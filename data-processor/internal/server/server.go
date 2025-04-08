package server

import (
	"data-processor/internal/transport"
)

type Server struct {
	transport.UnimplementedFileServiceServer
	service service
}

func New(service service) *Server {
	return &Server{service: service}
}
