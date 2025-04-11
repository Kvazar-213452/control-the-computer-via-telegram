package updata

import (
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

// ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠄⣼⡿⠟⠛⠉⠉⠉⠁⠄⠄⠄⠁⠈⠉⠙⠛⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
// ⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠛⠛⠉⠄⠈⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠉⠁⠄⠛⠻⠿⢿⣿⣿⣿⣿⣿
// ⣿⣿⣿⣿⠿⠛⠉⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠛⠿⣿
// ⣿⣿⠋⠄⠄⠄⣀⣤⡤⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈
// ⣿⣿⣧⠄⠄⢊⣵⣿⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⣶⣶⣶⣶⣶⣶⣶⣤⣤⡀
// ⣿⣿⣿⣷⣄⠞⢡⡟⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠹⣿⣿⣿⣿⣿⣿⣿⣟⠻
// ⣿⣿⣿⣿⣿⣿⡌⠄⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢹⣿⣿⣿⣿⣿⣿⣿⣷
// ⣿⣿⣿⣿⣿⣿⣿⠆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡟⣿⡿⡏⠻⣇⠈⠙
// ⣿⣿⣿⣿⣿⣿⡟⠄⠄⠄⠄⠄⠄⠄⠐⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢹⢻⠃⠄⠄⠈⠄⠄
// ⣿⣿⣿⣿⣿⣿⠁⠄⠄⠄⠄⠄⠄⠄⣾⣧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣸⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠄⠄⠄⠄⠄⢀⣠
// ⣿⣿⣿⣿⣿⡏⢀⠄⠄⠄⠄⠄⠄⣸⣿⣿⣆⠄⠄⠄⠄⠄⠄⠄⢀⣼⣿⣿⣿⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣘⣻⣿⣿
// ⣿⣿⣿⣿⣿⠁⡜⠄⠄⠄⠄⠄⠠⠿⠿⠭⢽⣧⣀⠄⠄⠄⢀⡤⠄⠄⠄⠙⠛⠿⠷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠹⣿⣿⣿
// ⣿⣿⣿⣿⣿⢀⡇⠄⠄⠄⠄⠄⣀⠤⡄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣷⠶⠦⢀⣀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢹⣿⣿
// ⣿⣿⣿⣿⡇⢸⡇⠄⠄⠄⠄⢀⣻⣧⣴⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣖⣿⡿⣷⣳⣷⣦⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿
// ⣿⣿⣿⣿⣷⣼⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿
// ⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿
// ⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⢿⣿⣿⣿⣿⣿⢟⠿⠿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿
// ⣿⣿⣿⣿⣿⣿⡄⠄⠄⠄⠄⠄⠈⢿⣿⣿⣿⡿⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣾⣿
// ⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠙⢿⡿⠄⠄⡀⠄⠄⠄⣠⣾⣿⣿⣿⣿⠟⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣼⣿⣿
// ⣿⣿⣿⣿⣿⣿⡇⠄⡄⠄⠄⣀⡴⣖⣶⣴⣶⡆⠄⠄⠄⠄⣾⠿⠿⠛⠋⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣴⣿⣿⣿⣿
// ⣿⣿⣿⣿⣿⣿⣿⠄⢧⣤⣾⣫⣾⣿⣿⠟⠛⣾⡦⠄⢰⣶⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠻⣿⣿⣿⣿⣿
// ⣿⣿⣿⣿⣿⣿⣫⣾⡿⣫⣾⣿⣿⣿⢽⣿⡶⣿⡧⢀⣷⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⢠⠄⠄⠄⠄⡀⠄⣰⣶⣶⣦⣄⠄⠄⠈⠻⣿⣿⣿
// ⣿⣿⣿⣿⣿⢿⣿⣿⣾⣿⣿⣿⣿⣾⣿⣷⣇⣿⡇⢸⣾⣿⣻⠃⢀⠄⠄⠄⠄⢀⣀⣸⠄⠄⢰⠄⡇⠄⣿⣿⣿⣿⣿⣷⣄⠄⠄⠈⢿⣿
// ⣿⣿⣿⣿⣏⣿⣿⣿⣿⣿⣿⣿⣿⢹⢯⣱⣟⣯⠁⣷⣿⣿⣿⠠⣿⣿⣿⣿⣿⡿⢫⣷⡆⠄⣿⢀⠇⠄⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠹
// ⣿⣿⣿⣿⣼⣿⣿⣿⣿⣿⣿⣿⣿⣐⠿⢿⣿⡷⣾⣿⣿⣿⢿⢳⣿⣿⣿⣿⣿⣶⢻⣽⣷⣸⣿⢸⠄⣼⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄

func Up_music(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	parts := strings.Fields(message.Caption)
	if len(parts) <= 1 {
		config_func.Log_append("#up_msuic not heve args")
		return 0
	}
	key := strings.Join(parts[1:], " ")

	jsonPath := "data/music.json"
	saveDir := "data_use/music"

	os.MkdirAll(saveDir, os.ModePerm)

	fileID := message.Audio.FileID
	fileInfo, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		config_func.Log_append("error get music")
		return 0
	}
	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, fileInfo.FilePath)

	filename := fmt.Sprintf("%x-%d.mp3", config_func.Md5hash(key), time.Now().UnixMilli())
	savePath := filepath.Join(saveDir, filename)

	resp, err := http.Get(url)
	if err != nil {
		config_func.Log_append("errro dwn music")
		return 0
	}
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		config_func.Log_append("errro create msuic file")
		return 0
	}
	defer out.Close()
	io.Copy(out, resp.Body)

	data := map[string]string{}
	readJSON(jsonPath, &data)

	if oldPath, ok := data[key]; ok {
		_ = os.Remove(oldPath)
	}

	data[key] = savePath
	err = writeJSON(jsonPath, data)
	if err != nil {
		config_func.Log_append("errro file music json")
		return 0
	}

	fmt.Printf("%+v\n", message.Audio)

	return 1
}
