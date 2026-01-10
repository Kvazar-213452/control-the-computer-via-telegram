package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Name_shel = "213452"
	Size_x    = "500"
	Size_y    = "500"
	Time_gif  = 5

	mu sync.Mutex

	lastSentTime = make(map[int64]time.Time)
	delaySeconds = 6 * time.Second
)

func main() {
	botToken := "7663786653:AAH4neWHUNntiQgUhhfFC-txBtuJjbvcii0"

	for {
		err := startBot(botToken)
		if err != nil {
			fmt.Println("error bot:", err)
		}

		fmt.Println("reload")
		time.Sleep(5 * time.Second)
	}
}

func startBot(botToken string) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("errro bot: %w", err)
	}

	bot.Debug = true
	fmt.Printf("bot good %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			text := update.Message.Text

			if text == "/ping" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
				bot.Send(msg)
				continue
			}

			if update.Message.Animation != nil {
				go handleGIF(bot, update.Message)
			}
		}
	}

	return fmt.Errorf("GetUpdatesChan close")
}

func handleGIF(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	userID := msg.From.ID

	if lastTime, ok := lastSentTime[userID]; ok {
		if time.Since(lastTime) < delaySeconds {
			remain := delaySeconds - time.Since(lastTime)
			sec := int(remain.Seconds()) + 1
			reply := fmt.Sprintf("чекай %d c для gif.", sec)
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, reply))
			return
		}
	}

	mu.Lock()
	defer mu.Unlock()

	lastSentTime[userID] = time.Now()

	fileID := msg.Animation.FileID
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		fmt.Println("errro get file:", err)
		return
	}

	url := file.Link(bot.Token)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("errro load file:", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err omsg:", err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	dataPath := "./shell/data.temp"
	err = saveBase64ToFile(encoded, dataPath)
	if err != nil {
		fmt.Println("errro save GIF:", err)
		return
	}

	fmt.Println("GIF save on", dataPath)

	err = runShellApp()
	if err != nil {
		fmt.Println("satrt error shell_web.exe:", err)
	}
}

func saveBase64ToFile(data, path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(data), 0644)
}

func runShellApp() error {
	cmd := exec.Command("./shell_web.exe", Name_shel, Size_x, Size_y)
	cmd.Dir = "./shell"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	fmt.Printf("shell_web.exe start on PID: %d\n", cmd.Process.Pid)

	time.Sleep(time.Duration(Time_gif) * time.Second)

	err = cmd.Process.Kill()
	if err != nil {
		return fmt.Errorf("errro close shell_web.exe: %w", err)
	}

	fmt.Printf("shell_web.exe close PID: %d\n", cmd.Process.Pid)
	return nil
}
