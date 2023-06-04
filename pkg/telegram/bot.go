// Package telegram ...
package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"

	"github.com/alexvlasov182/telegram-bot/pkg/config"
	"github.com/alexvlasov182/telegram-bot/pkg/repository"
)

// Bot function
type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string

	messages config.Messages
}

// NewBot function
func NewBot(
	bot *tgbotapi.BotAPI,
	pocketClient *pocket.Client,
	tr repository.TokenRepository,
	redirectURL string,
	messages config.Messages,
) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, redirectURL: redirectURL, tokenRepository: tr, messages: messages}
}

// Start function
func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
