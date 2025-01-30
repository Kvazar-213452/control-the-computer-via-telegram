package func_unix

import (
	"encoding/json"
	"fmt"
	"head/head_com/config"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Data_Updata map[string]string

func Updata_all() int {
	resp, err := http.Get(config.Server_data + "music")
	if err != nil {
		fmt.Println("error:", err)
		return 1
	}
	defer resp.Body.Close()

	var musicData Data_Updata
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
		err = download_file(url, "data_use/music/")
		if err != nil {
			fmt.Println("error dwn, mb I dont no:", err)
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

	// dwn foto// dwn foto// dwn foto
	// dwn foto// dwn foto// dwn foto
	// dwn foto// dwn foto// dwn foto

	var folot_data Data_Updata

	resp, err = http.Get(config.Server_data + "foto")
	if err != nil {
		fmt.Println("error:", err)
		return 1
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&folot_data)
	if err != nil {
		fmt.Println("error JSON:", err)
		return 1
	}

	err = clear_directory("data_use/foto")
	if err != nil {
		fmt.Println("error read:", err)
		return 1
	}

	for _, url := range folot_data {
		err = download_file(url, "data_use/foto/")
		if err != nil {
			fmt.Println("error dwn:", err)
			return 1
		}
	}

	for key, url := range folot_data {
		newUrl := "data_use/foto/" + filepath.Base(url)
		folot_data[key] = newUrl
	}

	err = save_data_file("data/foto.json", folot_data)
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

func download_file(url string, phat string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(phat + filepath.Base(url))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func save_data_file(filePath string, data Data_Updata) error {
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
