package service

import (
	"context"

	"github.com/Antoha2/tlgrmBot1/internal/geokoder"
	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Service interface {
	ProcessingResp(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig
	RepeatRequest(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig //tgbotapi.Update
	AddUser(ctx context.Context, update tgbotapi.Update) error                        //string
	UserVerification(ctx context.Context, update tgbotapi.Update)                     //(bool, error)
}

type service struct {
	rep repository.Repository
	ya  meteo.GetWinder
	gis meteo.GetWinder
	gk  geokoder.GeokoderIn
}

func NewService(rep repository.Repository, yandex meteo.GetWinder, gismeteo meteo.GetWinder, geokoder geokoder.GeokoderIn) *service {
	return &service{
		rep: rep,
		ya:  yandex,
		gis: gismeteo,
		gk:  geokoder,
	}
}
