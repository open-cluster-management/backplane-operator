// Copyright Contributors to the Open Cluster Management project

package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	pkgerrors "github.com/pkg/errors"
	backplanev1 "github.com/stolostron/backplane-operator/api/v1"
	"github.com/stolostron/backplane-operator/pkg/foundation"
	"github.com/stolostron/backplane-operator/pkg/overrides"
	renderer "github.com/stolostron/backplane-operator/pkg/rendering"
	"github.com/stolostron/backplane-operator/pkg/status"
	"github.com/stolostron/backplane-operator/pkg/toggle"
	"github.com/stolostron/backplane-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var ErrBadFormat = errors.New("bad format")

func (r *MultiClusterEngineReconciler) HostedReconcile(ctx context.Context, mce *backplanev1.MultiClusterEngine) (
	retRes ctrl.Result, retErr error) {
	r.Log = log.Log.WithName("reconcile")

	defer func() {
		mce.Status = r.StatusManager.ReportStatus(*mce)
		err := r.Client.Status().Update(ctx, mce)
		if mce.Status.Phase != backplanev1.MultiClusterEnginePhaseAvailable && !utils.IsPaused(mce) {
			retRes = ctrl.Result{RequeueAfter: requeuePeriod}
		}
		if err != nil {
			retErr = err
		}
	}()

	// If deletion detected, finalize backplane config
	if mce.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(mce, backplaneFinalizer) {
			err := r.finalizeHostedBackplaneConfig(ctx, mce) // returns all errors
			if err != nil {
				r.Log.Info(err.Error())
				return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
			}

			r.Log.Info("all subcomponents have been finalized successfully - removing finalizer")
			controllerutil.RemoveFinalizer(mce, backplaneFinalizer)
			if err := r.Client.Update(ctx, mce); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, nil // Object finalized successfully
	}

	// Add finalizer for this CR
	if !controllerutil.ContainsFinalizer(mce, backplaneFinalizer) {
		controllerutil.AddFinalizer(mce, backplaneFinalizer)
		if err := r.Client.Update(ctx, mce); err != nil {
			return ctrl.Result{}, err
		}
	}

	var result ctrl.Result
	var err error

	result, err = r.setHostedDefaults(ctx, mce)
	if result != (ctrl.Result{}) {
		return ctrl.Result{}, err
	}
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	result, err = r.validateNamespace(ctx, mce)
	if result != (ctrl.Result{}) {
		return ctrl.Result{}, err
	}
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	result, err = r.validateImagePullSecret(ctx, mce)
	if result != (ctrl.Result{}) {
		return ctrl.Result{}, err
	}
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	// Attempt to retrieve image overrides from environmental variables.
	imageOverrides := overrides.GetOverridesFromEnv(overrides.OperandImagePrefix)

	// If no overrides found using OperandImagePrefix, attempt to retrieve using OSBSImagePrefix.
	if len(imageOverrides) == 0 {
		imageOverrides = overrides.GetOverridesFromEnv(overrides.OSBSImagePrefix)
	}

	// Check if no image overrides were found using either prefix.
	if len(imageOverrides) == 0 {
		r.Log.Error(err, "Could not get map of image overrides")

		// If images are not set from environmental variables, fail
		r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing,
			metav1.ConditionFalse, status.RequirementsNotMetReason, "No image references defined in deployment"))

		return ctrl.Result{RequeueAfter: requeuePeriod}, errors.New(
			"no image references exist. images must be defined as environment variables")
	}

	// Apply image repository override from annotation if present.
	if imageRepo := utils.GetImageRepository(mce); imageRepo != "" {
		r.Log.Info(fmt.Sprintf("Overriding Image Repository from annotation 'imageRepository': %s", imageRepo))
		imageOverrides = utils.OverrideImageRepository(imageOverrides, imageRepo)
	}

	// Check for developer overrides in configmap.
	if cmName := utils.GetImageOverridesConfigmapName(mce); cmName != "" {
		imageOverrides, err = overrides.GetOverridesFromConfigmap(r.Client, imageOverrides,
			utils.OperatorNamespace(), cmName, false)

		if err != nil {
			r.Log.Error(err, fmt.Sprintf("Failed to find image override configmap: %s/%s",
				utils.OperatorNamespace(), cmName))

			r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing,
				metav1.ConditionFalse, status.RequirementsNotMetReason,
				fmt.Sprintf("Issue building image references: %s", err.Error())))

			return ctrl.Result{}, err
		}
	}

	// Update cache with image overrides and related information.
	r.CacheSpec.ImageOverrides = imageOverrides
	r.CacheSpec.ImageRepository = utils.GetImageRepository(mce)
	r.CacheSpec.ImageOverridesCM = utils.GetImageOverridesConfigmapName(mce)

	// Attempt to retrieve template overrides from environmental variables.
	templateOverrides := overrides.GetOverridesFromEnv(overrides.TemplateOverridePrefix)

	// Check for developer overrides in configmap
	if toConfigmapName := utils.GetTemplateOverridesConfigmapName(mce); toConfigmapName != "" {
		templateOverrides, err = overrides.GetOverridesFromConfigmap(r.Client, templateOverrides,
			utils.OperatorNamespace(), toConfigmapName, true)

		if err != nil {
			r.Log.Error(err, fmt.Sprintf("Failed to find template override configmap: %s/%s",
				utils.OperatorNamespace(), toConfigmapName))

			return ctrl.Result{}, err
		}
	}

	// Update cache with template overrides and related information.
	r.CacheSpec.TemplateOverrides = templateOverrides
	r.CacheSpec.TemplateOverridesCM = utils.GetTemplateOverridesConfigmapName(mce)

	// Do not reconcile objects if this instance of mce is labeled "paused"
	if utils.IsPaused(mce) {
		r.Log.Info("MultiClusterEngine reconciliation is paused. Nothing more to do.")
		r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing,
			metav1.ConditionUnknown, status.PausedReason, "Multiclusterengine is paused"))

		return ctrl.Result{}, nil
	}

	hostedClient, err := r.GetHostedClient(ctx, mce)
	if err != nil {
		r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing,
			metav1.ConditionFalse, status.RequirementsNotMetReason,
			fmt.Sprintf("couldn't connect to hosted environment: %s", err.Error())))

		return ctrl.Result{RequeueAfter: requeuePeriod}, err
	}

	err = hostedClient.Create(ctx, &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: mce.Spec.TargetNamespace},
	})
	if err != nil && !apierrors.IsAlreadyExists(err) {
		r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing,
			metav1.ConditionFalse, status.RequirementsNotMetReason, err.Error()))

		return ctrl.Result{RequeueAfter: requeuePeriod}, err
	}

	// Render CRD templates
	crdsDir := "pkg/templates/hosted-crds"
	crds, errs := renderer.RenderCRDs(crdsDir, mce)
	for _, err := range errs {
		r.Log.Info(err.Error())
	}
	if len(errs) > 0 {
		return ctrl.Result{RequeueAfter: requeuePeriod}, nil
	}

	for _, template := range crds {
		result, err := r.hostedApplyTemplate(ctx, mce, template, hostedClient)
		if err != nil {
			return result, err
		}
	}

	// Create hosted ClusterManager
	if mce.Enabled(backplanev1.ClusterManager) {
		result, err := r.ensureHostedClusterManager(ctx, mce)
		if result != (ctrl.Result{}) {
			return result, err
		}
	} else {
		result, err := r.ensureNoHostedClusterManager(ctx, mce)
		if result != (ctrl.Result{}) {
			return result, err
		}
	}

	// Create hosted ClusterManager
	if mce.Enabled(backplanev1.ServerFoundation) {
		result, err := r.ensureHostedImport(ctx, mce, hostedClient)
		if result != (ctrl.Result{}) {
			return result, err
		}
	} else {
		result, err := r.ensureNoHostedImport(ctx, mce, hostedClient)
		if result != (ctrl.Result{}) {
			return result, err
		}
	}

	r.StatusManager.AddCondition(status.NewCondition(backplanev1.MultiClusterEngineProgressing, metav1.ConditionTrue,
		status.DeploySuccessReason, "All components deployed"))

	return ctrl.Result{}, nil
}

