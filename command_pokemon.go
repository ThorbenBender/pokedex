package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackPokedex(cfg *config, args ...string) error {
  if len(cfg.caughtPokemon) == 0 {
    return errors.New("You dont have any pokemon")
  }
  fmt.Println("Your caught pokemon")
  for _, pokemon := range cfg.caughtPokemon {
    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Printf("Abilities: \n")
    for _,ab := range pokemon.Abilities {
      fmt.Printf("- %s\n", ab.Ability.Name)
    }
    fmt.Printf("Types:\n")
    for _, typ := range pokemon.Types {
      fmt.Printf("- %s\n", typ.Type.Name)
    }
  }
  return nil
}

func callbackInspect(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("Missing pokemon name")
  }

  pokemonName := args[0]

  pokemon, ok := cfg.caughtPokemon[pokemonName]

  if !ok {
    return errors.New("You dont have that pokemon")
  }

  fmt.Println("Info about the pokemon")
  fmt.Printf("Name: %s\n", pokemon.Name)
  for _, typ := range pokemon.Types {
    fmt.Printf("Element: %s\n", typ.Type.Name)
  }
  return nil
}

func callbackCatch(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("Missing pokemon name")
  }

  pokemonName := args[0]

  pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName) 

  if err != nil {
    return err
  }


  const threshold = 50
  randNum := rand.Intn(pokemon.BaseExperience)
  if randNum > threshold {
    return fmt.Errorf("Failed to catch %s\n", pokemonName)
  }

  fmt.Printf("%s was caught\n", pokemonName)
  cfg.caughtPokemon[pokemonName] = pokemon

  return nil
}

func callbackPokemonInfo(cfg *config, args ...string) error  {
  if len(args) != 1 {
    return errors.New("Missing pokemon name") 
  }
  pokemonName := args[0]

  pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
  if err != nil {
    return err
  }

  fmt.Printf("Pokemon Info: \n")
  fmt.Printf("Name: %s\n", pokemon.Name)
  fmt.Printf("Abilities: \n")
  for _, ab := range pokemon.Abilities {
    fmt.Printf("- %s - %d\n", ab.Ability.Name, ab.Slot)
  }
  return nil
}
