package main

import (
	"kubemanagement.com/global"
	"kubemanagement.com/initiallize"
)

// 项目启动入口
func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8S()
	panic(r.Run(global.CONF.System.Addr))
}
