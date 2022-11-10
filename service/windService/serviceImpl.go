package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/Antoha2/tlgrmBot1/internal/meteo"
	repository "github.com/Antoha2/tlgrmBot1/repository"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

//обработка ответа
func (s *service) ProcessingResp(ctx context.Context, tgMessage tgbotapi.Update) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig
	err := s.UserVerification(context.Background(), tgMessage)
	if err != nil {
		msg = tgbotapi.NewMessage(tgMessage.Message.Chat.ID, "неверное имя пользователя")
		return msg
	}

	if tgMessage.Message.IsCommand() {
		switch tgMessage.Message.Text {
		case "/start":
			msg = tgbotapi.NewMessage(tgMessage.Message.Chat.ID, fmt.Sprintf("Приветствую, %s", UserNameVerification(tgMessage)))
			return msg
		case "/repeat_last_request":
			msg = s.RepeatRequest(context.Background(), tgMessage)
			return msg
		case "/stavropol":
			tgMessage.Message.Text = "ставрополь"
		case "/moskow":
			tgMessage.Message.Text = "москва"
		case "/saint_petersburg":
			tgMessage.Message.Text = "санкт-петербург"
		case "/alma_ata":
			tgMessage.Message.Text = "алма-ата"

		}
	}

	var yaData string
	repMessage := new(repository.RepositoryMessagelist)

	testOk := checkMsgText(tgMessage.Message.Text)
	if !testOk {
		yaData = "некорректный ввод"
		msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, yaData)
		return msg
	}

	coordinates, err := s.gk.GetCoordinates(tgMessage.Message.Text)
	if err != nil {
		log.Println("GetCoordinates() - ", err)
		yaData = "не найдено"
		msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, yaData)
		return msg
	}

	reqCoord := &meteo.Querry{
		Lat:      coordinates.Lat,
		Lon:      coordinates.Lon,
		CityName: coordinates.CityName,
	}

	yaData, err = s.ya.GetWind(reqCoord)
	if err != nil {
		log.Println("GetWinder() - ", err)
		yaData = "некорректный ввод"

	}

	msg = tgbotapi.NewMessage(tgMessage.Message.Chat.ID, yaData)

	repMessage.Text = tgMessage.Message.Text
	repMessage.ChatId = tgMessage.Message.Chat.ID
	repMessage.Response = msg.Text
	repMessage.UserId = tgMessage.Message.From.ID

	err = s.rep.AddMessage(repMessage)
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
		msg := tgbotapi.NewMessage(tgMessage.Message.Chat.ID, "нельзя повторить запрос")
		return msg
	}
	tgMessage.Message.Text = repMessage.Text
	msg := s.ProcessingResp(ctx, tgMessage)

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
func (s *service) UserVerification(ctx context.Context, user tgbotapi.Update) error {
	repUser := &repository.RepositoryUserlist{
		UserId: user.Message.From.ID,
	}

	if !s.rep.UserVerification(repUser) {
		err := s.AddUser(ctx, user)
		if err != nil {
			log.Println("s.rep.AddUser() - ", err)
			return fmt.Errorf("ошибка добавления пользователя")
		}
	}
	return nil
}

//получение истории запросов пользователя
func (s *service) GetHistory(ctx context.Context, userId int) ([]ServMessage, error) {
	repHistory, err := s.rep.GetHistory(userId)
	if err != nil {
		log.Println("s.rep.GetHistory() - ", err)
		return nil, err
	}
	servHistory := make([]ServMessage, len(repHistory))

	for index, msg := range repHistory {
		servHistory[index].ChatId = msg.ChatId
		servHistory[index].MessageId = msg.MessageId
		servHistory[index].Response = msg.Response
		servHistory[index].Text = msg.Text
	}
	return servHistory, nil
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

//
