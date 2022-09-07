package file

import (
	"mime/multipart"

	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/response"
)

type FileReq struct {
}

func (f *FileReq) Upload(config *config.Config,
	db *database.Client, header multipart.FileHeader, filePath string) response.Response {
	var fileModel File
	fileModel.Name = header.Filename
	fileModel.Path = filePath
	fileModel.Size = header.Size
	if err := db.DB.Create(&fileModel).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Success("ok")
}
