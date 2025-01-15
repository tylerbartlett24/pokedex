package main

import (
	"time"

	"github.com/tylerbartlett24/pokedex/internal/pokeapi"
	"github.com/tylerbartlett24/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokecache.NewCache(time.Second * 15),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
