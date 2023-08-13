package k8s

import (
	"github.com/gin-gonic/gin"
	"kubemanagement.com/response"
)

type PVCApi struct{}

func (PVCApi) CreatePVC(c *gin.Context) {
	response.Success(c)
}
func (PVCApi) DeletePVC(c *gin.Context) {
	response.Success(c)
}
func (PVCApi) GetPVCList(c *gin.Context) {
	response.Success(c)
}
