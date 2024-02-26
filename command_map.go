package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
  resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
  
  if err != nil {
    return err
  }

  for _, area := range resp.Results {
    fmt.Printf(" - %s\n", area.Name)
  }
  cfg.nextLocationAreaURL = resp.Next
  cfg.prevLocationAreaURL = resp.Previous
  return nil
}


func callbackMapB(cfg *config, args ...string) error {
  if cfg.prevLocationAreaURL == nil {
    return errors.New("you are on the first page")
  }
  resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
  
  if err != nil {
    return err
  }

  for _, area := range resp.Results {
    fmt.Printf(" - %s\n", area.Name)
  }
  cfg.nextLocationAreaURL = resp.Next
  cfg.prevLocationAreaURL = resp.Previous
  return nil
}


func callbackExplore(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("No location area provided")
  }

  locationAreaName := args[0]

  locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
  
  if err != nil {
    return err
  }
  fmt.Printf("Pokemon in %s:\n", locationArea.Name)
  for _, pokemon := range locationArea.PokemonEncounters {
    fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
  }
  return nil
}
