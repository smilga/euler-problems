package main

import "fmt"

func main() {
	fList := []int{1, 2}
	next := 2
	var sum int

	for {
		next = fList[len(fList)-1] + fList[len(fList)-2]
		if next >= 4000000 {
			break
		}
		fList = append(fList, next)
	}

	for _, i := range fList {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)

}
