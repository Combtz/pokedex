package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaJSON(url string) (LocationData, error) {
	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationData{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationData{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}
	locationResp := LocationData{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationData{}, err
	}
	c.cache.Add(url, data)
	return locationResp, nil
}
