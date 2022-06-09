package main

import "fmt"

func main() {
	fmt.Println("Enter table")
	input := 0

	fmt.Scanln(&input)
	tab(input)

}
func tab(input int) {
	var i int
	j := 1

	count := (input * 10) + 1
	for i = input; i < count; i += input { //i=i+input
		fmt.Println(input, "*", j, "=", i)
		j += 1

	}

}
