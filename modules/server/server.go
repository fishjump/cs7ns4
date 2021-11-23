package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/withmandala/go-log"

	"github.com/fishjump/cs7ns4/modules/server/handler"
)

type (
	HttpHandler func(http.ResponseWriter, *http.Request)
)

var (
	handlers map[string]HttpHandler
	logger   *log.Logger
)

func init() {
	logger = log.New(os.Stderr)
}

func init() {
	handlers = make(map[string]HttpHandler)
	handlers["/"] = handler.WwwHandler
	handlers["/db/air-quality"] = handler.AirQualityHandler
	handlers["/db/user"] = handler.UserHandler
}

func listenAndServe(ip string, port int) {
	for k, v := range handlers {
		http.HandleFunc(k, v)
	}

	addr := fmt.Sprintf("%s:%d", ip, port)
	logger.Infof("Starting serve at %s", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		logger.Error(err)
	}
}

func Run(ip string, port int) {
	listenAndServe(ip, port)
}
