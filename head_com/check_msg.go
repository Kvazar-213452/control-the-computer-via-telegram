package head_com

import (
	"head/head_com/func_unix"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VolumeRequest struct {
	Volume float64 `json:"volume"`
}

func Check_msg_user(msg string, bot *tgbotapi.BotAPI, chatID int64) int {
	if msg == "#help" {
		msg := tgbotapi.NewMessage(chatID, "ok")
		bot.Send(msg)
		return 1
	} else if msg == "/start" {
		msg := tgbotapi.NewMessage(chatID, "viva 213452!")
		bot.Send(msg)
		return 1
	} else if msg == "#off" {
		msg := tgbotapi.NewMessage(chatID, "bot off")
		bot.Send(msg)
		os.Exit(0)

		return 1
	} else if strings.HasPrefix(msg, "#sound") {
		parts := strings.Fields(msg)
		if len(parts) == 2 {
			volume, err := strconv.ParseFloat(parts[1], 64)
			if err != nil || volume < 0 || volume > 100 {
				msg := tgbotapi.NewMessage(chatID, "Невірне значення для гучності (0-100).")
				bot.Send(msg)
				return 1
			}

			func_unix.Post_sound(volume, chatID, bot)

			msg := tgbotapi.NewMessage(chatID, "Гучність змінена на "+strconv.FormatFloat(volume, 'f', 0, 64)+"%")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, "Формат команди: #sound <value> (0-100).")
			bot.Send(msg)
		}
		return 1
	}

	return 0
}
