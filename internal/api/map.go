package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Abdulrahman-02/Pokedex/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      *cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check if the data is already in the cache
	if data, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return Locations{}, err
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	// Add the data to the cache
	c.cache.Add(url, data)

	return locations, nil
}
