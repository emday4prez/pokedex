package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
						description: "Displays the names of 20 location areas",
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

// MAP
// The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.

func commandMap() error {
    // Implement the logic to display the locations here
    for cmd, desc := range GetHelp() {
        fmt.Println(cmd, ": ", desc.description)
    }
    return nil // Returning nil indicates no error
}

// MAPB (MAP BACK)
// Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.

// If you're on the first "page" of results, this command should just print an error message.


func commandMapb() error {
    // Implement the logic to display the previous locations here
    for cmd, desc := range GetHelp() {
        fmt.Println(cmd, ": ", desc.description)
    }
    return nil // Returning nil indicates no error
}

func main() {
    reader := bufio.NewScanner(os.Stdin)
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