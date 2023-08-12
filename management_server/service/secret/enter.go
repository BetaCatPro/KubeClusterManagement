package secret

import "kubemanagement.com/convert"

type SeviceGroup struct {
	SecretService
}

var secretConvert = convert.ConvertGroupApp.SecretConvert
