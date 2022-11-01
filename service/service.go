package service

import (
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Service interface {
	ProcessingResp(update tgbotapi.Update) tgbotapi.MessageConfig
}

type serviceImpl struct {
}

func NewService(rep repository.Repository) *serviceImpl {
	return &serviceImpl{}
}
