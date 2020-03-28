package triangle

import (
	"fmt"
	"math"
)

const (
	NaT = 0
	Equ = 3
	Iso = 2
	Sca = 1
)

type Kind = int

func KindFromSides(a, b, c float64) Kind {

	if a+b+c < 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a == b && a == c && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}

	return NaT
}
