// Package telegram ...
package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

// msg.Text = "You are not authorized! Use command /start"
// msg.Text = "This is hot valid link!"
// msg.Text = "Sorry something wrong, try letter"
func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "unknown error happend")

	switch err {
	case errInvalidURL:
		msg.Text = "This is not valid link!"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "You are not authorized! Use command /start"
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Sorry something wrong, try letter"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
