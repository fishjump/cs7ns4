package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/fishjump/cs7ns4/modules/entities"

	"github.com/withmandala/go-log"
)

var (
	logger *log.Logger

	dir     string
	storage string
)

func init() {
	logger = log.New(os.Stderr)

	exe, err := os.Executable()
	if err != nil {
		logger.Error(err)
		return
	}

	dir = filepath.Dir(exe) + "/db/user/"
	storage = dir + "%d.json"

}

// Get latest user upload
func Get() (string, error) {
	files, err := ioutil.ReadDir(dir)
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

	data, err := ioutil.ReadFile(dir + files[0].Name())
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Put(data *entities.User) error {
	filename := fmt.Sprintf(storage, data.Timestamp)

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

	logger.Infof("Save new user data: %s", filename)
	return nil
}
