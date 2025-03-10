package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemonJson(url string) (LocationPokemon, error) {
	if val, ok := c.cache.Get(url); ok {
		locationPokemonResp := LocationPokemon{}
		err := json.Unmarshal(val, &locationPokemonResp)
		if err != nil {
			return LocationPokemon{}, err
		}
		return locationPokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemon{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationPokemon{}, err
	}
	locationPokemonResp := LocationPokemon{}
	err = json.Unmarshal(data, &locationPokemonResp)
	if err != nil {
		return LocationPokemon{}, err
	}
	c.cache.Add(url, data)
	return locationPokemonResp, nil
}
