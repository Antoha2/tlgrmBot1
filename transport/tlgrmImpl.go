package transport

import (
	"context"
	"log"

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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := wImpl.windService.ProcessingResp(context.Background(), update)
		bot.Send(msg)
	}
}
