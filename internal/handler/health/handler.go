package health

import (
	"ginkgo/internal/pkg/middleware/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler() func(c *gin.Context) {

	return func(c *gin.Context) {
		service := Health{}
		log.Sugar.Info("abc")
		res := service.Status()
		c.JSON(http.StatusOK, res)
	}
}
