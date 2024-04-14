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