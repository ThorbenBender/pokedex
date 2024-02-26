package pokeapi

import (
	"net/http"
	"time"

	"github.com/thorbenbender/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
  cache pokecache.Cache
  httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
  return Client{
    httpClient: http.Client{
      Timeout: time.Minute,

    },
    cache: pokecache.NewCache(cacheInterval),
  }
}


