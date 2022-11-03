package service

import (
	"context"
	"log"

	"github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (sImpl *serviceImpl) ProcessingResp(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	repMessage := new(repository.RepositoryMessage)

	//Chat := update.Message.Chat.ID
	repMessage.UserName = tgMessage.Message.From.UserName
	repMessage.Text = tgMessage.Message.Text
	repMessage.Chat.ChatId = tgMessage.Message.Chat.ID

	err := sImpl.repository.AddMessage(repMessage)
	if err != nil {
		log.Println(err)
	}

	//contact := update.Message.Contact.UserID
	//log.Println(sm)
	//log.Printf("[%s] %d %s \n", UserName, ChatID, Text)
	//reply := sm.Text
	msg := tgbotapi.NewMessage(repMessage.Chat.ChatId, repMessage.Text)

	return msg
}
