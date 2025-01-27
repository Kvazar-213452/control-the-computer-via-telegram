package head_com

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Check_msg_user(msg string, bot *tgbotapi.BotAPI, chatID int64) int {
	if msg == "#help" {
		msg := tgbotapi.NewMessage(chatID, "ok")
		bot.Send(msg)
		return 1
	} else if msg == "/start" {
		msg := tgbotapi.NewMessage(chatID, "cool 213452!")
		bot.Send(msg)
		return 1
	} else if msg == "#off" {
		msg := tgbotapi.NewMessage(chatID, "bot off")
		bot.Send(msg)
		os.Exit(0)

		return 1
	}

	return 0
}
