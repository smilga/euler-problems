package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	r := 20
	s := 2520
	m := false

	match := func(x int) bool {
		for i := 1; i <= r; i++ {
			if x%i != 0 {
				return false
			}
		}
		return true
	}

	for !m {
		s += 20
		if match(s) {
			m = true
		}
	}
	fmt.Println(s)
	fmt.Println(time.Now().Unix() - now)

}
