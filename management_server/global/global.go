package global

import (
	"k8s.io/client-go/kubernetes"
	"kubemanagement.com/config"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
