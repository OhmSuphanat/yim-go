package handler

import (
	"yim-go/api/internal/core/port"
	"yim-go/shared/gen/protobuf"
)

type handler struct {
	svc port.Service
	protobuf.UnimplementedBookServiceServer // Embed this struct
}



func New(svc port.Service) protobuf.BookServiceServer {
	return &handler{
		svc: svc,
	}
}
