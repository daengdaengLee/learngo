package main

import (
	"fmt"

	"github.com/daengdaengLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
