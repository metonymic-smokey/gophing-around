package pythagorean

import (
	"fmt"
)

type Triplet = [3]int

func Range(min, max int) []Triplet {
	var a, b, c int
	p := 0
	var x []Triplet

	for a = min; a <= max; a++ {
		for b = a; b <= max; b++ {
			for c = b; c <= max; c++ {
				if c*c == a*a+b*b && a < b && b < c {
					x = append(x, Triplet{a, b, c})
					p++
				}

			}
		}
	}
	return x
}

func Sum(n int) []Triplet {
	x := Range(1, n)
	var p []Triplet
	l := len(x)

	for i := 0; i < l; i++ {
		if x[i][0]+x[i][1]+x[i][2] == n {
			p = append(p, Triplet{x[i][0], x[i][1], x[i][2]})
		}
	}
	return p
}

