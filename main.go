package main

import (
	"github.com/joho/godotenv"
)

type StockItem struct {
	Name		string		`json:"name"`
	Amount		int			`json:"amount"`
	QtyType		string		`json:"qtytype"`
	PricePer	float		`json:"priceper"`
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Print("Cound not open connection to database")
	}
	dbQueries := database.New(db)
}