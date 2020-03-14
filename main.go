package main

import (
	"fmt"

	"github.com/daengdaengLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposite(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
