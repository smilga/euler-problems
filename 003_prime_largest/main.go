package main

import "fmt"

func main() {
	n := 600851475143
	i := 2

	for i*i <= n {
		if n%i == 0 {
			n = n / i
		} else {
			i++
		}
	}
	fmt.Println(n)

}
