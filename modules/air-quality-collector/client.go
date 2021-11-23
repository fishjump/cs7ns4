package airqualitycollector

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	client          *http.Client
	waqiUrlTemplate string = "https://api.waqi.info/feed/@%d/?token=%s"
	token           string = "cbe95f042c7d4d19a51da051fae8b3f47a7e04f8"
	serverUrl       string = "https://uc.superfish.me/db/air-quality"
)

func makeUrl(station int) string {
	return fmt.Sprintf(waqiUrlTemplate, station, token)
}

func init() {
	client = &http.Client{}
}

func getAirqualityFromWaqi() string {
	req, err := http.NewRequest(http.MethodGet, makeUrl(13377), nil)
	if err != nil {
		logger.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}

	return string(body)
}

func sendAirQualityToServer(body io.Reader) {
	req, err := http.NewRequest(http.MethodPost, serverUrl, body)
	if err != nil {
		logger.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
}

func doAirQualityUpdateRequest() {
	body := getAirqualityFromWaqi()
	payload := strings.NewReader(body)
	sendAirQualityToServer(payload)
}
