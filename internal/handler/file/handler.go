package file

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/file"
	"ginkgo/internal/pkg/response"
)

func UploadrHandler(
	config *config.Config,
	db *database.Client,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		service := FileReq{}
		c.Request.ParseMultipartForm(4 << 20)

		mFile, mFileHeader, err := c.Request.FormFile("file")

		if err != nil {
			c.JSON(200, response.Error(response.ErrCodeParameter))
		}
		defer mFile.Close()

		if err := file.CreateDir(config.File.Path); err != nil {
			c.JSON(200, response.Error(response.CreateDirectorErr))
		}

		filePath := config.File.Path + "/" + mFileHeader.Filename
		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			c.JSON(200, response.Error(response.CreateFileErr))
		}
		defer destFile.Close()
		io.Copy(destFile, mFile)

		res := service.Upload(config, db)
		c.JSON(http.StatusOK, res)

	}
}
