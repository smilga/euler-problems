package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const ll = 6

var sum = []int{1, 2, 3, 4, 5, 6, 7} //make([]int, ll)

func main() {
	n := `321042
	      321655`

	ns := bytes.Split([]byte(n), []byte("\n"))

	for _, s := range ns {
		addNumber(bytes.TrimSpace(s))
	}
	fmt.Println(sum)

}

func addNumber(n []byte) {
	sumLen := len(sum)
	j := 0
	mem := 0
	for i := sumLen - 1; i >= 0; i-- {
		this := len(n) - j
		next := len(n) - j - 1
		add := 0
		if next >= 0 {
			add, _ = strconv.Atoi(string(n)[next:this])
		}
		fmt.Println(sum[i], add, mem)
		if sum[i]+add+mem == 10 {
			sum[i] = 0
			mem = 1

		} else if sum[i]+add+mem > 9 {
			sum[i] += (sum[i] + add + mem) - 10
			mem = 1
		} else {
			sum[i] += add + mem
			mem = 0
		}
		j++
	}
	if mem == 1 {
		sum = append([]int{mem}, sum...)
	}

}
