package api

import (
	"kubemanagement.com/api/example"
	"kubemanagement.com/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
