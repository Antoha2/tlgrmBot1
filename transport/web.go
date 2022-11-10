package transport

import (
	"github.com/Antoha2/tlgrmBot1/config"
	service "github.com/Antoha2/tlgrmBot1/service/windService"
)

type Transport interface {
}

type webImpl struct {
	config      *config.Config
	windService service.Service
}

func NewWeb(service service.Service, cfg *config.Config) *webImpl {
	return &webImpl{
		windService: service,
		config:      cfg,
	}
}
