package func_unix

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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
		fmt.Println("error:", err)
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
		fmt.Println("error:", err)
		return
	}

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("error")
		} else {
			fmt.Println("error:", err)
		}
		return
	}

	cmd := exec.Command("cmd", "/C", "start", filePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err = cmd.Start()
	if err != nil {
		fmt.Println("error:", err)
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
		fmt.Println("Error putting system to sleep:", err)
	}
}

func Shutdown_pc() {
	cmd := exec.Command("shutdown", "/s", "/t", "1")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error shutting down the system:", err)
	}
}

func Reboot_pc() {
	cmd := exec.Command("shutdown", "/r", "/t", "1")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error rebooting the system:", err)
	}
}

func Disable_WiFi() error {
	cmd := exec.Command("netsh", "interface", "set", "interface", "Wi-Fi", "disabled")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error Wi-Fi: %v", err)
	}
	return nil
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
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("good")
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
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("good")
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
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("good")
	}
}

func Key_unix(key string) {
	cmd := exec.Command("./set_key.exe", key)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("good")
	}
}

func Speench_text(text string, chatID int64, bot *tgbotapi.BotAPI) {
	cmd := exec.Command("./main.exe", text)

	cmd.Dir = "./lib"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println("good")
	}

	f, err := os.Open("./lib/output.mp3")
	if err != nil {
		log.Fatalf("error: %v", err)
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
