package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	s := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			s = false
		}
	}
	return s
}
func main() {
	now := time.Now().Unix()

	max := 2000000
	var sum int

	for max > 2 {
		max--
		if isPrime(max) {
			sum += max
		}
	}
	fmt.Println("Result: ", sum)
	fmt.Printf("Calculated %v seconds", time.Now().Unix()-now)
}
