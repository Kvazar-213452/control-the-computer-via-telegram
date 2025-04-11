package updata

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Up_foto(bot *tgbotapi.BotAPI, message *tgbotapi.Message) int {
	// Отримуємо ключ з caption
	parts := strings.Fields(message.Caption)
	if len(parts) <= 1 {
		return 0
	}
	key := strings.Join(parts[1:], " ")

	// Шляхи
	jsonPath := "data/foto.json"
	saveDir := "data_use/foto"

	// Створення папки, якщо не існує
	os.MkdirAll(saveDir, os.ModePerm)

	// Завантажуємо фото
	fileID := message.Photo[len(message.Photo)-1].FileID
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		log.Println("Помилка отримання файлу:", err)
		return 0
	}

	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

	// Унікальне ім’я файла
	filename := fmt.Sprintf("%x-%d.jpg", md5hash(key), time.Now().UnixMilli())
	savePath := filepath.Join(saveDir, filename)

	// Завантаження фото
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Помилка завантаження фото:", err)
		return 0
	}
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		log.Println("Помилка створення файла:", err)
		return 0
	}
	defer out.Close()
	io.Copy(out, resp.Body)

	// Читання/оновлення JSON
	data := map[string]string{}
	_ = readJSON(jsonPath, &data)

	// Якщо ключ вже є — видаляємо старий файл
	if oldPath, ok := data[key]; ok {
		_ = os.Remove(oldPath)
	}

	data[key] = savePath

	// Збереження JSON
	err = writeJSON(jsonPath, data)
	if err != nil {
		log.Println("Помилка запису JSON:", err)
		return 0
	}

	return 1
}

func md5hash(text string) []byte {
	h := md5Sum()
	h.Write([]byte(text))
	return h.Sum(nil)
}

func md5Sum() hash.Hash {
	hash, _ := hashFromString("md5")
	return hash
}

func hashFromString(algo string) (hash.Hash, error) {
	switch strings.ToLower(algo) {
	case "md5":
		return md5.New(), nil
	default:
		return nil, fmt.Errorf("невідомий алгоритм хешування: %s", algo)
	}
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
