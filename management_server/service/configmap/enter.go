package configmap

import "kubemanagement.com/convert"

type ServiceGroup struct {
	ConfigMapService
}

var configConvert = convert.ConvertGroupApp.ConfigMapConvert
