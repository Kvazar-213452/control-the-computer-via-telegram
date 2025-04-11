package updata

import (
	"encoding/json"
	"fmt"
	"head/head_com/config_func"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Up_foto(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	parts := strings.Fields(message.Caption)
	if len(parts) <= 1 {
		return 0
	}
	key := strings.Join(parts[1:], " ")

	jsonPath := "data/foto.json"
	saveDir := "data_use/foto"

	os.MkdirAll(saveDir, os.ModePerm)

	fileID := message.Photo[len(message.Photo)-1].FileID
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		config_func.Log_append("errro get file")
		return 0
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

	filename := fmt.Sprintf("%x-%d.jpg", config_func.Md5hash(key), time.Now().UnixMilli())
	savePath := filepath.Join(saveDir, filename)

	resp, err := http.Get(url)
	if err != nil {
		config_func.Log_append("error dwn foto")
		return 0
	}
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		config_func.Log_append("error foto create")
		return 0
	}
	defer out.Close()
	io.Copy(out, resp.Body)

	data := map[string]string{}
	_ = readJSON(jsonPath, &data)

	if oldPath, ok := data[key]; ok {
		_ = os.Remove(oldPath)
	}

	data[key] = savePath

	err = writeJSON(jsonPath, data)
	if err != nil {
		config_func.Log_append("erro add json foto/music")
		return 0
	}

	return 1
}

func readJSON(path string, out *map[string]string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			*out = map[string]string{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, out)
}

func writeJSON(path string, data map[string]string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, jsonData, 0644)
}
