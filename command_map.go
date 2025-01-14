package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tylerbartlett24/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	if cfg.nextLocationsURL != nil {
		rawData, ok := cfg.cache.Get(*cfg.nextLocationsURL)
		if ok {
			locationsResp := pokeapi.RespShallowLocations{}
			err := json.Unmarshal(rawData, &locationsResp)
			if err != nil {
				return err
			}

			cfg.nextLocationsURL = locationsResp.Next
			cfg.prevLocationsURL = locationsResp.Previous

			for _, loc := range locationsResp.Results {
				fmt.Println(loc.Name)
			}
			return nil

		}
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	newEntry, err := json.Marshal(locationsResp)
	if err != nil {
		return err
	}
	if cfg.nextLocationsURL != nil {
		cfg.cache.Add(*cfg.nextLocationsURL, newEntry)
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	rawData, ok := cfg.cache.Get(*cfg.prevLocationsURL)
	if ok {
		locationsResp := pokeapi.RespShallowLocations{}
		err := json.Unmarshal(rawData, &locationsResp)
		if err != nil {
			return err
		}

		cfg.nextLocationsURL = locationsResp.Next
		cfg.prevLocationsURL = locationsResp.Previous

		for _, loc := range locationsResp.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
