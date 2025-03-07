package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreaJSON(url string) (LocationData, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationData{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationData{}, err
	}
	if res.StatusCode > 299 {
		errMes := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationData{}, errors.New(errMes)
	}
	var locationData LocationData
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return LocationData{}, err
	}
	return locationData, nil
}
