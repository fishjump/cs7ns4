package entities

import (
	"os"

	"github.com/withmandala/go-log"
)

type (
	HttpResponseBody struct {
		Successed bool        `json:"successed"`
		Reason    string      `json:"reason"`
		Data      interface{} `json:"data,omitempty"`
	}

	AirQualityRequestBody struct {
		Status string `json:"status"`
		Data   struct {
			AirQualityIndex int `json:"aqi"`
			Station         int `json:"idx"`
			City            struct {
				Geo []float64 `json:"geo"`
			} `json:"city"`
			Time struct {
				Timestamp int `json:"v"`
			} `json:"time"`
		} `json:"data,omitempty"`
	}

	UserRequestBody struct {
		Latitude  float64 `json:"locationLatitude,string"`
		Longitude float64 `json:"locationLongitude,string"`
		Timestamp float64 `json:"locationTimestamp_since1970,string"`
	}

	AirQuality struct {
		AirQualityIndex int     `json:"air-quality-index"`
		Station         int     `json:"station"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		Timestamp       int     `json:"timestamp"`
	}

	AirQualityList []AirQuality

	User struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Timestamp int     `json:"timestamp"`
	}

	UserList []User

	QueryReponse struct {
		AirQualityIndex int     `json:"air-quality-index"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		Timestamp       int     `json:"timestamp"`
	}

	QueryReponseList []QueryReponse
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr)
}

func SuccessedResponseBody(data interface{}) string {
	body := HttpResponseBody{
		Successed: true,
		Reason:    "",
		Data:      data,
	}

	return ToJson(body)
}

func FailedResponseBody(reason string) string {
	body := HttpResponseBody{
		Successed: false,
		Reason:    reason,
	}

	return ToJson(body)
}
