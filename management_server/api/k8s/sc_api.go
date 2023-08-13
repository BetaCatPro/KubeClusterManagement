package k8s

import (
	"github.com/gin-gonic/gin"
	"kubemanagement.com/response"
)

type SCApi struct{}

func (SCApi) CreateSC(c *gin.Context) {
	response.Success(c)
}
func (SCApi) DeleteSC(c *gin.Context) {
	response.Success(c)
}
func (SCApi) GetSCList(c *gin.Context) {
	response.Success(c)
}
