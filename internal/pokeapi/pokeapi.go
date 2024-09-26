package pokeapi

import (
	"net/http"
	"time"

	"github.com/ekastn/pokedex-cli/internal/pokecache"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	HttpClient http.Client
	Cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
		Cache: pokecache.NewCache(time.Minute * 5),
	}
}
