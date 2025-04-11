package head_com

import (
	"fmt"
	"log"
	"os"
)

func Read_file(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error %s: %v", filename, err)
	}
	return string(data), nil
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
