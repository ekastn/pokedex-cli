package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := BaseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationAreasResp{}, fmt.Errorf("failed to fetch location areas: %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}
