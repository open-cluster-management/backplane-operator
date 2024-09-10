// Copyright Contributors to the Open Cluster Management project

package renderer

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"

	loader "helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"

	v1 "github.com/stolostron/backplane-operator/api/v1"
	"github.com/stolostron/backplane-operator/pkg/utils"
	"helm.sh/helm/v3/pkg/engine"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

const (
	AlwaysChartsDir = "pkg/templates/charts/always"
)

type Values struct {
	Global    Global    `json:"global" structs:"global"`
	HubConfig HubConfig `json:"hubconfig" structs:"hubconfig"`
	Org       string    `json:"org" structs:"org"`
}

type Global struct {
	ImageOverrides    map[string]string `json:"imageOverrides" structs:"imageOverrides"`
	TemplateOverrides map[string]string `json:"templateOverrides" structs:"templateOverrides"`
	PullPolicy        string            `json:"pullPolicy" structs:"pullPolicy"`
	PullSecret        string            `json:"pullSecret" structs:"pullSecret"`
	Namespace         string            `json:"namespace" structs:"namespace"`
	ConfigSecret      string            `json:"configSecret" structs:"configSecret"`
	HubSize           v1.HubSize        `json:"hubSize" structs:"hubSize" yaml:"hubSize"`
}

type HubConfig struct {
	NodeSelector         map[string]string `json:"nodeSelector" structs:"nodeSelector"`
	ProxyConfigs         map[string]string `json:"proxyConfigs" structs:"proxyConfigs"`
	ReplicaCount         int               `json:"replicaCount" structs:"replicaCount"`
	Tolerations          []Toleration      `json:"tolerations" structs:"tolerations"`
	OCPVersion           string            `json:"ocpVersion" structs:"ocpVersion"`
	ClusterIngressDomain string            `json:"clusterIngressDomain" structs:"clusterIngressDomain"`
	HubType              string            `json:"hubType" structs:"hubType"`
}

type Toleration struct {
	Key               string                    `json:"Key" protobuf:"bytes,1,opt,name=key"`
	Operator          corev1.TolerationOperator `json:"Operator" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"`
	Value             string                    `json:"Value" protobuf:"bytes,3,opt,name=value"`
	Effect            corev1.TaintEffect        `json:"Effect" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"`
	TolerationSeconds *int64                    `json:"TolerationSeconds" protobuf:"varint,5,opt,name=tolerationSeconds"`
}

func convertTolerations(tols []corev1.Toleration) []Toleration {
	var tolerations []Toleration
	for _, t := range tols {
		tolerations = append(tolerations, Toleration{
			Key:               t.Key,
			Operator:          t.Operator,
			Value:             t.Value,
			Effect:            t.Effect,
			TolerationSeconds: t.TolerationSeconds,
		})
	}
	return tolerations
}

func (u *Toleration) MarshalJSON() ([]byte, error) {
	v := reflect.ValueOf(u)
	values := make([]string, reflect.Indirect(v).NumField())
	var operator corev1.TolerationOperator = u.Operator
	var effect corev1.TaintEffect = u.Effect

	//Marshal all Toleration fields that are a number or true/false into a string
	for i := 0; i < reflect.Indirect(v).NumField(); i++ {
		switch reflect.Indirect(v).Field(i).Kind() {
		case reflect.String:
			if str, ok := reflect.Indirect(v).Field(i).Interface().(string); ok {
				if _, err := strconv.Atoi(str); err == nil {
					values[i] = fmt.Sprintf("'%s'", str)
				} else if _, err := strconv.ParseFloat(str, 64); err == nil {
					values[i] = fmt.Sprintf("'%s'", str)
				} else if _, err := strconv.ParseBool(str); err == nil {
					values[i] = fmt.Sprintf("'%s'", str)
				} else {
					values[i] = str
				}
			}
			if tol, ok := reflect.Indirect(v).Field(i).Interface().(corev1.TolerationOperator); ok {
				str := string(tol)
				if _, err := strconv.Atoi(str); err == nil {
					operator = corev1.TolerationOperator(fmt.Sprintf("'%s'", str))
				} else if _, err := strconv.ParseFloat(str, 64); err == nil {
					operator = corev1.TolerationOperator(fmt.Sprintf("'%s'", str))
				} else if _, err := strconv.ParseBool(str); err == nil {
					operator = corev1.TolerationOperator(fmt.Sprintf("'%s'", str))
				}
			}
			if eff, ok := reflect.Indirect(v).Field(i).Interface().(corev1.TaintEffect); ok {
				str := string(eff)
				if _, err := strconv.Atoi(str); err == nil {
					effect = corev1.TaintEffect(fmt.Sprintf("'%s'", str))
				} else if _, err := strconv.ParseFloat(str, 64); err == nil {
					effect = corev1.TaintEffect(fmt.Sprintf("'%s'", str))
				} else if _, err := strconv.ParseBool(str); err == nil {
					effect = corev1.TaintEffect(fmt.Sprintf("'%s'", str))
				}
			}
		}

	}

	return json.Marshal(&struct {
		Key               string                    `json:"Key" protobuf:"bytes,1,opt,name=key"`
		Operator          corev1.TolerationOperator `json:"Operator" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"`
		Value             string                    `json:"Value" protobuf:"bytes,3,opt,name=value"`
		Effect            corev1.TaintEffect        `json:"Effect" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"`
		TolerationSeconds *int64                    `json:"TolerationSeconds" protobuf:"varint,5,opt,name=tolerationSeconds"`
	}{
		Key:               values[0],
		Operator:          operator,
		Value:             values[2],
		Effect:            effect,
		TolerationSeconds: u.TolerationSeconds,
	})
}

