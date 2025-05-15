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

type apiConfig struct {
	db			   *database.Queries
}


type cliCommand struct {
    name        string
    description string
    callback    func(args string) error
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
	}


	scanner := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommand{
		"1": {
			name:        "Create Stock",
			description: "Add new stock item",
			callback:    apiCfg.commandCreateNewStockItem,
		},
		"2": {
			name:        "Create Stock",
			description: "Add new stock item",
			callback:    apiCfg.CommandAddStock,
		},
	}

	for ;; {
		fmt.Println("1. Create New Stock")
		fmt.Print("command > ")
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