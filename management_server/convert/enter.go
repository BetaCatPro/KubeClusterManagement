package convert

import (
	"kubemanagement.com/convert/configmap"
	"kubemanagement.com/convert/node"
	"kubemanagement.com/convert/pod"
	"kubemanagement.com/convert/secret"
)

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.ConvertGroup
	SecretConvert    secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
