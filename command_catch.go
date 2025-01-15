package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/tylerbartlett24/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		fmt.Println("The Catch command should be followed by a Pokemon.")
		return nil
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + name
	cacheData, ok := cfg.cache.Get(url)

	if ok {
		pokemon := pokeapi.Pokemon{}
		err := json.Unmarshal(cacheData, &pokemon)
		if err != nil {
			return err
		}
		fmt.Println("Throwing a Pokeball at " + name + "...")
		roll := rand.Intn(610)
		if pokemon.BaseExperience > roll {
			fmt.Println(name + " escaped!")
		} else {
			fmt.Println(name + " was caught!")
			cfg.pokedex[name] = pokemon
		}
		return nil
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(url)
	if err != nil {
		return nil
	}

	newEntry, err := json.Marshal(pokemon)
	if err != nil {
		return err
	}
	cfg.cache.Add(url, newEntry)

	fmt.Println("Throwing a Pokeball at " + name + "...")
	roll := rand.Intn(610)
	if pokemon.BaseExperience > roll {
		fmt.Println(name + " escaped!")
	} else {
		fmt.Println(name + " was caught!")
		cfg.pokedex[name] = pokemon
	}
	return nil
}
