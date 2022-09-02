package file

import (
	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadrHandler(
	config *config.Config,
	db *database.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		service := FileReq{}
		if err := c.ShouldBind(&service); err != nil {
			c.JSON(200, response.Error(response.ErrCodeParameter))
		} else {
			res := service.Upload(config, db)
			c.JSON(http.StatusOK, res)
		}
	}
}
