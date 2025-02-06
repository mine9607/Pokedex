package pokeapi

import (
	"net/http"
	"time"

	"github.com/mine9607/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetCache() *pokecache.Cache {
	return c.cache
}
