package harbor

import (
	"github.com/gin-gonic/gin"
	"kubemanagement.com/response"
)

type ProjectsApi struct {
}

func (*ProjectsApi) GetProjectsDetailOrList(ctx *gin.Context) {
	response.SuccessWithDetailed(ctx, "获取Secret成功", nil)
}
