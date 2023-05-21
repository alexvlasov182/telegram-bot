// Main package ...
package main

import (
	"log"

	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"

	"github.com/alexvlasov182/telegram-bot/pkg/repository"
	"github.com/alexvlasov182/telegram-bot/pkg/repository/boltdb"
	"github.com/alexvlasov182/telegram-bot/pkg/server"
	"github.com/alexvlasov182/telegram-bot/pkg/telegram"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")

	authorizationServer := server.NewAuthorizationServer(
		pocketClient,
		tokenRepository,
		"https://t.me/Pokcet_bot",
	)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
