package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListEncounters(name string) (EncounterJson, error) {
	url := baseURL + "location-area/" + name
	val, exist := c.cache.Get(url)
	if exist {
		encounters := EncounterJson{}
		err := json.Unmarshal(val, &encounters)
		if err != nil {
			fmt.Println(err)
			return EncounterJson{}, err
		}
		return encounters, nil

	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return EncounterJson{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return EncounterJson{}, err
	}
	if res.StatusCode > 299 {
		return EncounterJson{}, fmt.Errorf("bad status: %d\n", res.StatusCode)
	}

	encounters := EncounterJson{}
	err = json.Unmarshal(body, &encounters)
	if err != nil {
		fmt.Println(err)
		return EncounterJson{}, err
	}
	c.cache.Add(url, body)
	return encounters, nil
}

type EncounterJson struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string
			Url  string
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
