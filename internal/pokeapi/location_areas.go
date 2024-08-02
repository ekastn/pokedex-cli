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

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := BaseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if val, ok := c.Cache.Get(fullUrl); ok {
		locationAreas := LocationAreasResp{}
		err := json.Unmarshal(val, &locationAreas)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreas, nil
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

	c.Cache.Add(fullUrl, data)
	return locationAreas, nil
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	fullUrl := BaseUrl + "/location-area/" + name

	if val, ok := c.Cache.Get(fullUrl); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("failed to fetch: %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.Cache.Add(fullUrl, data)
	return locationArea, nil
}
