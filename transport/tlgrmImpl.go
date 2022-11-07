package transport

import (
	"context"
	"log"
	"reflect"

	"github.com/Antoha2/tlgrmBot1/config"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (wImpl *webImpl) StartBot() {

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Println("NewBotAPI() - ", err)
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println("GetUpdatesChan() - ", err)
		panic(err)
	}

	var msg tgbotapi.MessageConfig

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {

			case "/stavropol":
				update.Message.Text = "ставрополь"
				msg = wImpl.windService.ProcessingResp(context.Background(), update)
				bot.Send(msg)
			case "/moskow":
				update.Message.Text = "москва"
				msg = wImpl.windService.ProcessingResp(context.Background(), update)
				bot.Send(msg)
			case "/saint_petersburg":
				update.Message.Text = "санкт-петербург"
				msg = wImpl.windService.ProcessingResp(context.Background(), update)
				bot.Send(msg)
			case "/alma_ata":
				update.Message.Text = "алма-ата"
				msg = wImpl.windService.ProcessingResp(context.Background(), update)
				bot.Send(msg)
			case "/repeat_last_request":
				msg = wImpl.windService.RepeatRequest(context.Background(), update)
				bot.Send(msg)
			default:
				msg = wImpl.windService.ProcessingResp(context.Background(), update)
				bot.Send(msg)
			}
		}
	}
}
