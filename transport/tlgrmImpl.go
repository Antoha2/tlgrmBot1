package transport

import (
	"context"

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

		msg := wImpl.windService.ProcessingResp(context.Background(), update)
		bot.Send(msg)
	}
}

/*
func getGismeteo(apiUrl string, offset int) error {

	buf, err := json.Marshal(GismeteoToken)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", GismeteoApi, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return err
	}
	req.Header.Set("X-Gismeteo-Token", GismeteoToken)
	log.Println("req - ", req)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return err
	}

	//resp, err := http.Get(apiUrl)
	/* buf, err := json.Marshal(GismeteoToken)
	if err != nil {
		return err
	}
	var w http.ResponseWriter
	//resp, err := http.Get(GismeteoApi)
	//resp, err := http.Post(GismeteoApi, "application/json", bytes.NewBuffer(buf))
	// if err != nil {
	// 	log.Println("http.Get() - ", err)
	// 	return err
	// }
	defer resp.Body.Close()

	log.Println("resp - ", resp)
	//resp.Header.Set()

	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return err
	}

	//var
	/* restResponse := new(Gismeteo)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return err
	}
	log.Println(restResponse)
	return nil
}

// func respond(botUrl string, update Update) error {
// 	var botMessage BotMessage
// 	botMessage.ChatId = update.Message.Chat.ChatId
// 	botMessage.Text = update.Message.Text
// 	buf, err := json.Marshal(botMessage)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
// 	if err != nil {
// 		log.Println("http.Post() -", err)
// 		return err
// 	}

// 	return nil
// }
*/
