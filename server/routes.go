package server

import (
	"titan/handlers"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
}
