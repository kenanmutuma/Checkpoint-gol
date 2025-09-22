package main

import (
	"fmt"
	"os"
	"strconv"

	"bank/account"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <credit|debit> <amount>")
		return
	}
	amt, _ := strconv.Atoi(os.Args[2])

	before := bank.Balance

	switch os.Args[1] {
	case "credit":
		bank.Credit(amt)
		fmt.Println("Before:", before, "After:", bank.Balance)
	case "debit":
		if !bank.Debit(amt) {
			fmt.Println("insufficient funds")
		} else {
			fmt.Println("Before:", before, "After:", bank.Balance)
		}
	default:
		fmt.Println("unknown action")
	}
}
