package adapter

import "yim-go/api/internal/core/port"

type adapter struct {
}

func New() port.Adapter {
	return &adapter{}
}