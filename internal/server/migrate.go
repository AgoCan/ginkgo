package server

import "ginkgo/internal/handler/auth"

func (s *Server) migrate() {
	auth.AutoMigrate(s.DB.DB)
}
