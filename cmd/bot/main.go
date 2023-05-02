package main

import (
	"log"

	"github.com/alexvlasov182/telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6149137591:AAGIyIRDeaL3HfHmmKOhUpMF0dW8QmqXkxQ")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("106868-68946fd24af1ea54a6051e7")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
