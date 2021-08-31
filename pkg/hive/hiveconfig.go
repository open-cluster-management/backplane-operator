// Copyright Contributors to the Open Cluster Management project

package hive

import (
	v1alpha1 "github.com/open-cluster-management/backplane-operator/api/v1alpha1"
	"github.com/open-cluster-management/backplane-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func HiveConfig(bpc *v1alpha1.MultiClusterEngine) *unstructured.Unstructured {

	cm := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "hive.openshift.io/v1",
			"kind":       "HiveConfig",
			"metadata": map[string]interface{}{
				"name": "hive",
			},
			"spec": map[string]interface{}{},
		},
	}

	utils.AddBackplaneConfigLabels(cm, bpc.GetName())

	return cm
}