func (val *Values) ToValues() (chartutil.Values, error) {
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	vals, err := chartutil.ReadValues(inrec)
	if err != nil {
		return vals, err
	}
	return vals, nil
}

func RenderCRDs(crdDir string, backplaneConfig *v1.MultiClusterEngine) ([]*unstructured.Unstructured, []error) {
	var crds []*unstructured.Unstructured
	errs := []error{}

	if val, ok := os.LookupEnv("DIRECTORY_OVERRIDE"); ok {
		crdDir = path.Join(val, crdDir)
	}

	// Read CRD files
	err := filepath.Walk(crdDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		crd := &unstructured.Unstructured{}
		if info == nil || info.IsDir() {
			return nil
		}

		bytesFile, e := os.ReadFile(path)
		if e != nil {
			errs = append(errs, fmt.Errorf("%s - error reading file: %v", info.Name(), err.Error()))
		}

		if err = yaml.Unmarshal(bytesFile, crd); err != nil {
			errs = append(errs, fmt.Errorf("%s - error unmarshalling file to unstructured: %v", info.Name(), err.Error()))
		}

		if backplaneConfig != nil {
			_, conversion, _ := unstructured.NestedMap(crd.Object, "spec", "conversion", "webhook", "clientConfig", "service")
			if conversion {
				crd.Object["spec"].(map[string]interface{})["conversion"].(map[string]interface{})["webhook"].(map[string]interface{})["clientConfig"].(map[string]interface{})["service"].(map[string]interface{})["namespace"] = backplaneConfig.Spec.TargetNamespace
			}
		}
		crds = append(crds, crd)
		return nil
	})
	if err != nil {
		return crds, errs
	}

	return crds, errs
}

func RenderCharts(chartDir string, backplaneConfig *v1.MultiClusterEngine, images map[string]string,
	templateOverrides map[string]string) ([]*unstructured.Unstructured, []error) {

	log := log.Log.WithName("reconcile")
	var templates []*unstructured.Unstructured
	errs := []error{}

	if val, ok := os.LookupEnv("DIRECTORY_OVERRIDE"); ok {
		chartDir = path.Join(val, chartDir)
	}

	// Read chart directory
	charts, err := os.ReadDir(chartDir)
	if err != nil {
		// Append error and return early since chart directory cannot be read
		errs = append(errs, err)
		return nil, errs
	}

	for _, chart := range charts {
		chartPath := filepath.Join(chartDir, chart.Name())
		chartTemplates, chartErrs := renderTemplates(chartPath, backplaneConfig, images, templateOverrides)

		if len(chartErrs) > 0 {
			for _, err := range chartErrs {
				log.Info(err.Error())
			}

			errs = append(errs, chartErrs...)
			continue // Continue processing other charts even if one fails
		}

		templates = append(templates, chartTemplates...)
	}

	// Return both collected templates and errors (if any)
	if len(errs) > 0 {
		return templates, errs
	}

	return templates, nil
}

func RenderChart(chartPath string, backplaneConfig *v1.MultiClusterEngine, images map[string]string,
	templates map[string]string) ([]*unstructured.Unstructured, []error) {

	log := log.Log.WithName("reconcile")
	if val, ok := os.LookupEnv("DIRECTORY_OVERRIDE"); ok {
		chartPath = path.Join(val, chartPath)
	}

	chartTemplates, errs := renderTemplates(chartPath, backplaneConfig, images, templates)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Info(err.Error())
		}
		return nil, errs
	}
	return chartTemplates, nil

}

