package config_func

import (
	"crypto/md5"
	"fmt"
	"hash"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Md5hash(text string) []byte {
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

func Del_temp(dirPath string) {
	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			Del_temp(filePath)
			os.Remove(filePath)
		} else {
			os.Remove(filePath)
		}
	}
}

func Log_append(text string) {
	f, err := os.OpenFile("data/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error open log.txt:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(text + "\n"); err != nil {
		log.Println("error in log.txt:", err)
	}
}
