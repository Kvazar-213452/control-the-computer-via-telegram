package func_unix

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type MusicData map[string]string

func Updata_all() int {
	resp, err := http.Get("http://localhost:3000/music")
	if err != nil {
		fmt.Println("error:", err)
		return 1
	}
	defer resp.Body.Close()

	var musicData MusicData
	err = json.NewDecoder(resp.Body).Decode(&musicData)
	if err != nil {
		fmt.Println("error JSON:", err)
		return 1
	}

	err = clear_directory("data_use/music")
	if err != nil {
		fmt.Println("error read:", err)
		return 1
	}

	for _, url := range musicData {
		err = download_file(url)
		if err != nil {
			fmt.Println("error dwn:", err)
			return 1
		}
	}

	for key, url := range musicData {
		newUrl := "data_use/music/" + filepath.Base(url)
		musicData[key] = newUrl
	}

	err = save_data_file("data/music.json", musicData)
	if err != nil {
		fmt.Println("error save:", err)
		return 1
	}

	return 0
}

func clear_directory(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.RemoveAll(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func download_file(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("data_use/music/" + filepath.Base(url))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func save_data_file(filePath string, data MusicData) error {
	err := ioutil.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		return err
	}

	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, updatedData, 0644)
}
