package allyourbase

import (
	"errors"
	"math"
)

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func digits(n int) []int {
	q := n
	var dig []int

	for q != 0 {
		r := q % 10
		dig = append(dig, r)
		q /= 10
	}
	dig = reverse(dig)
	return dig
}

func ConvertToBase(ib int, dig []int, ob int) ([]int, error) {

	var dec int //decimal rep of each ip
	var ind, i int
	var res []int

	if ib < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}

	if ob < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	if ib < 0 || ob < 0 {
		return []int{}, errors.New("bases cannot be negative")
	}

	for i = 0; i < len(dig); i++ {
		if dig[i] < 0 || dig[i] >= ib {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	if len(dig) == 1 && dig[0] == 0 {
		return []int{0}, nil
	}

	//empty list
	if len(dig) == 0 {
		return []int{0}, nil
	}

	for i = len(dig) - 1; i >= 0; i-- {
		dec += (dig[i] * int(math.Pow(float64(ib), float64(ind))))
		ind += 1
	}

	if dec == 0 {
		return []int{0}, nil
	}

	x := digits(dec)

	if ob == 10 {
		return x, nil
	}

	q := dec

	for q != 0 {
		r := q % ob
		res = append(res, r)
		q /= ob
	}

	res = reverse(res)
	return res, nil
}

