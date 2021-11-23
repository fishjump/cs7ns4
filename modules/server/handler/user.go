package handler

import (
	"io"
	"net/http"

	"github.com/fishjump/cs7ns4/modules/db/user"
	"github.com/fishjump/cs7ns4/modules/entities"
)

func handleUserGet(w http.ResponseWriter, req *http.Request) {
	json, err := user.Get()
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}
	entities.FromJson(json, &data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(data)))
}

func handleUserPost(w http.ResponseWriter, req *http.Request) {
	jsonStr, err := io.ReadAll(req.Body)
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	body := &entities.UserRequestBody{}
	entities.FromJson(string(jsonStr), body)
	data := entities.UserRequestBodyToUser(body)

	err = user.Put(&data)
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(nil)))
}

func UserHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	logger.Infof("Method: %s, Path: %s", req.Method, req.RequestURI)

	switch req.Method {
	case http.MethodGet:
		handleUserGet(w, req)
	case http.MethodPost:
		handleUserPost(w, req)
	default:
		makeError(
			&w, http.StatusBadRequest,
			"GET/POST method is required",
		)
	}
}
