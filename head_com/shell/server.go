package shell

import (
	"io"
	"net/http"
	"os"
)

func Upload_MP3(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "error", http.StatusBadRequest)
			return
		}
		defer file.Close()

		outFile, err := os.Create("uploaded_output.mp3")
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("error"))
	} else {
		http.Error(w, "blat", http.StatusMethodNotAllowed)
	}
}
