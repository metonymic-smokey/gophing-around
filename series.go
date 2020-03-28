package series

import (
	"fmt"
)

func All(n int, s string) []string {
	var i int
	var res []string
	fmt.Println(len(res))

	for i = 0; i < len(s)-n+1; i = i + 1 {
		res = append(res, s[i:i+n])
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	var res string

	res = s[0:n]
	return res
}

func First(n int, s string) (string, bool) {
	if n < 0 || n > len(s) {
		return s, false
	}

	return s[0:n], true

}

func series() {
	fmt.Println(UnsafeFirst(3, "49142"))

}
