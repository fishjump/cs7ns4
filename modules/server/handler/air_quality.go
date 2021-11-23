package handler

import (
	"io"
	"net/http"

	airquality "github.com/fishjump/cs7ns4/modules/db/air-quality"
	"github.com/fishjump/cs7ns4/modules/entities"
)

func handleAirQualityGet(w http.ResponseWriter, req *http.Request) {
	json, err := airquality.Get()
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	var data interface{}
	entities.FromJson(json, &data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(data)))

}

func handleAirQualityPost(w http.ResponseWriter, req *http.Request) {
	jsonStr, err := io.ReadAll(req.Body)
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	body := &entities.AirQualityRequestBody{}
	entities.FromJson(string(jsonStr), body)
	data := entities.AirQualityRequestBodyToAirQuality(body)

	err = airquality.Put(&data)
	if err != nil {
		makeError(&w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(entities.SuccessedResponseBody(nil)))
}

func AirQualityHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	logger.Infof("Method: %s, Path: %s", req.Method, req.RequestURI)

	switch req.Method {
	case http.MethodGet:
		handleAirQualityGet(w, req)
	case http.MethodPost:
		handleAirQualityPost(w, req)
	default:
		makeError(
			&w, http.StatusBadRequest,
			"GET/POST method is required",
		)
	}

}
