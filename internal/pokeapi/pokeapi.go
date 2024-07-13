package pokeapi

import (
	"net/http"
	"time"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	HttpClient http.Client
}

func NewClient() Client {
	return Client{
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
