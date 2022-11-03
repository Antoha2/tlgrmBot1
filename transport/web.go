package transport

import (
	service "github.com/Antoha2/tlgrmBot1/service/windService"
)

type Transport interface {
}

type webImpl struct {
	windService service.Service
}

func NewWeb(service service.Service) *webImpl {
	return &webImpl{
		windService: service,
	}
}