func (r *MultiClusterEngineReconciler) GetHostedClient(ctx context.Context, mce *backplanev1.MultiClusterEngine) (
	client.Client, error) {
	secretNN, err := utils.GetHostedCredentialsSecret(mce)
	if err != nil {
		return nil, err
	}

	// Parse Kube credentials from secret
	kubeConfigSecret := &corev1.Secret{}
	if err := r.Client.Get(context.TODO(), secretNN, kubeConfigSecret); err != nil {
		if apierrors.IsNotFound(err) {
			if err != nil {
				return nil, err
			}
		}
	}
	kubeconfig, err := parseKubeCreds(kubeConfigSecret)
	if err != nil {
		return nil, err
	}

	restconfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	uncachedClient, err := client.New(restconfig, client.Options{
		Scheme: r.Client.Scheme(),
	})
	if err != nil {
		return nil, err
	}

	return uncachedClient, nil
}

// parseKubeCreds takes a secret cotaining credentials and returns the stored Kubeconfig.
func parseKubeCreds(secret *corev1.Secret) ([]byte, error) {
	kubeconfig, ok := secret.Data["kubeconfig"]
	if !ok {
		return []byte{}, fmt.Errorf("%s: %w", secret.Name, ErrBadFormat)
	}

	return kubeconfig, nil
}

