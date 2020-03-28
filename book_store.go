package main

import (
	"fmt"
	"sort"
)

func Cost(books []int) int {

	l := len(books)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return 800
	}

	if l == 2 {
		if books[0] == books[1] {
			return 1600
		}

		return int(0.95 * 1600)
	}

	sort.Ints(books)

	x := make(map[int]int)

	for i := 0; i < l; i++ {
		x[books[i]] += 1
	}

	if len(x) == 1 {
		return 800*len(x)
	}

	if len(x) == 2 {
		return int(800 + 0.95*1600)
	}

	

	return 2400

}

func main() {
	n := []int{1, 2,1}

	fmt.Println(Cost(n))
}
