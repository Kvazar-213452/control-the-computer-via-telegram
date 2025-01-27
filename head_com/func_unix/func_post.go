package func_unix

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VolumeRequest struct {
	Volume float64 `json:"volume"`
}

func Post_sound(volume float64, chatID int64, bot *tgbotapi.BotAPI) {
	volumeRequest := VolumeRequest{Volume: volume}

	requestBody, err := json.Marshal(volumeRequest)
	if err != nil {
		log.Fatalf("Error JSON marshaling: %v", err)
	}

	url := "http://127.0.0.1:5000/set_volume"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error during HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	msg := tgbotapi.NewMessage(chatID, "Гучність успішно змінена!")
	bot.Send(msg)
}