func (r *MultiClusterEngineReconciler) ensureHostedImport(ctx context.Context,
	backplaneConfig *backplanev1.MultiClusterEngine, hostedClient client.Client) (ctrl.Result, error) {
	log := log.Log.WithName("reconcile")

	r.StatusManager.AddComponent(toggle.EnabledStatus(types.NamespacedName{Name: "managedcluster-import-controller-v2",
		Namespace: backplaneConfig.Spec.TargetNamespace}))

	r.StatusManager.RemoveComponent(toggle.DisabledStatus(
		types.NamespacedName{Name: "managedcluster-import-controller-v2",
			Namespace: backplaneConfig.Spec.TargetNamespace}, []*unstructured.Unstructured{}))

	templates, errs := renderer.RenderChart(toggle.HostingImportChartDir, backplaneConfig,
		r.CacheSpec.ImageOverrides, r.CacheSpec.TemplateOverrides)

	if len(errs) > 0 {
		for _, err := range errs {
			log.Info(err.Error())
		}
		return ctrl.Result{RequeueAfter: requeuePeriod}, nil
	}

	// Applies all templates
	for _, template := range templates {
		result, err := r.applyTemplate(ctx, backplaneConfig, template)
		if err != nil {
			return result, err
		}
	}

	templates, errs = renderer.RenderChart(toggle.HostedImportChartDir, backplaneConfig,
		r.CacheSpec.ImageOverrides, r.CacheSpec.TemplateOverrides)

	if len(errs) > 0 {
		for _, err := range errs {
			log.Info(err.Error())
		}
		return ctrl.Result{RequeueAfter: requeuePeriod}, nil
	}

	// Applies all templates
	for _, template := range templates {
		result, err := r.hostedApplyTemplate(ctx, backplaneConfig, template, hostedClient)
		if err != nil {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MultiClusterEngineReconciler) ensureNoHostedImport(ctx context.Context,
	backplaneConfig *backplanev1.MultiClusterEngine, hostedClient client.Client) (ctrl.Result, error) {
	log := log.Log.WithName("reconcile")

	r.StatusManager.RemoveComponent(toggle.EnabledStatus(
		types.NamespacedName{Name: "managedcluster-import-controller-v2",
			Namespace: backplaneConfig.Spec.TargetNamespace}))

	r.StatusManager.AddComponent(toggle.DisabledStatus(
		types.NamespacedName{Name: "managedcluster-import-controller-v2",
			Namespace: backplaneConfig.Spec.TargetNamespace}, []*unstructured.Unstructured{}))

	templates, errs := renderer.RenderChart(toggle.HostingImportChartDir, backplaneConfig,
		r.CacheSpec.ImageOverrides, r.CacheSpec.TemplateOverrides)

	if len(errs) > 0 {
		for _, err := range errs {
			log.Info(err.Error())
		}
		return ctrl.Result{RequeueAfter: requeuePeriod}, nil
	}

	// Applies all templates
	for _, template := range templates {
		result, err := r.deleteTemplate(ctx, backplaneConfig, template)
		if err != nil {
			return result, err
		}
	}

	templates, errs = renderer.RenderChart(toggle.HostedImportChartDir, backplaneConfig,
		r.CacheSpec.ImageOverrides, r.CacheSpec.TemplateOverrides)

	if len(errs) > 0 {
		for _, err := range errs {
			log.Info(err.Error())
		}
		return ctrl.Result{RequeueAfter: requeuePeriod}, nil
	}

	// Applies all templates
	for _, template := range templates {
		result, err := r.hostedDeleteTemplate(ctx, backplaneConfig, template, hostedClient)
		if err != nil {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MultiClusterEngineReconciler) ensureHostedClusterManager(ctx context.Context,
	mce *backplanev1.MultiClusterEngine) (ctrl.Result, error) {
	log := log.Log.WithName("reconcile")
	cmName := fmt.Sprintf("%s-cluster-manager", mce.Name)

	r.StatusManager.AddComponent(status.ClusterManagerStatus{
		NamespacedName: types.NamespacedName{Name: cmName},
	})

	// Apply namespace
	newNs := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: cmName,
		},
	}
	checkNs := &corev1.Namespace{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{Name: cmName}, checkNs)
	if err != nil && apierrors.IsNotFound(err) {
		err = r.Client.Create(context.TODO(), newNs)
		if err != nil {
			log.Error(err, "Could not create namespace")
			return ctrl.Result{}, err
		}
		log.Info("Namespace for hosted cluster-manager created")
		return ctrl.Result{Requeue: true}, nil
	}
	if err != nil && !apierrors.IsNotFound(err) {
		return ctrl.Result{Requeue: true}, err
	}

	// Apply secret in namespace
	kubeconfigSecret := &corev1.Secret{}
	secretNN, err := utils.GetHostedCredentialsSecret(mce)
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}
	err = r.Client.Get(ctx, secretNN, kubeconfigSecret)
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	cmSecret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kubeconfigSecret.APIVersion,
			Kind:       kubeconfigSecret.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "external-hub-kubeconfig",
			Namespace: cmName,
			Labels:    kubeconfigSecret.Labels,
		},
		Data: kubeconfigSecret.Data,
		Type: corev1.SecretTypeOpaque,
	}

	force := true
	err = r.Client.Patch(context.TODO(), cmSecret, client.Apply,
		&client.PatchOptions{Force: &force, FieldManager: "backplane-operator"})

	if err != nil {
		log.Info(fmt.Sprintf("Error applying kubeconfig secret to hosted cluster-manager namespace: %s", err.Error()))
		return ctrl.Result{Requeue: true}, nil
	}

	// Apply clustermanager
	cmTemplate := foundation.HostedClusterManager(mce, r.CacheSpec.ImageOverrides)
	if err := ctrl.SetControllerReference(mce, cmTemplate, r.Scheme); err != nil {
		return ctrl.Result{}, fmt.Errorf("error setting controller reference on resource `%s`: %w",
			cmTemplate.GetName(), err)
	}

	force = true
	err = r.Client.Patch(ctx, cmTemplate, client.Apply,
		&client.PatchOptions{Force: &force, FieldManager: "backplane-operator"})

	if err != nil {
		return ctrl.Result{}, fmt.Errorf("error applying object Name: %s Kind: %s, %w", cmTemplate.GetName(),
			cmTemplate.GetKind(), err)
	}

	return ctrl.Result{}, nil
}

