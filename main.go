package main

import (
	"fmt"
)

func canIDring(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDring(16))
}
