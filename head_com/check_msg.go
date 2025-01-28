package head_com

import (
	"encoding/json"
	"fmt"
	"head/head_com/func_unix"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VolumeRequest struct {
	Volume float64 `json:"volume"`
}

var let_start = 1

func Check_msg_user(msg string, bot *tgbotapi.BotAPI, chatID int64) int {
	if msg == "#bot1" {
		let_start = 1
		bot.Send(tgbotapi.NewMessage(chatID, "set_1 unix"))
		return 1
	} else if msg == "#bot0" {
		let_start = 0
		bot.Send(tgbotapi.NewMessage(chatID, "set0 unix"))
		return 1
	} else if let_start == 1 {
		return 1
	} else if msg == "#help" {
		text, err := Read_file("data/help.unix")
		if err != nil {
			log.Fatalf("error %v", err)
		}

		bot.Send(tgbotapi.NewMessage(chatID, text))
		return 1
	} else if msg == "/start" {
		bot.Send(tgbotapi.NewMessage(chatID, "viva 213452!"))
		return 1
	} else if msg == "#off" {
		bot.Send(tgbotapi.NewMessage(chatID, "bot off"))
		os.Exit(0)

		return 1
	} else if strings.HasPrefix(msg, "#sound") {
		parts := strings.Fields(msg)
		if len(parts) == 2 {
			volume, err := strconv.ParseFloat(parts[1], 64)
			if err != nil || volume < 0 || volume > 100 {
				bot.Send(tgbotapi.NewMessage(chatID, "diapazon (0-100)."))
				return 1
			}

			func_unix.Post_sound(volume, chatID, bot)

			bot.Send(tgbotapi.NewMessage(chatID, "change sound "+strconv.FormatFloat(volume, 'f', 0, 64)+"%"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "format command #sound <value> (0-100)."))
		}
		return 1
	} else if strings.HasPrefix(msg, "#mouse") {
		parts := strings.Fields(msg)
		if len(parts) == 2 {
			volume, err := strconv.ParseFloat(parts[1], 64)
			if err != nil || volume < 0 || volume > 100 {
				bot.Send(tgbotapi.NewMessage(chatID, "diapazon (0-100)."))
				return 1
			}

			func_unix.Post_set_mouse_speed(volume, chatID, bot)

			bot.Send(tgbotapi.NewMessage(chatID, "change mouse "+strconv.FormatFloat(volume, 'f', 0, 64)+"%"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "format command #mouse <value> (0-100)."))
		}
		return 1
	} else if strings.HasPrefix(msg, "#screen") {
		parts := strings.Fields(msg)
		if len(parts) == 2 {
			volume, err := strconv.ParseFloat(parts[1], 64)
			if err != nil || volume < 0 || volume > 100 {
				bot.Send(tgbotapi.NewMessage(chatID, "diapazon (0-100)."))
				return 1
			}

			func_unix.Post_change_screen_brightness(volume, chatID, bot)

			bot.Send(tgbotapi.NewMessage(chatID, "change screen "+strconv.FormatFloat(volume, 'f', 0, 64)+"%"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "format command #screen <value> (0-100)."))
		}
		return 1
	} else if strings.HasPrefix(msg, "#music") {
		parts := strings.Fields(msg)
		volume, err := strconv.ParseFloat(parts[1], 64)
		if err != nil || volume < 0 || volume > 100 {
			bot.Send(tgbotapi.NewMessage(chatID, "diapazon (0-100)."))
			return 0
		}

		go func_unix.Start_music(int(volume), chatID, bot)

		bot.Send(tgbotapi.NewMessage(chatID, "start music"))
		return 1
	} else if msg == "#sleep" {
		func_unix.Sleep_pc()
		bot.Send(tgbotapi.NewMessage(chatID, "sleep"))
		return 1
	} else if msg == "#shutdown" {
		func_unix.Shutdown_pc()
		bot.Send(tgbotapi.NewMessage(chatID, "Shutdown"))
		return 1
	} else if msg == "#reboot" {
		func_unix.Reboot_pc()
		bot.Send(tgbotapi.NewMessage(chatID, "reboot"))
		return 1
	} else if strings.HasPrefix(msg, "#key") {
		parts := strings.Fields(msg)

		args := parts[1:]

		argsJSON, err := json.Marshal(args)
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return 1
		}

		func_unix.Post_key_unix(string(argsJSON), chatID, bot)

		bot.Send(tgbotapi.NewMessage(chatID, "start key"))
		return 1
	} else if strings.HasPrefix(msg, "#text") {
		parts := strings.Fields(msg)

		if len(parts) > 1 {
			text := strings.Join(parts[1:], " ")

			go func_unix.Post_speench_text(text, chatID, bot)

			bot.Send(tgbotapi.NewMessage(chatID, "start speech "+text))

			return 1
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "de text?"))
			return 1
		}
	} else if strings.HasPrefix(msg, "#foto") {
		parts := strings.Fields(msg)
		volume, err := strconv.ParseFloat(parts[1], 64)
		if err != nil || volume < 0 || volume > 100 {
			bot.Send(tgbotapi.NewMessage(chatID, "diapazon (0-100)."))
			return 0
		}

		func_unix.Open_foto(int(volume), chatID, bot)

		bot.Send(tgbotapi.NewMessage(chatID, "start music"))
		return 1
	}

	return 0
}
