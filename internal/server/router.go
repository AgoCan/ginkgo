package server

import (
	"ginkgo/internal/handler/auth"
	"ginkgo/internal/handler/health"
)

// SetupRouter 初始化gin入口，路由信息
func (s *Server) SetupRouter() {
	s.Gin.GET("/health", health.HealthHandler())
	s.Gin.GET("/api/v1/login", auth.LoginHandler(s.Config, s.DB))
	s.Gin.POST("/api/v1/register", auth.RegisterHandler(s.Config, s.DB))
}
