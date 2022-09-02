package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/middleware/cors"
	"ginkgo/internal/pkg/middleware/log"
	"ginkgo/internal/router"
)

type Server struct {
	Config *config.Config
	Gin    *gin.Engine
	// 导入日志
	Log *log.Client
	DB  *database.Client
}

func NewServer() *Server {
	return &Server{
		Gin: gin.New(),
	}
}

func (s *Server) Run() {

	c := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})

	s.Gin.Use(log.GinLogger(s.Log.Logger),
		log.GinRecovery(s.Log.Logger, true),
		c)
	router.SetupRouter(s.Gin)

	err := s.Gin.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
}
