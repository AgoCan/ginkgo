package server

import (
	"ginkgo/internal/handler/auth"
	"ginkgo/internal/handler/file"
)

func (s *Server) migrate() {
	auth.AutoMigrate(s.DB.DB)
	file.AutoMigrate(s.DB.DB)
}
