package main

import (
	"fmt"

	"github.com/wozlsla/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
