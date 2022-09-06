package file

import (
	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/response"
)

type FileReq struct {
	
}

func (f *FileReq) Upload(config *config.Config,
	db *database.Client) response.Response {
	
	return response.Success("ok")
}
