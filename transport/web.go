package transport

import (
	"github.com/Antoha2/tlgrmBot1/windService/service"
)

type Transport interface {
}

type webImpl struct {
	service service.Service
}

func NewWeb(service service.Service) *webImpl {
	return &webImpl{
		service: service,
	}
}
