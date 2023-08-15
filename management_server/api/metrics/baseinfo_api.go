package metrics

import (
	"github.com/gin-gonic/gin"
	"kubemanagement.com/response"
)

type MetricsApi struct {
}

func (*MetricsApi) GetMetricsDetailOrList(ctx *gin.Context) {
	response.SuccessWithDetailed(ctx, "获取Secret成功", nil)
}
