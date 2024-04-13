package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
	
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
        callback:    os.Exit(0),
    },
}
}

var cliName string = "Pokedex >"

func printPrompt(){
	fmt.Println(cliName)
}


func cleanInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}



func main() {
    reader := bufio.NewScanner(os.Stdin)
				helpMenu := GetHelp()
				printPrompt()
				for reader.Scan(){
						text := cleanInput(reader.Text())
						if text == "help"{
							for cmd, desc := range helpMenu {
								fmt.Println(cmd, desc)
							}
						}else if text == "exit" {
							 helpMenu["exit"]
						}
				}			printPrompt()
}