func (r *MultiClusterEngineReconciler) ensureNoHostedClusterManager(ctx context.Context,
	mce *backplanev1.MultiClusterEngine) (ctrl.Result, error) {
	cmName := fmt.Sprintf("%s-cluster-manager", mce.Name)

	r.StatusManager.RemoveComponent(status.ClusterManagerStatus{
		NamespacedName: types.NamespacedName{Name: cmName},
	})

	// Delete clustermanager
	clusterManager := &unstructured.Unstructured{}
	clusterManager.SetGroupVersionKind(
		schema.GroupVersionKind{
			Group:   "operator.open-cluster-management.io",
			Version: "v1",
			Kind:    "ClusterManager",
		},
	)
	err := r.Client.Get(ctx, types.NamespacedName{Name: cmName}, clusterManager)
	if err == nil { // If resource exists, delete
		err := r.Client.Delete(ctx, clusterManager)
		if err != nil {
			return ctrl.Result{RequeueAfter: requeuePeriod}, err
		}
	} else if err != nil && !apierrors.IsNotFound(err) { // Return error, if error is not not found error
		return ctrl.Result{RequeueAfter: requeuePeriod}, err
	}

	// Verify clustermanager namespace deleted
	checkNs := &corev1.Namespace{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: cmName}, checkNs)
	if err == nil {
		return ctrl.Result{RequeueAfter: requeuePeriod}, fmt.Errorf(
			"waiting for hosted-clustermanager namespace to be terminated before proceeding with clustermanager cleanup")
	}

	if err != nil && !apierrors.IsNotFound(err) { // Return error, if error is not not found error
		return ctrl.Result{RequeueAfter: requeuePeriod}, err
	}

	return ctrl.Result{}, nil
}

