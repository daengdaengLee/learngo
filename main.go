package main

import (
	"fmt"

	"github.com/daengdaengLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposite(10)
	account.ChangeOwner("daengdaenglee")
	fmt.Println(account)
}
