package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (LocationJson, error) {
	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	val, exist := c.cache.Get(url)
	if exist {
		location := LocationJson{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			fmt.Println(err)
			return LocationJson{}, err
		}
		return location, nil

	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationJson{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationJson{}, err
	}
	if res.StatusCode > 299 {
		return LocationJson{}, fmt.Errorf("bad status: %d", res.StatusCode)
	}

	location := LocationJson{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
		return LocationJson{}, err
	}
	c.cache.Add(url, body)
	return location, nil
}

type LocationJson struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string
		Url  string
	} `json:"results"`
}
