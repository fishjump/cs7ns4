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
	stations        map[string]int
)

func makeUrl(station int) string {
	return fmt.Sprintf(waqiUrlTemplate, station, token)
}

func init() {
	client = &http.Client{}
	stations = make(map[string]int)

	stations["Davitt Road, Inchicore, Dublin 12, Ireland"] = 13377
	stations["Marino, Dublin 3, Ireland"] = 13390
	stations["St. John’s Road, Kilmainham, Dublin 8, Ireland"] = 13405
	stations["Dublin Port, Dublin 1, Ireland"] = 13379
	stations["Rathmines, Ireland"] = 5112
	stations["Ringsend, Dublin 4, Ireland"] = 13400
	stations["St. Anne's Park, Dublin 5, Ireland"] = 13404
	stations["Clonskeagh, Ireland"] = 13372
	stations["Finglas, Dublin 11, Ireland"] = 13384
	stations["Sandymount Green, Dublin 4, Ireland"] = 13402
	stations["Weaver Park, Dublin 8, Ireland"] = 13412
	stations["Drumcondra Library, Dublin 9, Ireland"] = 13378
	stations["Amiens Street, Dublin 1, Ireland"] = 13358
	stations["Coolock, Dublin 5, Ireland"] = 13374
	stations["Ballymun Library, Dublin 9, Ireland"] = 13363
	stations["Walkinstown Library, Dublin 12, Ireland"] = 13411
	stations["Lord Edward Street, Dublin 2, Ireland"] = 13387
	stations["Custom House Quay, Dublin 1, Ireland"] = 13376
}

func getAirqualityFromWaqi(station int) string {
	req, err := http.NewRequest(http.MethodGet, makeUrl(station), nil)
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
	for _, station := range stations {
		body := getAirqualityFromWaqi(station)
		payload := strings.NewReader(body)
		sendAirQualityToServer(payload)
	}
}
