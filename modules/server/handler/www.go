package handler

import (
	"net/http"

	"github.com/fishjump/cs7ns4/modules/www"
)

func WwwHandler(w http.ResponseWriter, req *http.Request) {
	resp, err := www.Serve()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
