package head_com

import (
	"fmt"
	"os"
	"os/exec"
)

func Read_file(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error %s: %v", filename, err)
	}
	return string(data), nil
}

func Restart() {
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	os.Exit(0)
}
