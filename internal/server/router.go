package server

import (
	"github.com/gin-gonic/gin"

	"ginkgo/internal/handler/auth"
	"ginkgo/internal/handler/health"
)

// SetupRouter 初始化gin入口，路由信息
func (s *Server) SetupRouter() {
	s.Gin.GET("/health", health.HealthHandler())
	v1NoAuth := s.Gin.Group("/api/v1")
	s.authRouter(v1NoAuth)

	v1Auth := s.Gin.Group("/api/v1")
	v1Auth.Use(jwtAuth(s.Config, s.DB))
	v1Auth.GET("/auth/health", health.HealthHandler())
}

func (s *Server) authRouter(engine *gin.RouterGroup) {
	engine.POST("/login", auth.LoginHandler(s.Config, s.DB))
	engine.POST("/register", auth.RegisterHandler(s.Config, s.DB))

}
