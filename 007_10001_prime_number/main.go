package main

import "fmt"

func main() {
	n := 10001
	c := 0

	isPrime := func(x int) bool {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				return false
			}
			return true
		}
	}
	fmt.Prinln(isPrime(3))
}
