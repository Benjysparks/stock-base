package main

import (
	"fmt"
	"bufio"
	"os"
)


	

func (cfg *apiConfig) LogInMenu() {

	var err error

		usercommands = map[string]usercliCommand{
		"1": {
			name:        "Create User",
			description: "Creates New User",
			callback:    cfg.commandCreateUser,
		},
		"2": {
			name:        "Log in",
			description: "Log in",
			callback:    cfg.commandLogIn,
		},
	}

		scanner := bufio.NewScanner(os.Stdin)
		loggedin := false
		for loggedin == false {
			fmt.Println(" ")
			fmt.Println("1. Add User")
			fmt.Println("2. Log In")
			fmt.Print("\ncommand > ")
			scanner.Scan()
			cleanedInput := CleanInput(scanner.Text())

			if len(cleanedInput) == 0 {
				continue
			}

			commandName := cleanedInput[0]
			loggedBool := false

			command, exists := usercommands[commandName]
			if exists {
				
				if len(cleanedInput) == 1 {

					args := ""
					loggedBool, err = command.callback(args, loggedin)
					if err != nil {
						fmt.Println(err)
				}
				} else {
				args := cleanedInput[1]
				loggedBool, err = command.callback(args, loggedin)
				if err != nil {
					fmt.Println(err)
				}
				}
			} else {
				fmt.Println("Unknown command")
			}
			loggedin = loggedBool
		}
}