package auth

import (
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/response"
	"net/http"

	"ginkgo/internal/config"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(
	config *config.Config,
	db *database.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		service := Registration{}
		if err := c.ShouldBind(&service); err != nil {
			c.JSON(200, response.Error(response.ErrCodeParameter))
		} else {
			res := service.Create(config, db)
			c.JSON(http.StatusOK, res)
		}
	}
}

func LoginHandler(
	config *config.Config,
	db *database.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		service := UserLogin{}
		if err := c.ShouldBind(&service); err != nil {
			c.JSON(200, response.Error(response.ErrCodeParameter))
		} else {
			res := service.Login(config, db)
			c.JSON(http.StatusOK, res)
		}
	}
}
