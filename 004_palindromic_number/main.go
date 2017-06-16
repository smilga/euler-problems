package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindromic(x int) bool {
	str := strconv.Itoa(x)
	l := len(str)
	rev := make([]string, l)

	for i := 0; i < l; i++ {
		rev[i] = string(str[l-i-1])
	}
	rev_str := strings.Join(rev, "")

	return str == rev_str
}

func main() {
	o := 999
	i := 999
	var res int

	for o > 100 {
		for i > 100 {
			if isPalindromic(o*i) && o*i > res {
				res = o * i
			} else {
				i--
			}
		}
		o--
		i = 999
	}
	fmt.Println(res)
}
