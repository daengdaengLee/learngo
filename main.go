package main

import "fmt"

func main() {
	const name1 string = "kunho"
	fmt.Println(name1)

	var name2 string = "nico"
	name2 = "lynn"
	fmt.Println(name2)

	// 함수 안에서만 가능한 축약 표현
	name3 := "nico"
	name3 = "lynn"
	fmt.Println(name3)
}
