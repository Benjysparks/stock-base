package main

import (
	"fmt"
	"context"
	"bufio"
	"strconv"
	"os"
	"workspace/github.com/Benjysparks/stock-base/internal/database"
)

func (cfg *apiConfig) commandCreateNewStockItem(args string) error {

	scanner := bufio.NewScanner(os.Stdin)
	stockName := ""
	var stockAmount int32
	stockQtyType := ""
	var stockPricePer int32


	for ;; {
		fmt.Print("New stock name > ")
		scanner.Scan()
		stockName = scanner.Text()
		break
		}
	for ;; {
		fmt.Print("New stock amount > ")
		scanner.Scan()
		tempAmount, _ := strconv.Atoi(scanner.Text())
		stockAmount = int32(tempAmount)
		break
		}
	for ;; {
		fmt.Print("New stock quantity type > ")
		scanner.Scan()
		stockQtyType = scanner.Text()
		break
		}
	for ;; {
		fmt.Print("New price per stock > ")
		scanner.Scan()
		tempPricePer, _ := strconv.Atoi(scanner.Text())
		stockPricePer = int32(tempPricePer)
		break
		}
	cfg.db.CreateStockItem(context.Background(), database.CreateStockItemParams{
		Stockname:		stockName,
		Amount:			stockAmount,
		QtyType:		stockQtyType,
		PricePer:		stockPricePer,
	})
	return nil
}


func addSpaces (args ...string) string {
		wordLine := ""
		for _, word := range args {
		spacesNeeded := 20 - len(word)
		spaces := ""
		for i := 0; i < spacesNeeded; i++ {
			spaces = spaces + " "
		}
		wordLine = wordLine + (word + spaces)
	}
	return wordLine
}

func (cfg *apiConfig) CommandShowAllStock(args string) error {
	stock, err := cfg.db.ShowAllStock(context.Background())
	if err != nil {
		fmt.Sprintf("")
	}
	fmt.Println(addSpaces("Stockname", "Amount", "Quantity Type", "Price Per Unit"))
	fmt.Println(" ")
	for _, item := range stock {
		fmt.Println(addSpaces(item.Stockname, fmt.Sprint(item.Amount), item.QtyType, fmt.Sprint(item.PricePer)))
	}
	fmt.Println(" ")
	return nil
}

func (cfg *apiConfig) commandCreateUser(args string, loggedin bool) (bool, error) {
	
	scanner := bufio.NewScanner(os.Stdin)
	userName := ""
	passWord := ""

		for ;; {
		fmt.Print("New User Name > ")
		scanner.Scan()
		userName = scanner.Text()
		break
		}
		for ;; {
		fmt.Print("New Password > ")
		scanner.Scan()
		passWord = scanner.Text()
		break
		}
		cfg.db.CreateNewUser(context.Background(), database.CreateNewUserParams{
			UserName:		userName,
			PassWord:		passWord,
		})
		return false, nil
}

func (cfg *apiConfig) commandLogIn(args string, loggedin bool) (bool, error) {
		scanner := bufio.NewScanner(os.Stdin)
	userName := ""
	passWord := ""
	storedPass := ""
	var err error

		for ;; {
		fmt.Print("Enter User Name > ")
		scanner.Scan()
		userName = scanner.Text()
		storedPass, err = cfg.db.GetPassword(context.Background(), userName)
		if err != nil {
			fmt.Println(" ")
			fmt.Println("Invalid Username")
			fmt.Println(" ")
		} else{
			break
		}
		}



		for ;; {
		fmt.Print("New Password > ")
		scanner.Scan()
		passWord = scanner.Text()
		if passWord == storedPass{
			cfg.user = userName
			return true, nil
		} else {
			fmt.Println(" ")
			fmt.Println("Invalid Username")
			fmt.Println(" ")
		}
		break
		}
		return false, nil
}