package func_unix

import (
	"encoding/json"
	"fmt"
	"head/head_com/config_func"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Get_file_key(key int, file_1 string) (string, error) {
	file, err := os.Open(file_1)
	if err != nil {
		return "", fmt.Errorf("error open the door: %v", err)
	}
	defer file.Close()

	var musicMap map[string]string

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&musicMap)
	if err != nil {
		return "", fmt.Errorf("error json: %v", err)
	}

	keyStr := fmt.Sprintf("%d", key)

	if value, exists := musicMap[keyStr]; exists {
		return value, nil
	}

	return "", fmt.Errorf("none key")
}

// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func
// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func
// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func// unix_func

func Start_music(volume int, chatID int64, bot *tgbotapi.BotAPI) {
	filePath, err := Get_file_key(volume, "data/music.json")
	if err != nil {
		config_func.Log_append("errro start music")
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid nema music"))
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid format"))
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalid cod gavno"))
		return
	}

	speaker.Play(streamer)

	select {}
}

func Open_foto(volume int, chatID int64, bot *tgbotapi.BotAPI) {
	filePath, err := Get_file_key(volume, "data/foto.json")
	if err != nil {
		config_func.Log_append("errro open foto")
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			config_func.Log_append("errro open foto")
		} else {
			config_func.Log_append("errro open foto")
		}
		return
	}

	cmd := exec.Command("cmd", "/C", "start", filePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err = cmd.Start()
	if err != nil {
		config_func.Log_append("errro open foto")
	} else {
		fmt.Println("good")
	}
}

// pc_func// pc_func
// pc_func// pc_func
// pc_func// pc_func

func Sleep_pc() {
	cmd := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "Sleep")
	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error putting system to sleep")
	}
}

func Shutdown_pc() {
	cmd := exec.Command("shutdown", "/s", "/t", "1")
	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error shutting down the system")
	}
}

func Reboot_pc() {
	cmd := exec.Command("shutdown", "/r", "/t", "1")
	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error rebooting the system")
	}
}

// use lib// use lib// use lib// use lib// use lib// use lib// use lib
// use lib// use lib// use lib// use lib// use lib// use lib// use lib
// use lib// use lib// use lib// use lib// use lib// use lib// use lib

func Close_window() {
	cmd := exec.Command("./close.exe")

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error clone window")
	} else {
		config_func.Log_append("close window")
	}
}

func Sound(volume float64) {
	volumeStr := strconv.FormatFloat(volume, 'f', -1, 64)

	cmd := exec.Command("./sound_volume.exe", volumeStr)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error sound")
	} else {
		config_func.Log_append("good sound change")
	}
}

func Mouse(volume float64) {
	volumeStr := strconv.FormatFloat(volume, 'f', -1, 64)

	cmd := exec.Command("./mouse_speed.exe", volumeStr)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		config_func.Log_append("Error mouse spead")
	} else {
		config_func.Log_append("mouse spead change")
	}
}

func Key_unix(key string) {
	cmd := exec.Command("./set_key.exe", key)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		config_func.Log_append("error key")
	} else {
		config_func.Log_append("key start")
	}
}

func Speench_text(text string, chatID int64, bot *tgbotapi.BotAPI) {
	cmd := exec.Command("./main.exe", text)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		config_func.Log_append("srrro lib text")
	} else {
		config_func.Log_append("good text lib")
	}

	f, err := os.Open("./lib/output.mp3")
	if err != nil {
		config_func.Log_append("mp3 error")
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "invalib MP3"))
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "error sound"))
		return
	}

	speaker.Play(streamer)

	bot.Send(tgbotapi.NewMessage(chatID, "play text"))

	select {}
}

func Bg_window(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	parts := strings.Fields(message.Caption)
	if len(parts) <= 1 {
		return 0
	}

	key := strings.Join(parts[1:], " ")
	saveDir := "data_use/temp"

	os.MkdirAll(saveDir, os.ModePerm)

	fileID := message.Photo[len(message.Photo)-1].FileID
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		config_func.Log_append("error get file for bg")
		return 0
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

	filename := fmt.Sprintf("%x-%d.jpg", config_func.Md5hash(key), time.Now().UnixMilli())
	savePath := filepath.Join(saveDir, filename)

	resp, err := http.Get(url)
	if err != nil {
		config_func.Log_append("error get bg")
		return 0
	}
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		config_func.Log_append("error create file bg")
		return 0
	}
	defer out.Close()
	io.Copy(out, resp.Body)

	dir, _ := os.Getwd()
	phat := filepath.Join(dir, saveDir, filename)

	cmd := exec.Command("./set_bg.exe", phat)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		config_func.Log_append("lib error bg")
	} else {
		config_func.Log_append("lib good bg")
	}

	return 1
}
