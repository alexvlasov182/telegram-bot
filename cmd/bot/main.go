package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"github.com/alexvlasov182/telegram-bot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6149137591:AAGIyIRDeaL3HfHmmKOhUpMF0dW8QmqXkxQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true


	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
