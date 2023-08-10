package pod

import "kubemanagement.com/convert"

type PodServiceGroup struct {
	PodService
}

var podConvert = convert.ConvertGroupApp.PodConvert
