package main

import (
	"github.com/joho/godotenv"
	"fmt"
	"bufio"
	"os"
	"log"
	_ "github.com/lib/pq"
	"database/sql"
	"strings"
	"workspace/github.com/Benjysparks/stock-base/internal/database"
)

var commands map[string]cliCommand
var usercommands map[string]usercliCommand

type apiConfig struct {
	db			   *database.Queries
	user			string
}


type cliCommand struct {
    name        string
    description string
    callback    func(args string) error
}

type usercliCommand struct {
    name        string
    description string
    callback    func(args string, loggedin bool) (bool, error)
}

type StockItem struct {
	Name		string		`json:"name"`
	Amount		int			`json:"amount"`
	QtyType		string		`json:"qtytype"`
	PricePer	float64		`json:"priceper"`
}

func CleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	stringSlice := strings.Fields(lowerString)
	return stringSlice
}


func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print("Cound not open connection to database")
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:				dbQueries,
		user:			" ",
	}

	usercommands = map[string]usercliCommand{
	"1": {
		name:        "Create User",
		description: "Creates New User",
		callback:    apiCfg.commandCreateUser,
	},
	"2": {
		name:        "Log in",
		description: "Log in",
		callback:    apiCfg.commandLogIn,
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


	commands = map[string]cliCommand{
		"1": {
			name:        "Create Stock",
			description: "Add new stock item",
			callback:    apiCfg.commandCreateNewStockItem,
		},
		"2": {
			name:        "Check Stock Levels",
			description: "Shows all stock in database",
			callback:    apiCfg.CommandShowAllStock,
		},
	}

	for ;; {
		fmt.Println("1. Create New Stock")
		fmt.Println("2. Show All Stock Levels")
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