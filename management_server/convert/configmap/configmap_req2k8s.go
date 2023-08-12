package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	configmapreq "kubemanagement.com/model/configmap/request"
	"kubemanagement.com/utils"
)

type Req2K8s struct {
}

func (*Req2K8s) CmReq2K8sConvert(configMapReq configmapreq.ConfigMap) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapReq.Name,
			Namespace: configMapReq.Namespace,
			Labels:    utils.ToMap(configMapReq.Labels),
		},
		Data: utils.ToMap(configMapReq.Data),
	}
}
