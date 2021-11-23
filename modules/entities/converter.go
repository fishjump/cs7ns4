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

func AirQualityRequestBodyToAirQuality(src *AirQualityRequestBody) AirQuality {
	data := AirQuality{
		AirQualityIndex: src.Data.AirQualityIndex,
		Station:         src.Data.Station,
		Longitude:       src.Data.City.Geo[0],
		Latitude:        src.Data.City.Geo[1],
		Timestamp:       src.Data.Time.Timestamp,
	}

	return data
}

func UserRequestBodyToUser(src *UserRequestBody) User {
	data := User{
		Latitude:  src.Latitude,
		Longitude: src.Longitude,
		Timestamp: int(src.Timestamp),
	}

	return data
}
