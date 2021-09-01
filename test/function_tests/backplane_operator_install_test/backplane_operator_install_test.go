// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package backplane_install_test

import (
	"context"
<<<<<<< HEAD
	"io/ioutil"

	"github.com/ghodss/yaml"
=======
	"fmt"
>>>>>>> add node selector to spec
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"

	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"

	backplane "github.com/open-cluster-management/backplane-operator/api/v1alpha1"
<<<<<<< HEAD

=======
	appsv1 "k8s.io/api/apps/v1"
	// apierrors "k8s.io/apimachinery/pkg/api/errors"
>>>>>>> add node selector to spec
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/types"
)

const (
	BackplaneConfigName        = "backplane"
	BackplaneOperatorNamespace = "backplane-operator-system"
	installTimeout             = time.Minute * 5
	duration                   = time.Second * 1
	interval                   = time.Millisecond * 250
)

var (
	ctx                = context.Background()
	globalsInitialized = false
	baseURL            = ""

	k8sClient client.Client

<<<<<<< HEAD
	backplaneConfig = types.NamespacedName{}

	blockCreationResources = []struct {
		Name     string
		GVK      schema.GroupVersionKind
		Filepath string
		crdPath  string
		Expected string
	}{
		{
			Name: "MultiClusterHub",
			GVK: schema.GroupVersionKind{
				Group:   "operator.open-cluster-management.io",
				Version: "v1",
				Kind:    "MultiClusterHub",
			},
			Filepath: "../resources/multiclusterhub.yaml",
			crdPath:  "../resources/multiclusterhub_crd.yaml",
			Expected: "Existing MultiClusterHub resources must first be deleted",
		},
	}
	blockDeletionResources = []struct {
		Name     string
		GVK      schema.GroupVersionKind
		Filepath string
		crdPath  string
		Expected string
	}{
		{
			Name: "BareMetalAsset",
			GVK: schema.GroupVersionKind{
				Group:   "inventory.open-cluster-management.io",
				Version: "v1alpha1",
				Kind:    "BareMetalAsset",
			},
			Filepath: "../resources/baremetalassets.yaml",
			Expected: "Existing BareMetalAsset resources must first be deleted",
		},
		{
			Name: "MultiClusterObservability",
			GVK: schema.GroupVersionKind{
				Group:   "observability.open-cluster-management.io",
				Version: "v1beta2",
				Kind:    "MultiClusterObservability",
			},
			crdPath:  "../resources/multiclusterobservabilities_crd.yaml",
			Filepath: "../resources/multiclusterobservability.yaml",
			Expected: "Existing MultiClusterObservability resources must first be deleted",
		},
		{
			Name: "ManagedCluster",
			GVK: schema.GroupVersionKind{
				Group:   "cluster.open-cluster-management.io",
				Version: "v1",
				Kind:    "ManagedClusterList",
			},
			Filepath: "../resources/managedcluster.yaml",
			Expected: "Existing ManagedCluster resources must first be deleted",
		},
	}
=======
	backplaneConfig       = types.NamespacedName{}
	backplaneNodeSelector map[string]string
>>>>>>> add node selector to spec
)

func initializeGlobals() {
	// baseURL = *BaseURL
	backplaneConfig = types.NamespacedName{
		Name: BackplaneConfigName,
	}
	backplaneNodeSelector = map[string]string{"beta.kubernetes.io/os": "linux"}
}

