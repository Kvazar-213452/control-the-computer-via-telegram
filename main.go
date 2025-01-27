package main

import (
	"fmt"
	"head/head_com"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Зчитуємо токен з файлу
	token, err := readTokenFromFile("data.unix")
	if err != nil {
		log.Fatalf("Помилка при читанні токена: %v", err)
	}

	// Створюємо бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Помилка при створенні бота: %v", err)
	}

	// Включаємо налагодження
	bot.Debug = true

	log.Printf("Авторизовано як %s", bot.Self.UserName)

	// Оновлення бота
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Канал для обробки сигналів завершення програми
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Основний цикл для обробки оновлень
	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				// Перевіряємо повідомлення за допомогою функції з head_com
				ver := head_com.Check_msg_user(update.Message.Text, bot, update.Message.Chat.ID)
				if ver == 0 {
					// Якщо не знайдено команду, відправляємо повідомлення
					reply := "govno napusav: " + update.Message.Text
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					bot.Send(msg)
				}
			}
		case <-sigChan:
			log.Println("Бот завершив роботу.")
			return
		}
	}
}

// Читання токену з файлу
func readTokenFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("не вдалося прочитати файл %s: %v", filename, err)
	}
	return string(data), nil
}
