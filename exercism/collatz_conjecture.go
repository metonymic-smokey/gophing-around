//package main

package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	var c int
	if n <= 0 {
		return 0, errors.New("0 and negatives are errors")
	}

	if n == 1 {
		return 0, nil
	}

	for n != 1 {
		c++
		if n%2 == 0 {
			n = n / 2
			continue
		}

		if n%2 != 0 {
			n = 3*n + 1
			continue
		}
		fmt.Println(n)
	}

	return c, nil

}

func collatzconjecture() {
	fmt.Println(CollatzConjecture(12))
}
