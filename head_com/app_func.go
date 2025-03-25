package head_com

import (
	"fmt"
	"os"
)

func Read_file(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error %s: %v", filename, err)
	}
	return string(data), nil
}
