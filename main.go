package main

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	_ "github.com/lib/pq"
	"database/sql"
	"strings"
	"io/ioutil"
	"workspace/github.com/Benjysparks/stock-base/internal/database"
)

var commands map[string]cliCommand
var usercommands map[string]usercliCommand
var stockcommands map[string]stockcliCommand

type apiConfig struct {
	db			   *database.Queries
	User			string
	CurrentStock	string
	CurrentInvoice	int
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

type stockcliCommand struct {
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
		User:			"",
		CurrentStock:	"",
		CurrentInvoice:	0001,
	}

	files, _ := ioutil.ReadDir("./Invoices")
    apiCfg.CurrentInvoice = (len(files)) + 1
	


	apiCfg.LogInMenu()


	
}