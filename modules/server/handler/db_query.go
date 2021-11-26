package handler

import (
	"net/http"

	"github.com/fishjump/cs7ns4/modules/db"
	airquality "github.com/fishjump/cs7ns4/modules/db/air-quality"
	"github.com/fishjump/cs7ns4/modules/entities"
)

func handleDbQueryUserGet(w http.ResponseWriter, req *http.Request) {
	json, err := db.Query(req.URL.Query().Get("username"))
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}
	entities.FromJson(json, &data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(data)))
}

func handleDbQueryStationGet(w http.ResponseWriter, req *http.Request) {
	json, err := airquality.GetStation(req.URL.Query().Get("station"))
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}
	entities.FromJson(json, &data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(data)))
}

func handleDbQueryGet(w http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Has("username") {
		handleDbQueryUserGet(w, req)
		return
	}

	if req.URL.Query().Has("station") {
		handleDbQueryStationGet(w, req)
		return
	}

	makeError(&w, http.StatusBadRequest, "expect query parameter: username or station")
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
