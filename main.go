package main

import (
	"fmt"
)

func canIDring(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDring(16))
}
