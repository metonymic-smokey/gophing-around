package raindrops

import (
	"fmt"
	"strconv"
)

func Convert(n int) string {
	var rain string

	if n%3 == 0 {
		rain += "Pling"
	}

	if n%5 == 0 {
		rain += "Plang"
	}

	if n%7 == 0 {
		rain += "Plong"
	}

	if rain == "" {
		rain = strconv.Itoa(n)
	}

	return rain
}

func raindrops() {
	var n int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&n)

	fmt.Println(Convert(n))
}
