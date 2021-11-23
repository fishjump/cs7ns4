package entities

import (
	"encoding/json"
)

func ToJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
	}

	return string(jsonStr)
}

func FromJson(jsonStr string, data interface{}) {
	err := json.Unmarshal([]byte(jsonStr), data)
	if err != nil {
		logger.Error(err)
	}
}

func AirQualityRequestBodyToAirQuality(src *AirQualityRequestBody) (AirQuality, error) {
	return AirQuality{}, nil
}

func UserRequestBodyToUser(src *UserRequestBody) (User, error) {
	return User{}, nil
}
