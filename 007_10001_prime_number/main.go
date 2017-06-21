package main

import "fmt"

func main() {
	n := 10001

	isPrime := func(x int) bool {
		out := true
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				out = false
				break
			}
		}
		return out
	}

	i := 1
	for n > 0 {
		i++
		if isPrime(i) {
			n--
		}
	}
	fmt.Println("counter: ", n, "Prime nr.: ", i)

}
