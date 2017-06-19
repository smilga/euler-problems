package main

import "fmt"

func main() {
	start := 1
	end := 100

	var sum int
	var square int

	for i := start; i <= end; i++ {
		sum += i * i
		square += i
	}
	square *= square

	fmt.Println(square - sum)

}
