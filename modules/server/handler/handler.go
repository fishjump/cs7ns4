package handler

import (
	"net/http"
	"os"

	"github.com/fishjump/cs7ns4/modules/entities"
	"github.com/withmandala/go-log"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr)
}

func makeError(w *http.ResponseWriter, status int, err string) {
	(*w).WriteHeader(status)
	(*w).Write([]byte(entities.FailedResponseBody(err)))
	logger.Error(err)
}
