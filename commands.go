package main

import (
	"fmt"
	"context"
	"bufio"
	"strconv"
	"os"
	"io"
	"log"
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
	stockAdjustment := fmt.Sprintf("New Stock Line Created: %v", stockName)
	cfg.db.LogHistory(context.Background(), database.LogHistoryParams{
		UserName:		cfg.User,
		Stockname:		stockName,
		Adjustment:		stockAdjustment,
	})
	return nil
}

func (cfg *apiConfig) commandShowStockHistory(args string) error {
	stock := cfg.CurrentStock
	fmt.Println(" ")
	fmt.Printf("\nShowing adjustment history for %v", stock)
	fmt.Println(" ")

	logs, err := cfg.db.ShowStockHistory(context.Background(), stock)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(addSpaces("Date edited","User","Adjustment"))
	fmt.Println(" ")
	for _, log := range logs {
		fmt.Println(addSpaces(log.ToChar), addSpaces(log.UserName), addSpaces(log.Adjustment))
		fmt.Println(" ")
		fmt.Println(len(addSpaces(log.UserName)))
	}
	return nil
}

func (cfg *apiConfig) commandEditStockLevel(args string) error {

	var stockAdjustment int32
	adjustmentReason := ""
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("How would you like to adjust stock level? (+/- amount) > ")
		scanner.Scan()
		tempAmount, _ := strconv.Atoi(scanner.Text())
		stockAdjustment = int32(tempAmount)
		break
		}
	for ;; {
		fmt.Print("What is the reason for the adjustment? > ")
		scanner.Scan()
		adjustmentReason = scanner.Text()
		break
		}
	
		
	
	cfg.db.AdjustStockAmount(context.Background(), database.AdjustStockAmountParams{
		Amount:		stockAdjustment,
		Stockname:	cfg.CurrentStock,
	})

	stockSlice, err := cfg.db.VagueStockSearch(context.Background(), cfg.CurrentStock)
	if err != nil {
		fmt.Println(" ")
		return err
	}

	fmt.Println(" ")
	fmt.Println("New stock level:")
	fmt.Println(" ")
	fmt.Println(addSpaces("Stockname", "Amount", "Quantity Type", "Price Per Unit"))
	fmt.Println(" ")
	fmt.Println(addSpaces(stockSlice[0].Stockname, fmt.Sprint(stockSlice[0].Amount), stockSlice[0].QtyType, fmt.Sprint(stockSlice[0].PricePer)))
	
	logAdjustment := fmt.Sprintf("New Stock Adjustment: %v \n%vStock adjusted by: %v \n%vReason: %v", stockSlice[0].Stockname, addCustomSpaces(60), stockAdjustment, addCustomSpaces(60),adjustmentReason)
	cfg.db.LogHistory(context.Background(), database.LogHistoryParams{
		UserName:	cfg.User,
		Stockname:	cfg.CurrentStock,
		Adjustment:	logAdjustment,
	})
	return nil
}

func (cfg *apiConfig) commandReturnToMainMenu(args string) error {
	cfg.MainMenu()
	return nil
}

func addSpaces (args ...string) string {
		wordLine := ""
		for _, word := range args {
		spacesNeeded := 29 - len(word)
		spaces := ""
		for i := 0; i < spacesNeeded; i++ {
			spaces = spaces + " "
		}
		wordLine = wordLine + (word + spaces)
	}
	return wordLine
}

func addCustomSpaces (spacesNeeded int) string {
	spaces := ""
		for i := 0; i < spacesNeeded; i++ {
			spaces = spaces + " "
		}
	return spaces
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

func (cfg *apiConfig) CommandSearchForStock(args string) error {

	scanner := bufio.NewScanner(os.Stdin)
	stockName := ""

	for ;; {
		fmt.Print("Enter Stock Name > ")
		scanner.Scan()
		stockName = scanner.Text()
		break
		}
	
	stockName = "%" + stockName + "%"

	stockSlice, err := cfg.db.VagueStockSearch(context.Background(), stockName)
	if err != nil {
		fmt.Println(" ")
		return err
	}
	if len(stockSlice) == 0 {
		fmt.Println(" ")
		fmt.Println("No stock found, please refine search")
		fmt.Println(" ")
	} else if len(stockSlice) == 1 {
		fmt.Println(" ")
		fmt.Println(addSpaces("Stockname", "Amount", "Quantity Type", "Price Per Unit"))
		fmt.Println(" ")
		fmt.Println(addSpaces(stockSlice[0].Stockname, fmt.Sprint(stockSlice[0].Amount), stockSlice[0].QtyType, fmt.Sprint(stockSlice[0].PricePer)))
		fmt.Println(" ")
		cfg.CurrentStock = stockSlice[0].Stockname
		cfg.StockEditMenu(stockSlice[0].Stockname)
		return nil
	}
	fmt.Println(" ")
	fmt.Println("Multiple stock lines found, for stock editing search for a specific stock line")
	fmt.Println(" ")
	fmt.Println(addSpaces("Stockname", "Amount", "Quantity Type", "Price Per Unit"))
	fmt.Println(" ")
	for _, item := range stockSlice {
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
		fmt.Println(" ")
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
		fmt.Print("Enter Password > ")
		scanner.Scan()
		passWord = scanner.Text()
		if passWord == storedPass{
			cfg.User = userName
			cfg.MainMenu()
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

func (cfg *apiConfig) commandReadFile(args string) error {
	f, err := os.Open("Test.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    defer f.Close()
    buf := make([]byte, 1024)
    for {
        n, err := f.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		fmt.Println(err)
		continue
	}
	if n > 0 {
		fmt.Println(string(buf[:n]))
	}
    }
	return nil
}