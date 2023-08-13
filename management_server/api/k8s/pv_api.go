package k8s

import (
	"github.com/gin-gonic/gin"
	"kubemanagement.com/response"
)

type PVApi struct{}

func (PVApi) CreatePV(c *gin.Context) {
	response.Success(c)
}
func (PVApi) DeletePV(c *gin.Context) {
	response.Success(c)
}
func (PVApi) GetPVList(c *gin.Context) {
	response.Success(c)
}
