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

func calc() chan int {
	c := make(chan int)
	done := make(chan bool)

	max := 2000000

	for i := 2; i < max; i++ {
		go func(i int) {
			if isPrime(i) {
				c <- i
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 2; i < max; i++ {
			<-done
		}
		close(c)
	}()

	return c

}

func main() {
	now := time.Now().Unix()

	c := calc()
	var sum int

	for n := range c {
		sum += n
	}

	fmt.Println("Result: ", sum)
	fmt.Printf("Calculated %v seconds", time.Now().Unix()-now)
}
