package user

import "github.com/fishjump/cs7ns4/modules/entities"

func Get(key string) (string, error) {
	data := &entities.User{
		Latitude:  77,
		Longitude: 19,
		Timestamp: 1000000,
	}

	return entities.ToJson(data), nil
}

func Put(data *entities.User) error {
	return nil
}
