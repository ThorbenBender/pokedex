package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func startRepl(cfg *config) {
  fmt.Println("eh")
  scanner := bufio.NewScanner(os.Stdin)
  for {

    fmt.Println("Please enter some text >")
    scanner.Scan()
    text := scanner.Text()
    
    cleaned := cleanInput(text)

    if len(cleaned) == 0 {
      continue
    }

    commandName := cleaned[0]
    args := []string{}
    if len(cleaned) > 1 {
      args = cleaned[1:]
    }

    availableCommands := getCommands()

    command, ok := availableCommands[commandName]

    if !ok {
      fmt.Println("Invalid command")
      continue
    }

    err := command.callback(cfg, args...)
    
    if err != nil {
      fmt.Println(err)
    }
  }
}

func getCommands() map[string]clicommand {
  return map[string]clicommand{
    "help": {
      name: "help",
      description: "Prints the help menu",
      callback: callbackHelp, 
    },
    "exit": {
      name: "exit",
      description: "Turns off pokedex",
      callback: callbackExit,
    },
    "map": {
      name: "map",
      description: "Get pokemon map locations",
      callback: callbackMap,
    },
    "mapb": {
      name: "mapb",
      description: "Get previous map locations",
      callback: callbackMapB,
    },
    "explore": {
      name: "explore {location_area}",
      description: "Explore a location in pokemon",
      callback: callbackExplore,
    },
    "info": {
      name: "info {pokemon_name}",
      description: "Get pokemon info",
      callback: callbackPokemonInfo,
    },
    "catch": {
      name: "catch {pokemon_name}",
      description: "Try to catch pokemon by name",
      callback: callbackCatch,
    },
    "inspect": {
      name: "inspect {pokemon_name}",
      description: "Inspect caught pokemon stats",
      callback: callbackInspect,
    },
    "pokedex": {
      name: "pokedex",
      description: "List caught pokemon",
      callback: callbackPokedex,
    },
  }
}

type clicommand struct {
  name string
  description string 
  callback func(*config, ...string) error
}

func cleanInput(str string) []string {
  lowered := strings.ToLower(str)
  words := strings.Fields(lowered)
  return words
}
