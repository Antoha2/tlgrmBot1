package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (wImpl *webImpl) StartBot() {

	bot, err := tgbotapi.NewBotAPI(botToken) //(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		bot.Send(wImpl.service.ProcessingResp(context.Background(), update))

	}
}

func (wImpl *webImpl) StartWindRequest() {

	GismeteoUrl := GismeteoApi + GismeteoToken
	offset := 0

	updates, err := getUpdates(GismeteoUrl, offset)
	if err != nil {
		log.Println("getUpdates() -", err)
	}
	for _, update := range updates {
		err := respond(GismeteoUrl, update)
		if err != nil {
			log.Println("respond() -", err)
		}
		offset = update.UpdateId + 1

	}
	log.Println(updates)

}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		log.Println("http.Get() - ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Println("http.Post() -", err)
		return err
	}

	return nil
}
