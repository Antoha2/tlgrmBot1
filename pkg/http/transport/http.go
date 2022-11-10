package http

import (
	"net/http"

	"github.com/Antoha2/tlgrmBot1/config"
	service "github.com/Antoha2/tlgrmBot1/service/windService"
)

type httpImpl struct {
	service service.Service
	server  *http.Server
	config  *config.Config
}

func NewHTTP(service service.Service, cfg *config.Config) *httpImpl {
	return &httpImpl{
		service: service,
		config:  cfg,
	}
}
