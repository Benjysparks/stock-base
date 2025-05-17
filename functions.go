package main

import (
	"fmt"
	"bufio"
	"os"
)


func (cfg *apiConfig) MainMenu() {



	commands = map[string]cliCommand{
		"1": {
			name:        "Create Stock",
			description: "Add new stock item",
			callback:    cfg.commandCreateNewStockItem,
		},
		"2": {
			name:        "Check Stock Levels",
			description: "Shows all stock in database",
			callback:    cfg.CommandShowAllStock,
		},
		"3": {
			name:        "Search For Stock Item",
			description: "Searchs for a specific stock item",
			callback:    cfg.CommandSearchForStock,
		},
		"0": {
			name:        "Search For Stock Item",
			description: "Searchs for a specific stock item",
			callback:    cfg.commandReadFile,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Println(" ")
		fmt.Printf("Current User: %v ", cfg.User)
		fmt.Println(" ")
		fmt.Println("1. Create New Stock")
		fmt.Println("2. Show All Stock Levels")
		fmt.Println("3. Search For Stock")
		fmt.Print("\ncommand > ")
		scanner.Scan()
		cleanedInput := CleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]


		command, exists := commands[commandName]
		if exists {
			
			if len(cleanedInput) == 1 {

				args := ""
				err := command.callback(args)
				if err != nil {
					fmt.Println(err)
			}
			} else {
			args := cleanedInput[1]
			err := command.callback(args)
			if err != nil {
				fmt.Println(err)
			}
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}	

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

func (cfg *apiConfig) StockEditMenu(stock string) {

	var err error

		stockcommands = map[string]stockcliCommand{
		"1": {
			name:        "Edit Stock Amount",
			description: "edits stock amount",
			callback:    cfg.commandEditStockLevel,
		},
		"2": {
			name:        "Delete Stock Line",
			description: "change stock lines name",
			callback:    cfg.commandShowStockHistory,
		},
		"3": {
			name:        "Show Stock History",
			description: "shows all ajustments to current stock",
			callback:    cfg.commandShowStockHistory,
		},
		"0": {
			name:        "Return",
			description: "returns to previous menu",
			callback:    cfg.commandReturnToMainMenu,
		},
	}

		scanner := bufio.NewScanner(os.Stdin)
		loggedin := false
		for loggedin == false {
			fmt.Println(" ")
			fmt.Printf("Current User: %v ", cfg.User)
			fmt.Println(" ")
			fmt.Println("1. Edit Stock Amount")
			fmt.Println("2. Delete Stock Line")
			fmt.Println("3. Show history")
			fmt.Println("0. Return")
			fmt.Print("\ncommand > ")
			scanner.Scan()
			cleanedInput := CleanInput(scanner.Text())

			if len(cleanedInput) == 0 {
				continue
			}

			commandName := cleanedInput[0]

			command, exists := stockcommands[commandName]
			if exists {
				
				if len(cleanedInput) == 1 {

					args := ""
					err = command.callback(args)
					if err != nil {
						fmt.Println(err)
				}
				} else {
				args := cleanedInput[1]
				err = command.callback(args)
				if err != nil {
					fmt.Println(err)
				}
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
}