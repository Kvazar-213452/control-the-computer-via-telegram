package shell

import (
	"encoding/base64"
	"fmt"
	"head/head_com/config_func"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ⣿⣿⣿⣿⠛⠛⠉⠄⠁⠄⠄⠉⠛⢿⣿⣿⣿⣿⣿⣿⣿
// ⣿⣿⡟⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿
// ⣿⣿⡇⠄⠄⠄⠐⠄⠄⠄⠄⠄⠄⠄⠠⣿⣿⣿⣿⣿⣿
// ⣿⣿⡇⠄⢀⡀⠠⠃⡐⡀⠠⣶⠄⠄⢀⣿⣿⣿⣿⣿⣿
// ⣿⣿⣶⠄⠰⣤⣕⣿⣾⡇⠄⢛⠃⠄⢈⣿⣿⣿⣿⣿⣿
// ⣿⣿⣿⡇⢀⣻⠟⣻⣿⡇⠄⠧⠄⢀⣾⣿⣿⣿⣿⣿⣿
// ⣿⣿⣿⣟⢸⣻⣭⡙⢄⢀⠄⠄⠄⠈⢹⣯⣿⣿⣿⣿⣿
// ⣿⣿⣿⣭⣿⣿⣿⣧⢸⠄⠄⠄⠄⠄⠈⢸⣿⣿⣿⣿⣿
// ⣿⣿⣿⣼⣿⣿⣿⣽⠘⡄⠄⠄⠄⠄⢀⠸⣿⣿⣿⣿⣿
// ⡿⣿⣳⣿⣿⣿⣿⣿⠄⠓⠦⠤⠤⠤⠼⢸⣿⣿⣿⣿⣿
// ⡹⣧⣿⣿⣿⠿⣿⣿⣿⣿⣿⣿⣿⢇⣓⣾⣿⣿⣿⣿⣿
// ⡞⣸⣿⣿⢏⣼⣶⣶⣶⣶⣤⣶⡤⠐⣿⣿⣿⣿⣿⣿⣿
// ⣯⣽⣛⠅⣾⣿⣿⣿⣿⣿⡽⣿⣧⡸⢿⣿⣿⣿⣿⣿⣿
// ⣿⣿⣿⡷⠹⠛⠉⠁⠄⠄⠄⠄⠄⠄⠐⠛⠻⣿⣿⣿⣿
// ⣿⣿⣿⠃⠄⠄⠄⠄⠄⣠⣤⣤⣤⡄⢤⣤⣤⣤⡘⠻⣿
// ⣿⣿⡟⠄⠄⣀⣤⣶⣿⣿⣿⣿⣿⣿⣆⢻⣿⣿⣿⡎⠝
// ⣿⡏⠄⢀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡎⣿⣿⣿⣿⠐
// ⣿⡏⣲⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢇⣿⣿⣿⡟⣼
// ⣿⡠⠜⣿⣿⣿⣿⣟⡛⠿⠿⠿⠿⠟⠃⠾⠿⢟⡋⢶⣿
// ⣿⣧⣄⠙⢿⣿⣿⣿⣿⣿⣷⣦⡀⢰⣾⣿⣿⡿⢣⣿⣿
// ⣿⣿⣿⠂⣷⣶⣬⣭⣭⣭⣭⣵⢰⣴⣤⣤⣶⡾⢐⣿⣿
// ⣿⣿⣿⣷⡘⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⢃⣼⣿⣿

func Open_gif(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	var fileID string
	if message.Animation != nil {
		fileID = message.Animation.FileID
	} else {
		config_func.Log_append("gif none")
		return 0
	}

	fileInfo, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		config_func.Log_append("fiel gif none")
		return 0
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, fileInfo.FilePath)

	resp, err := http.Get(url)
	if err != nil {
		config_func.Log_append("error dwn gif")
		return 0
	}
	defer resp.Body.Close()

	videoData, err := io.ReadAll(resp.Body)
	if err != nil {
		config_func.Log_append("error")
		return 0
	}

	encoded := base64.StdEncoding.EncodeToString(videoData)

	dataPath := "./lib/shell/data.temp"
	err = os.WriteFile(dataPath, []byte(encoded), 0644)
	if err != nil {
		config_func.Log_append("base64 gif error")
	}

	cmd := exec.Command("./shell_web.exe", config_func.Name_shel, config_func.Size_x, config_func.Size_y)
	cmd.Dir = "./lib/shell"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := cmd.Run()
		if err != nil {
			config_func.Log_append("error gif")
		} else {
			config_func.Log_append("end work gif")
		}
	}()

	go func() {
		time.Sleep(time.Duration(config_func.Time_gif) * time.Second)
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				config_func.Log_append("error end gif")
			} else {
				config_func.Log_append("end 5s gif")
			}
		}
	}()

	return 1
}
