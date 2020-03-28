package hamming

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	var c int
	c = 0

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				c += 1
			}
		}

		return c, nil
	}

	return c, errors.New("true")

}

func hamming() {
	var s1, s2 string
	var d int

	fmt.Println("Enter a string: ")
	fmt.Scanln(&s1)

	fmt.Println("Enter second string: ")
	fmt.Scanln(&s2)

	d, _ = Distance(s1, s2)
	fmt.Println(d)

}
