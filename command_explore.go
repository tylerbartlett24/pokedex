package main

import (
	"encoding/json"
	"fmt"

	"github.com/tylerbartlett24/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, name string) error {
	if name == "" {
		fmt.Println("The Explore command should be followed by an area name.")
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + name
	cacheData, ok := cfg.cache.Get(url)
	if ok {
		location := pokeapi.Location{}
		err := json.Unmarshal(cacheData, &location)
		if err != nil {
			return err
		}
		for _, pokemon := range location.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}
		return nil
	}

	location, err := cfg.pokeapiClient.ExploreLocation(url)
	if err != nil {
		return err
	}
	newEntry, err := json.Marshal(location)
	if err != nil {
		return err
	}
	cfg.cache.Add(url, newEntry)
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
