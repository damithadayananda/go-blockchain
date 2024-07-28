package http

import (
	"go-blockchain/ui"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
)

func InitServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := r.URL.Path
		if requestedPath == "/" {
			requestedPath = "/index.html"
		}
		file, err := ui.UI.Open("frontend/go-blockchain/build" + requestedPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fileInfo, _ := file.Stat()
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, path.Base(requestedPath), fileInfo.ModTime(), strings.NewReader(string(content)))
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf(err.Error())
	}
}
