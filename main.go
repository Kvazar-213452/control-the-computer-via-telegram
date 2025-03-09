package main

import (
	"fmt"
	"head/head_com"
	"head/head_com/config"
	"head/head_com/shell"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token, err := head_com.Read_file("data/data.unix")
	if err != nil {
		log.Fatalf("error %v", err)
	}

	go func() {
		http.HandleFunc("/upload_mp3", shell.Upload_MP3)
		if err := http.ListenAndServe(config.Port, nil); err != nil {
			fmt.Println("error", err)
		}
	}()

	go func() {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Debug
		bot.Debug = true

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
					ver := head_com.Check_msg_user(update.Message.Text, bot, update.Message.Chat.ID)
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
	fmt.Println("end")
}
