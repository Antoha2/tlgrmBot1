package service

import (
	"context"

	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Service interface {
	ProcessingResp(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig
}

type serviceImpl struct {
	repository repository.Repository
}

func NewService(rep repository.Repository) *serviceImpl {
	return &serviceImpl{
		repository: rep,
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
