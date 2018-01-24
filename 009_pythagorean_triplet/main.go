package main

import "fmt"

func main() {
	triplet(1)
}

func triplet(a int) {
	b := (a*a - 1) / 2
	c := (a*a + 1) / 2

	if a+b+c >= 1000 {
		fmt.Printf("Result: a: %v b: %v c: %v sum: %v", a, b, c, a+b+c)
	} else {
		triplet(a + 1)
	}
}
