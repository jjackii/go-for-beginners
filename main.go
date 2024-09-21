package main

import (
	"fmt"

	"github.com/wozlsla/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}