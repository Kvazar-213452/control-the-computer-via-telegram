package func_unix

import (
	"bytes"
	"encoding/json"
	"head/head_com/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Volume_request struct {
	Volume float64 `json:"volume"`
}

type Volume_request_1 struct {
	Volume string `json:"volume"`
}

func Post_request(endpoint string, requestBody []byte, chatID int64, bot *tgbotapi.BotAPI) {
	url := config.Server_url_back + endpoint

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Error during HTTP request: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Server is unreachable")
		bot.Send(msg)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(body))
		msg := tgbotapi.NewMessage(chatID, "error")
		bot.Send(msg)
		return
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error decoding response: %v", err)
	}
}

func Post_sound(volume float64, chatID int64, bot *tgbotapi.BotAPI) {
	Volume_request := Volume_request{Volume: volume}
	requestBody, err := json.Marshal(Volume_request)
	if err != nil {
		log.Printf("Error JSON marshaling: %v", err)
		return
	}
	Post_request("set_volume", requestBody, chatID, bot)
}

func Post_set_mouse_speed(volume float64, chatID int64, bot *tgbotapi.BotAPI) {
	Volume_request := Volume_request{Volume: volume}
	requestBody, err := json.Marshal(Volume_request)
	if err != nil {
		log.Printf("Error JSON marshaling: %v", err)
		return
	}
	Post_request("set_mouse_speed", requestBody, chatID, bot)
}

func Post_change_screen_brightness(volume float64, chatID int64, bot *tgbotapi.BotAPI) {
	Volume_request := Volume_request{Volume: volume}
	requestBody, err := json.Marshal(Volume_request)
	if err != nil {
		log.Printf("Error JSON marshaling: %v", err)
		return
	}
	Post_request("change_screen_brightness", requestBody, chatID, bot)
}

func Post_key_unix(volume string, chatID int64, bot *tgbotapi.BotAPI) {
	Volume_request := Volume_request_1{Volume: volume}
	requestBody, err := json.Marshal(Volume_request)
	if err != nil {
		log.Printf("Error JSON marshaling: %v", err)
		return
	}
	Post_request("key_unix", requestBody, chatID, bot)
}

func Post_speench_text(volume string, chatID int64, bot *tgbotapi.BotAPI) {
	Volume_request := Volume_request_1{Volume: volume}
	requestBody, err := json.Marshal(Volume_request)
	if err != nil {
		log.Printf("Error JSON marshaling: %v", err)
		return
	}
	Post_request("speench_text", requestBody, chatID, bot)

	done := make(chan struct{})
	go start_music(chatID, bot, done)
	<-done
}

func Post_close() {
	url := config.Server_url_back + "close"

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Error during HTTP request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK HTTP status: %s", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(body))
		return
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error decoding response: %v", err)
	}
}

// unuisx_misuc// unuisx_misuc// unuisx_misuc
// unuisx_misuc// unuisx_misuc// unuisx_misuc
// unuisx_misuc// unuisx_misuc// unuisx_misuc

func start_music(chatID int64, bot *tgbotapi.BotAPI, done chan struct{}) {
	f, err := os.Open("data_use/uploaded_output.mp3")
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid nema music"))
		close(done)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid format"))
		close(done)
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid cod gavno"))
		close(done)
		return
	}

	speaker.Play(streamer)

	select {}

	close(done)
}
