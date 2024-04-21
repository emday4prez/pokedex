package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	pokecache "pokedex/internal"
	"strings"
	"time"
)
	
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next string
	previous string
}


var started bool = false
var myCache *pokecache.Cache 


func GetHelp()map[string]cliCommand{
return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
				"map": {
						name: "map",
						description: " Displays the names of 20 location areas",
						callback: commandMap,
				},
								"mapb": {
						name: "mapb",
						description: "Displays the previous 20 location areas",
						callback: commandMapb,
				},
}
}

var cliName string = "Pokedex >"

func printPrompt(){
	fmt.Print(cliName)
}



func cleanInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}

func commandHelp() error {
    // Implement the logic to display the help message here
    for cmd, desc := range GetHelp() {
        fmt.Println(cmd, ": ", desc.description)
    }
    return nil // Returning nil indicates no error
}

func commandExit() error {
    os.Exit(0)
    return nil 
}

var location config

// MAP
// The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.



func commandMap() error {
    // Implement the logic to display the next locations here
			var URL string 

			if cachedData, ok := myCache.Get(URL); ok{
					var result LocationGroupResponse
					if err := json.Unmarshal(cachedData, &result); err != nil{
						return fmt.Errorf("error unmarshalling cached data: %v", err)
					}
			}else{
								if started {
					 URL = location.next
				}else{
					URL = "https://pokeapi.co/api/v2/location-area"
				}	
 started = true	
	resp, err := http.Get(URL)
	if err != nil {
	// handle error
	fmt.Println("...there was an error")
			}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
    var result LocationGroupResponse
    if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
        fmt.Println("Can not unmarshal JSON")
    }
				location.next = result.Next

				if(result.Previous != nil){
					location.previous = result.Previous.(string)
				}else{
					location.previous = ""
				}
				for _, loc := range result.Results{
					fmt.Println(loc.Name)
				}
				myCache.Add(URL, body)
}
// Returning nil indicates no error
 return nil
}

// MAPB (MAP BACK)
// Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.

// If you're on the first "page" of results, this command should just print an error message.

func commandMapb() error {
	if !started || location.previous ==  "" {
		return errors.New("nowhere to go")
	}

 resp, err := http.Get(location.previous)
	if err != nil {
	// handle error
	fmt.Sprintln("...there was an error :::: \n", err)
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)

    var result LocationGroupResponse
    if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
        fmt.Println("Can not unmarshal JSON")
    }
				
				for _, loc := range result.Results{
					fmt.Println(loc.Name)
				}
    return nil // Returning nil indicates no error
}







func main() {
				
    reader := bufio.NewScanner(os.Stdin)
				myCache = pokecache.NewCache(5 * time.Second) 
				helpMenu := GetHelp()
				printPrompt()
			
				for reader.Scan(){
						text := cleanInput(reader.Text())
			       if command, exists := helpMenu[text]; exists {
            err := command.callback()
            if err != nil {
                fmt.Println("Error executing command:", err)
            }
        } else {
            fmt.Println("Unknown command")
        }

        printPrompt() 
				}		
}




























type LocationGroup struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}




type LocationGroupResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}