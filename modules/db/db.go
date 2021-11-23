package db

import (
	"fmt"
	"math"
	"os"
	"sort"

	airquality "github.com/fishjump/cs7ns4/modules/db/air-quality"
	"github.com/fishjump/cs7ns4/modules/db/user"
	"github.com/fishjump/cs7ns4/modules/entities"

	"github.com/withmandala/go-log"
)

var (
	logger *log.Logger

	kNearest = 3
	kLatest  = 100
)

func init() {
	logger = log.New(os.Stderr)

}

func distance(u entities.User, a entities.AirQuality) float64 {
	return math.Sqrt(
		math.Pow(u.Latitude-a.Latitude, 2) +
			math.Pow(u.Longitude-a.Longitude, 2))
}

func getQueryReponse(u entities.User, al entities.AirQualityList) entities.QueryReponse {
	sort.Slice(al, func(i, j int) bool {
		return distance(u, al[i]) < distance(u, al[j])
	})

	var avg int = 0
	for i := 0; i < kNearest; i++ {
		avg += al[i].AirQualityIndex
	}
	avg /= kNearest

	result := entities.QueryReponse{
		Latitude:        u.Latitude,
		Longitude:       u.Longitude,
		Timestamp:       u.Timestamp,
		AirQualityIndex: avg,
	}

	return result
}

// Find the 3 nearest stations and get the average AQI
func Query() (string, error) {
	var uList entities.UserList
	uJson, err := user.GetLatest(kLatest)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	entities.FromJson(uJson, &uList)

	var aList entities.AirQualityList
	aJson, err := airquality.Get()
	if err != nil {
		logger.Error(err)
		return "", err
	}
	entities.FromJson(aJson, &aList)

	if l := len(aList); l < kNearest {
		err := fmt.Errorf("too short station list: %d, expect: %d", l, kNearest)
		logger.Error(err)
		return "", err
	}

	var qList entities.QueryReponseList
	for _, u := range uList {
		qList = append(qList, getQueryReponse(u, aList))
	}

	return entities.ToJson(qList), nil
}
