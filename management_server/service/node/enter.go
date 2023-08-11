package node

import "kubemanagement.com/convert"

type Group struct {
	NodeService
}

var nodeConvert = convert.ConvertGroupApp.NodeConvert
