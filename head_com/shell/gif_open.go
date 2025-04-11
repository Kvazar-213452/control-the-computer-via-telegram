package shell

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Open_gif(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	var fileID string
	if message.Animation != nil {
		fileID = message.Animation.FileID
	} else {
		log.Println("Помилка: анімація не знайдена")
		return 0
	}

	fileInfo, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		log.Println("Помилка отримання файлу:", err)
		return 0
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, fileInfo.FilePath)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Помилка завантаження файлу:", err)
		return 0
	}
	defer resp.Body.Close()

	videoData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Помилка читання тіла відповіді:", err)
		return 0
	}

	encoded := base64.StdEncoding.EncodeToString(videoData)

	dataPath := "./lib/shell/data.temp"
	err = os.WriteFile(dataPath, []byte(encoded), 0644)
	if err != nil {
		log.Fatalln("Не вдалося записати base64 у файл:", err)
	}

	cmd := exec.Command("./shell_web.exe", "unix", "500", "500")
	cmd.Dir = "./lib/shell"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Помилка при виконанні процесу: %v\n", err)
		} else {
			fmt.Println("Процес завершено успішно.")
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				log.Printf("Помилка при завершенні процесу: %v\n", err)
			} else {
				log.Println("Процес завершено після 5 секунд.")
			}
		}
	}()

	return 1
}
