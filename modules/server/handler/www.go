package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

var (
	file string
)

func init() {
	exe, err := os.Executable()
	if err != nil {
		logger.Error(err)
		return
	}

	file = filepath.Dir(exe) + "/resource/index.html"
}

func WwwHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, file)
}
