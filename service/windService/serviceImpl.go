package service

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (s *service) ProcessingResp(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	s.UserVerification(ctx, tgMessage)

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
				////////////////////////////////////////////////////////
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
		log.Println("s.rep.AddMessage() - ", err)
	}
	return msg
}

//повтор крайнего запроса пользователя
func (s *service) RepeatRequest(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	repMessage, err := s.rep.RepeatMessage(tgMessage.Message.Chat.ID)
	if err != nil {
		log.Println("s.rep.RepeatMessage() - ", err)
		///////////////////////////////////////////////////
	}
	tgMessage.Message.Text = repMessage.Text
	msg := s.ProcessingResp(ctx, tgMessage)

	//text := fmt.Sprintf("запрос: %s \n ответ: %s", repMessage.Text, repMessage.Response)
	//msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, text)

	return msg
}

//добавление нового пользователя
func (s *service) AddUser(ctx context.Context, tgMessage tgbotapi.Update) error {
	repUser := &repository.RepositoryUserlist{
		UserId:   tgMessage.Message.From.ID,
		UserName: tgMessage.Message.From.UserName,
		Add_date: string(time.Now().Format(time.RFC3339)),
	}

	err := s.rep.AddUser(repUser)
	if err != nil {
		log.Println("s.rep.AddUser() - ", err)
		return err
	}
	return nil
}

//проверка пользователя
func (s *service) UserVerification(ctx context.Context, user tgbotapi.Update) {
	repUser := &repository.RepositoryUserlist{
		UserId: user.Message.From.ID,
	}

	if !s.rep.UserVerification(repUser) {
		err := s.AddUser(ctx, user)
		if err != nil {
			log.Println("s.rep.AddUser() - ", err)
			//////////////////////////////////////////////
		}
	}
}

//проверка введенного текста
func checkMsgText(msg string) bool {
	m, _ := regexp.MatchString("^[а-яA-Я]", msg)
	return m
}

//проверка на наличие UserName
func UserNameVerification(user tgbotapi.Update) string {
	if user.Message.From.UserName == "" {
		return strconv.Itoa(user.Message.From.ID)
	}
	return user.Message.From.UserName
}
