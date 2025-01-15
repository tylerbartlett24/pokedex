package main

import (
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name: " + pokemon.Name)
	fmt.Println("Name: ", pokemon.Height)
	fmt.Println("Name: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  -%s\n", pType.Type.Name)
	}
	return nil

}
