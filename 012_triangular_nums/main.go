package main

import "fmt"

var seq = []int{}

func main() {
	i := 0
	for i < 11 {
		nextNum()
		i++
	}
	fmt.Println(seq)

}

func nextNum() {
	next := 0

	for i := 1; i <= len(seq); i++ {
		next += i
	}

	seq = append(seq, next)
}
