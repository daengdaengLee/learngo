package main

import "fmt"

func main() {
	names1 := [5]string{"noci", "lynn", "dal"}
	names1[3] = "alala"
	names1[4] = "alala"
	fmt.Println(names1)

	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "flynn")
	fmt.Println(names2)
}