var _ = Describe("BackplaneConfig Test Suite", func() {

	BeforeEach(func() {
		if !globalsInitialized {
			initializeGlobals()
			globalsInitialized = true
		}
	})

	Context("Creating a BackplaneConfig", func() {
		It("Should install all components ", func() {
			By("By creating a new BackplaneConfig", func() {
				Expect(k8sClient.Create(ctx, defaultBackplaneConfig())).Should(Succeed())
			})
		})

		It("Should check that all components were installed correctly", func() {
			By("Ensuring the BackplaneConfig becomes available", func() {
				Eventually(func() bool {
					key := &backplane.MultiClusterEngine{}
					k8sClient.Get(context.Background(), types.NamespacedName{
						Name: BackplaneConfigName,
					}, key)
					return key.Status.Phase == backplane.MultiClusterEnginePhaseAvailable
				}, installTimeout, interval).Should(BeTrue())

			})
		})

		It("Should check for a healthy status", func() {
			config := &backplane.MultiClusterEngine{}
			Expect(k8sClient.Get(ctx, backplaneConfig, config)).To(Succeed())

			By("Checking the phase", func() {
				Expect(config.Status.Phase).To(Equal(backplane.MultiClusterEnginePhaseAvailable))
			})
			By("Checking the components", func() {
				Expect(len(config.Status.Components)).Should(BeNumerically(">=", 6), "Expected at least 6 components in status")
			})
			By("Checking the conditions", func() {
				available := backplane.MultiClusterEngineCondition{}
				for _, c := range config.Status.Conditions {
					if c.Type == backplane.MultiClusterEngineAvailable {
						available = c
					}
				}
				Expect(available.Status).To(Equal(metav1.ConditionTrue))
			})
		})
<<<<<<< HEAD

		It("Should ensure validatingwebhook blocks deletion if resouces exist", func() {
			for _, r := range blockDeletionResources {
				By("Creating a new "+r.Name, func() {

					if r.crdPath != "" {
						applyResource(r.crdPath)
						defer deleteResource(r.crdPath)
					}
					applyResource(r.Filepath)
					defer deleteResource(r.Filepath)

					config := &backplane.MultiClusterEngine{}
					Expect(k8sClient.Get(ctx, backplaneConfig, config)).To(Succeed()) // Get Backplaneconfig

					err := k8sClient.Delete(ctx, config) // Attempt to delete backplaneconfig. Ensure it does not succeed.
					Expect(err).ShouldNot(BeNil())
					Expect(err.Error()).Should(ContainSubstring(r.Expected))
				})
			}
		})

		It("Should ensure validatingwebhook blocks creation if resouces exist", func() {
			for _, r := range blockCreationResources {
				By("Creating a new "+r.Name, func() {

					if r.crdPath != "" {
						applyResource(r.crdPath)
						defer deleteResource(r.crdPath)
					}
					applyResource(r.Filepath)
					defer deleteResource(r.Filepath)

					backplaneConfig := defaultBackplaneConfig()
					backplaneConfig.Name = "test"

					err := k8sClient.Create(ctx, backplaneConfig)
					Expect(err).ShouldNot(BeNil())
					Expect(err.Error()).Should(ContainSubstring(r.Expected))
				})
=======
		It("Should check that the config spec has propagated", func() {


			tests := []struct {
				Name           string
				NamespacedName types.NamespacedName
				ResourceType   client.Object

			}{
				{
					Name:           "OCM Webhook",
					NamespacedName: types.NamespacedName{Name: "ocm-webhook", Namespace: BackplaneOperatorNamespace},
					ResourceType:   &appsv1.Deployment{},
				},
				{
					Name:           "OCM Controller",
					NamespacedName: types.NamespacedName{Name: "ocm-controller", Namespace: BackplaneOperatorNamespace},
					ResourceType:   &appsv1.Deployment{},
				},
				{
					Name:           "OCM Proxy Server",
					NamespacedName: types.NamespacedName{Name: "ocm-proxyserver", Namespace: BackplaneOperatorNamespace},
					ResourceType:   &appsv1.Deployment{},
				},
				{
					Name:           "Cluster Manager Deployment",
					NamespacedName: types.NamespacedName{Name: "cluster-manager", Namespace: BackplaneOperatorNamespace},
					ResourceType:   &appsv1.Deployment{},
				},
				{
					Name:           "Hive Operator Deployment",
					NamespacedName: types.NamespacedName{Name: "hive-operator", Namespace: BackplaneOperatorNamespace},
					ResourceType:   &appsv1.Deployment{},
				},
			}

			By("Ensuring the spec is correct")
			for _, test := range tests {

				Eventually(func() bool {
					// component := &unstructured.Unstructured{}
					err := k8sClient.Get(ctx, test.NamespacedName, test.ResourceType)
					if err != nil {
						fmt.Fprintf(GinkgoWriter, "could not get component %s\n", test.Name)
					}
					
					componentSelector := test.ResourceType.(*appsv1.Deployment).Spec.Template.Spec.NodeSelector


					return reflect.DeepEqual(componentSelector, backplaneNodeSelector)
					
				}, installTimeout, interval).Should(BeTrue())

>>>>>>> add node selector to spec
			}
		})
	})
})

func applyResource(resourceFile string) {
	resourceData, err := ioutil.ReadFile(resourceFile) // Get resource as bytes
	Expect(err).To(BeNil())

	unstructured := &unstructured.Unstructured{Object: map[string]interface{}{}}
	err = yaml.Unmarshal(resourceData, &unstructured.Object) // Render resource as unstructured
	Expect(err).To(BeNil())

	Expect(k8sClient.Create(ctx, unstructured)).Should(Succeed()) // Create resource on cluster
}

func deleteResource(resourceFile string) {
	resourceData, err := ioutil.ReadFile(resourceFile) // Get resource as bytes
	Expect(err).To(BeNil())

	unstructured := &unstructured.Unstructured{Object: map[string]interface{}{}}
	err = yaml.Unmarshal(resourceData, &unstructured.Object) // Render resource as unstructured
	Expect(err).To(BeNil())

	Expect(k8sClient.Delete(ctx, unstructured)).Should(Succeed()) // Delete resource on cluster
}

func defaultBackplaneConfig() *backplane.MultiClusterEngine {
	return &backplane.MultiClusterEngine{
		ObjectMeta: metav1.ObjectMeta{
			Name: BackplaneConfigName,
		},
<<<<<<< HEAD
		Spec: backplane.MultiClusterEngineSpec{
			Foo: "bar",
=======
		Spec: backplane.BackplaneConfigSpec{
			Foo:          "bar",
			NodeSelector: backplaneNodeSelector,
>>>>>>> add node selector to spec
		},
		Status: backplane.MultiClusterEngineStatus{
			Phase: "",
		},
	}
}
