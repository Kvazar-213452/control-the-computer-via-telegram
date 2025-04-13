package main

import (
	"head/head_com"
	"head/head_com/config_func"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	go func() {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Debug

		bot.Debug = config_func.Debug

		log.Printf("active on %s", bot.Self.UserName)

		// updata
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case update := <-updates:
				if update.Message != nil {
					ver := head_com.Check_msg_user(bot, update.Message.Chat.ID, update.Message)
					if ver == 0 {
						reply := "govno napusav: " + update.Message.Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
						bot.Send(msg)
					}
				}
			case <-sigChan:
				return
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	config_func.Log_append("end app")
}
