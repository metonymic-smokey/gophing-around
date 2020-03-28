package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(s string, n int) (int, error) {
	var i, pro, j, max int
	pro = 1

	if n > len(s) {
		return 0, errors.New("n too large")
	}

	if n < 0 {
		return 0, errors.New("n less than 0")
	}

	for i = 0; i < 1+len(s)-n; i++ {
		pro = 1
		for j = 0; j < n; j++ {
			x, e := strconv.Atoi(string(s[i+j]))
			if e != nil {
				return 0, errors.New("not a digit")
			}
			pro *= x
		}

		if pro > max {
			max = pro
		}
	}

	return max, nil
}
