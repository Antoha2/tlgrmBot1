package service

import (
	"context"
	"log"
	"regexp"

	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (s *service) ProcessingResp(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	var yaData string

	testOk := checkMsgText(tgMessage.Message.Text)
	log.Println("!!!!!!!!!!!!!!!! test - ", testOk)

	if testOk {
		coordinates, err := s.gk.GetCoordinates(tgMessage.Message.Text)
		if err != nil {
			log.Println("GetCoordinates() - ", err)
		}

		reqCoord := &meteo.Querry{
			Lat:      coordinates.Lat,
			Lon:      coordinates.Lon,
			CityName: coordinates.CityName,
		}

		log.Println("!!!!!!!!!!!!! - ", reqCoord)

		yaData, err = s.ya.GetWind(reqCoord)
		if err != nil {
			log.Println(err)
		}
	} else {
		yaData = "некорректный ввод"
	}

	repMessage := new(repository.RepositoryMessage)

	//Chat := update.Message.Chat.ID
	repMessage.UserName = tgMessage.Message.From.UserName
	repMessage.Text = tgMessage.Message.Text
	repMessage.Chat.ChatId = tgMessage.Message.Chat.ID

	err := s.rep.AddMessage(repMessage)
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(repMessage.Chat.ChatId, yaData)
	return msg
}

func checkMsgText(msg string) bool {
	//m, _ := regexp.MatchString("^[a-zA-Z]", msg)
	m, _ := regexp.MatchString("^[а-яA-Я]", msg)
	return m
}

// func (sImpl *serviceImpl) StartWindRequest() {

// 	sImpl.clientWeather.GetWind()

// 	sImpl.repository.AddMessage()
// 	err :=  sImpl.clientWeather   //getYandex(YandexUrl, offset)
// 	if err != nil {
// 	log.Println("getUpdates() -", err)
// 	}
// 	for _, update := range updates {
// 		err := respond(GismeteoUrl, update)
// 		if err != nil {
// 			log.Println("respond() -", err)
// 		}
// 		offset = update.UpdateId + 1

// 	}
// 	log.Println(updates)

// }
