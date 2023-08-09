package router

import (
	"kubemanagement.com/router/example"
	"kubemanagement.com/router/k8s"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8SRouterGroup     k8s.K8sRouter
}

var RouterGroupApp = new(RouterGroup)
