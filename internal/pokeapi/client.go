package pokeapi

import (
	"time"
	"net/http"

	"github.com/joe-maitan/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache 	   pokecache.Cache
	// pokedex // map[string]Pokemon data structure
} // End Client struct {}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
		
	}
} // End NewClient() function