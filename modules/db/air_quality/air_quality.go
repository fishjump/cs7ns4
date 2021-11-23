package airquality

import (
	"github.com/fishjump/cs7ns4/modules/entities"
)

func Get(key string, data *entities.AirQuality) error {
	*data = entities.AirQuality{
		AirQualityIndex: 77,
		Station:         19,
		Timestamp:       1000000,
	}

	return nil
}

func Put(data *entities.AirQuality) error {
	return nil
}
