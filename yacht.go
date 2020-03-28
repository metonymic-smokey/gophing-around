package yacht

import (
	"sort"
)

func Score(dice []int, cat string) int {
	sorted := dice

	sort.Ints(sorted)

	if cat == "little straight" {
		var c int
		for i := 0; i < len(sorted); i++ {
			if sorted[i] == i+1 {
				c += 1
			}
		}

		if c == len(sorted) {
			return 30
		}
		return 0
	}

	if cat == "big straight" {
		var c int
		for i := 0; i < len(sorted); i++ {
			if sorted[i] == i+2 {
				c += 1
			}
		}

		if c == len(sorted) {
			return 30
		}
		return 0
	}

	if cat == "ones" {
		c := 0
		for i := 0; i < len(sorted); i++ {
			if dice[i] == 1 {
				c++
			}
		}
		return c
	}

	if cat == "twos" {
		c := 0

		for i := 0; i < len(sorted); i++ {
			if dice[i] == 2 {
				c++
			}
		}
		return c * 2
	}

	if cat == "threes" {
		c := 0

		for i := 0; i < len(sorted); i++ {
			if dice[i] == 3 {
				c++
			}
		}
		return c * 3
	}

	if cat == "fours" {
		c := 0

		for i := 0; i < len(sorted); i++ {
			if dice[i] == 4 {
				c++
			}
		}
		return c * 4
	}

	if cat == "fives" {
		c := 0

		for i := 0; i < len(sorted); i++ {
			if dice[i] == 5 {
				c++
			}
		}
		return c * 5
	}

	if cat == "sixes" {
		c := 0

		for i := 0; i < len(sorted); i++ {
			if dice[i] == 6 {
				c++
			}
		}
		return c * 6
	}

	if cat == "choice" {
		var c int
		for i := 0; i < len(sorted); i++ {
			c += sorted[i]
		}
		return c
	}

	if cat == "full house" {
		var s int

		if sorted[0] == sorted[1] && sorted[1] == sorted[2] && sorted[2] != sorted[3] && sorted[3] == sorted[4] {
			for i := 0; i < len(sorted); i++ {
				s += sorted[i]
			}
		}

		if sorted[0] == sorted[1] && sorted[1] != sorted[2] && sorted[3] == sorted[2] && sorted[3] == sorted[4] {
			for i := 0; i < len(sorted); i++ {
				s += sorted[i]
			}
		}
		return s
	}

	//four of a kind
	if cat == "four of a kind" {
		if sorted[0] != sorted[1] {
			if sorted[1] == sorted[2] {
				return 4 * sorted[1]
			}
			return 0
		}

		if sorted[3] != sorted[4] {
			if sorted[2] == sorted[3] {
				return 4 * sorted[2]
			}
			return 0
		}

		var c int
		for i := 0; i < len(sorted); i++ {
			if sorted[i] == sorted[0] {
				c++
			}
		}

		if c == len(sorted) {
			return 4 * sorted[0]
		}

	}

	if cat == "yacht" {
		var c int
		for i := 0; i < len(sorted); i++ {
			if sorted[i] == sorted[0] {
				c++
			}
		}
		if c == len(dice) {
			return 50
		}

		return 0
	}

	return 0

}