// RenderChartWithNamespace wraps the RenderChart function, overriding the target namespace
func RenderChartWithNamespace(chartPath string, backplaneConfig *v1.MultiClusterEngine,
	images map[string]string, templates map[string]string, namespace string) ([]*unstructured.Unstructured, []error) {

	mce := backplaneConfig.DeepCopy()
	mce.Spec.TargetNamespace = namespace
	return RenderChart(chartPath, mce, images, templates)
}

func renderTemplates(chartPath string, backplaneConfig *v1.MultiClusterEngine, images map[string]string,
	templateOverrides map[string]string) ([]*unstructured.Unstructured, []error) {

	log := log.Log.WithName("reconcile")
	var templates []*unstructured.Unstructured
	errs := []error{}

	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Info(fmt.Sprintf("error loading chart: %s", chart.Name()))
		return nil, append(errs, err)
	}

	valuesYaml := &Values{}
	injectValuesOverrides(valuesYaml, backplaneConfig, images, templateOverrides)
	helmEngine := engine.Engine{
		Strict:   true,
		LintMode: false,
	}

	vals, err := valuesYaml.ToValues()
	if err != nil {
		log.Info(fmt.Sprintf("error rendering chart: %s", chart.Name()))
		return nil, append(errs, err)
	}

	rawTemplates, err := helmEngine.Render(chart, chartutil.Values{"Values": vals.AsMap()})
	if err != nil {
		log.Info(fmt.Sprintf("error rendering chart: %s", chart.Name()))
		return nil, append(errs, err)
	}

	for fileName, templateFile := range rawTemplates {
		unstructured := &unstructured.Unstructured{}
		if err = yaml.Unmarshal([]byte(templateFile), unstructured); err != nil {
			return nil, append(errs, fmt.Errorf("error converting file %s to unstructured", fileName))
		}

		utils.AddBackplaneConfigLabels(unstructured, backplaneConfig.Name)

		// Add namespace to namespaced resources
		switch unstructured.GetKind() {
		case "Deployment", "ServiceAccount", "Role", "RoleBinding", "Service", "ConfigMap", "Route":
			unstructured.SetNamespace(backplaneConfig.Spec.TargetNamespace)
		}
		templates = append(templates, unstructured)
	}

	return templates, errs
}

func injectValuesOverrides(values *Values, backplaneConfig *v1.MultiClusterEngine, images map[string]string,
	templates map[string]string) {

	values.Global.ImageOverrides = images

	values.Global.TemplateOverrides = templates

	values.Global.PullPolicy = string(utils.GetImagePullPolicy(backplaneConfig))

	values.Global.Namespace = backplaneConfig.Spec.TargetNamespace

	// TODO: remove this after backplaneConfig.Spec.HubSize is back
	values.Global.HubSize = utils.GetHubSize(backplaneConfig)

	// TODO: put this back later
	// values.Global.HubSize = backplaneConfig.Spec.HubSize

	values.Global.PullSecret = backplaneConfig.Spec.ImagePullSecret

	if v1.IsInHostedMode(backplaneConfig) {
		secretNN, err := utils.GetHostedCredentialsSecret(backplaneConfig)
		if err == nil {
			values.Global.ConfigSecret = secretNN.Name
		}
	}
	values.HubConfig.ReplicaCount = utils.DefaultReplicaCount(backplaneConfig)

	values.HubConfig.NodeSelector = backplaneConfig.Spec.NodeSelector

	if len(backplaneConfig.Spec.Tolerations) > 0 {
		values.HubConfig.Tolerations = convertTolerations(backplaneConfig.Spec.Tolerations)
	} else {
		values.HubConfig.Tolerations = convertTolerations(utils.DefaultTolerations())
	}

	values.Org = "open-cluster-management"

	values.HubConfig.OCPVersion = os.Getenv("ACM_HUB_OCP_VERSION")

	values.HubConfig.ClusterIngressDomain = os.Getenv("ACM_CLUSTER_INGRESS_DOMAIN")

	values.HubConfig.HubType = utils.GetHubType(backplaneConfig)

	if utils.ProxyEnvVarsAreSet() {
		proxyVar := map[string]string{}
		proxyVar["HTTP_PROXY"] = os.Getenv("HTTP_PROXY")
		proxyVar["HTTPS_PROXY"] = os.Getenv("HTTPS_PROXY")
		proxyVar["NO_PROXY"] = os.Getenv("NO_PROXY")
		values.HubConfig.ProxyConfigs = proxyVar
	}
}
