package service

import (
	"kubemanagement.com/service/configmap"
	"kubemanagement.com/service/node"
	"kubemanagement.com/service/pod"
	"kubemanagement.com/service/secret"
)

type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
	SecretServiceGroup    secret.SeviceGroup
}

var ServiceGroupApp = new(ServiceGroup)
