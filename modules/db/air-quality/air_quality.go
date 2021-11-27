package airquality

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/fishjump/cs7ns4/modules/entities"

	"github.com/withmandala/go-log"
)

var (
	dir     string
	storage string

	logger *log.Logger
)

func init() {
	logger = log.New(os.Stderr)

	exe, err := os.Executable()
	if err != nil {
		logger.Error(err)
		return
	}

	dir = filepath.Dir(exe) + "/db/air-quality/"
	storage = dir + "%d/%d.json"

}

func getNearestData(timestamp int, path string) (string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("empty data")
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() {
			return true
		}

		if files[j].IsDir() {
			return false
		}

		return files[i].Name() > files[j].Name()
	})

	if timestamp == -1 {
		data, err := ioutil.ReadFile(path + "/" + files[0].Name())
		if err != nil {
			return "", err
		}

		return string(data), nil
	}

	nearest := -1
	for i := 0; i < len(files); i++ {
		fTimestamp := strings.TrimSuffix(files[i].Name(), ".json")
		fTs, err := strconv.Atoi(fTimestamp)
		if err != nil {
			return "", err
		}

		if fTs < timestamp {
			nearest = i
			break
		}
	}

	if nearest == -1 {
		return "", errors.New("not find nearest data")
	}

	data, err := ioutil.ReadFile(path + "/" + files[nearest].Name())
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getAllData(path string) (string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("empty data")
	}

	data := "["
	for idx, file := range files {
		if file.IsDir() {
			continue
		}

		tmp, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			return "", err
		}

		data += string(tmp)
		if idx != len(files)-1 {
			data += ","
		}
	}
	data += "]"

	return data, nil
}

func GetStation(name string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("empty data")
	}

	// iterate station directories
	var data string
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if file.Name() != name {
			continue
		}

		data, err = getAllData(dir + file.Name())
		if err != nil {
			return "", err
		}

		break
	}

	if data == "" {
		return "", errors.New("empty data")
	}

	return data, nil
}

// Get latest air quality data from each station
func Get(timestamp int) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", errors.New("empty data")
	}

	// iterate station directories
	data := "["
	for idx, file := range files {
		if !file.IsDir() {
			continue
		}

		tmp, err := getNearestData(timestamp, dir+file.Name())
		if err != nil {
			return "", err
		}

		data += tmp
		if idx != len(files)-1 {
			data += ","
		}
	}
	data += "]"

	return data, nil
}

func Put(data *entities.AirQuality) error {
	filename := fmt.Sprintf(storage, data.Station, data.Timestamp)

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = file.WriteString(entities.ToJson(*data))
	if err != nil {
		return err
	}

	logger.Infof("Save new air quality data: %s", filename)
	return nil
}
