package main

import (
	"fmt"
	"context"
)

func (cfg *apiConfig) CommandAddStock(args string) error {
	fmt.Println("Command issued")
	stock, err := cfg.db.ShowAllStock(context.Background())
	if err != nil {
		fmt.Println("error")
	}
	fmt.Print(stock)
	return nil
}