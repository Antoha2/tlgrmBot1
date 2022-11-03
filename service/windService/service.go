package service

import (
	"context"

	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Service interface {
	ProcessingResp(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig
}

type serviceImpl struct {
	rep repository.Repository
	ya  meteo.GetWinder
	gis meteo.GetWinder
	//clientWeather meteo.GetWinder
}

func NewService(rep repository.Repository, yandex meteo.GetWinder, gismeteo meteo.GetWinder) *serviceImpl {
	return &serviceImpl{
		rep: rep,
		ya:  yandex,
		gis: gismeteo,
		//clientWeather: meteo,
	}
}

type ServiceMessage struct {
	UserName string `json:"user_name"`
	Chat     chat   `json:"chat"`
	Text     string `json:"text"`
}

type chat struct {
	ChatId int64 `json:"id"`
}
