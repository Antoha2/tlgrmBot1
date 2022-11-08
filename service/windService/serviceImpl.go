package service

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (s *service) ProcessingResp(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	var yaData string
	repMessage := new(repository.RepositoryMessagelist)

	testOk := checkMsgText(tgMessage.Message.Text)

	if testOk {
		coordinates, err := s.gk.GetCoordinates(tgMessage.Message.Text)
		if err != nil {
			log.Println("GetCoordinates() - ", err)
			yaData = "не найдено"
		} else {

			reqCoord := &meteo.Querry{
				Lat:      coordinates.Lat,
				Lon:      coordinates.Lon,
				CityName: coordinates.CityName,
			}

			yaData, err = s.ya.GetWind(reqCoord)
			if err != nil {
				log.Println("GetWinder() - ", err)
			}
		}
	} else {
		yaData = "некорректный ввод"
	}
	msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, yaData)
	//Chat := update.Message.Chat.ID

	//repMessage.UserName = tgMessage.Message.From.UserName
	repMessage.Text = tgMessage.Message.Text
	repMessage.ChatId = tgMessage.Message.Chat.ID
	repMessage.Response = msg.Text
	repMessage.UserId = tgMessage.Message.From.ID //tgMessage.Message.Contact.UserID

	err := s.rep.AddMessage(repMessage)
	if err != nil {
		log.Println(err)
	}
	return msg
}

func (s *service) RepeatRequest(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	repMessage, err := s.rep.RepeatMessage(tgMessage.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		//return nil
	}

	text := fmt.Sprintf("запрос: %s \n ответ: %s", repMessage.Text, repMessage.Response)
	msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, text)

	// msg.Message.From.UserName = repMessage.UserName
	// msg.Message.Text = fmt.Sprintf("запрос: %s \n ответ: %s", repMessage.Text, repMessage.Response)
	// msg.Message.Chat.ID = repMessage.Chat.ChatId

	return msg
}

func checkMsgText(msg string) bool {
	m, _ := regexp.MatchString("^[а-яA-Я]", msg)
	return m
}
