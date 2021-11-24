package handler

import (
	"net/http"

	"github.com/fishjump/cs7ns4/modules/db"
	"github.com/fishjump/cs7ns4/modules/entities"
)

func handleDbQueryGet(w http.ResponseWriter, req *http.Request) {
	json, err := db.Query()
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}
	entities.FromJson(json, &data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(data)))

}

func DbQueryHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	logger.Infof("Method: %s, Path: %s", req.Method, req.RequestURI)

	switch req.Method {
	case http.MethodGet:
		handleDbQueryGet(w, req)
	default:
		makeError(
			&w, http.StatusBadRequest,
			"GET method is required",
		)
	}

}
