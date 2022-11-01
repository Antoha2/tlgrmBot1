package service

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (sImpl *serviceImpl) ProcessingResp(update tgbotapi.Update) tgbotapi.MessageConfig {

	UserName := update.Message.From.UserName
	ChatID := update.Message.Chat.ID
	Text := update.Message.Text
	log.Printf("[%s] %d %s \n", UserName, ChatID, Text)
	reply := Text
	msg := tgbotapi.NewMessage(ChatID, reply)

	return msg
}
