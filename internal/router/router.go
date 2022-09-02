package router

import (
	"github.com/gin-gonic/gin"

	"ginkgo/internal/handler/health"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter(r *gin.Engine) {
	r.GET("/health", health.HealthHandler)
}
