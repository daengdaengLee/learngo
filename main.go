package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, number := range numbers {
		fmt.Println(index, number)
	}

	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
