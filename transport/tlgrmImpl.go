package transport

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/* var botToken string = "5610958233:AAF7iMAJBZdrEJdJsQF9GJ_Oxcm7S_TiNio"
var botApi string = "https://api.telegram.org/bot"

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
} */

func StartBot() {

	botUrl := botApi + botToken
	offset := 0

	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("getUpdates() -", err)
		}
		for _, update := range updates {
			err := respond(botUrl, update)
			if err != nil {
				log.Println("respond() -", err)
			}
			offset = update.UpdateId + 1

		}
		log.Println(updates)

	}

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