// setHostedDefaults configures the MCE with default values and updates
func (r *MultiClusterEngineReconciler) setHostedDefaults(ctx context.Context, m *backplanev1.MultiClusterEngine) (
	ctrl.Result, error) {
	r.Log = log.Log.WithName("reconcile")

	updateNecessary := false
	if !utils.AvailabilityConfigIsValid(m.Spec.AvailabilityConfig) {
		m.Spec.AvailabilityConfig = backplanev1.HAHigh
		updateNecessary = true
	}

	if len(m.Spec.TargetNamespace) == 0 {
		m.Spec.TargetNamespace = backplanev1.DefaultTargetNamespace
		updateNecessary = true
	}

	if utils.SetHostedDefaultComponents(m) {
		updateNecessary = true
	}

	if utils.DeduplicateComponents(m) {
		updateNecessary = true
	}

	// Apply defaults to server
	if updateNecessary {
		r.Log.Info("Setting hosted defaults")
		err := r.Client.Update(ctx, m)
		if err != nil {
			r.Log.Error(err, "Failed to update MultiClusterEngine")
			return ctrl.Result{}, err
		}

		r.Log.Info("MultiClusterEngine successfully updated")
		return ctrl.Result{Requeue: true}, nil
	} else {
		return ctrl.Result{}, nil
	}
}

func (r *MultiClusterEngineReconciler) finalizeHostedBackplaneConfig(ctx context.Context,
	mce *backplanev1.MultiClusterEngine) error {
	_, err := r.ensureNoHostedClusterManager(ctx, mce)
	if err != nil {
		return err
	}
	if utils.IsUnitTest() {
		return nil
	}
	hostedClient, err := r.GetHostedClient(ctx, mce)
	if err != nil {
		return err
	}
	_, err = r.ensureNoHostedImport(ctx, mce, hostedClient)
	if err != nil {
		return err
	}
	return nil
}

func (r *MultiClusterEngineReconciler) hostedApplyTemplate(ctx context.Context,
	backplaneConfig *backplanev1.MultiClusterEngine, template *unstructured.Unstructured, hostedClient client.Client) (
	ctrl.Result, error) {
	// Set owner reference.
	err := ctrl.SetControllerReference(backplaneConfig, template, r.Scheme)
	if err != nil {
		return ctrl.Result{}, pkgerrors.Wrapf(err, "Error setting controller reference on resource %s", template.GetName())
	}

	// Apply the object data.
	force := true
	err = hostedClient.Patch(ctx, template, client.Apply,
		&client.PatchOptions{Force: &force, FieldManager: "backplane-operator"})

	if err != nil {
		return ctrl.Result{}, pkgerrors.Wrapf(err, "error applying object Name: %s Kind: %s", template.GetName(),
			template.GetKind())
	}

	return ctrl.Result{}, nil
}

/*
deleteTemplate return true if resource does not exist and returns an error if a GET or DELETE errors unexpectedly.
A false response without error means the resource is in the process of deleting.
*/
func (r *MultiClusterEngineReconciler) hostedDeleteTemplate(ctx context.Context,
	backplaneConfig *backplanev1.MultiClusterEngine, template *unstructured.Unstructured, hostedClient client.Client) (
	ctrl.Result, error) {
	r.Log = log.Log.WithName("reconcile")

	err := hostedClient.Get(ctx, types.NamespacedName{Name: template.GetName(), Namespace: template.GetNamespace()},
		template)

	if err != nil && apierrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	}

	// set status progressing condition
	if err != nil {
		r.Log.Error(err, "Odd error delete template")
		return ctrl.Result{}, err
	}

	r.Log.Info(fmt.Sprintf("finalizing template: %s\n", template.GetName()))
	err = hostedClient.Delete(ctx, template)
	if err != nil {
		r.Log.Error(err, "Failed to delete template")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
