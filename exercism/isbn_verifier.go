package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(s string) bool {
	var sum, ind int
	ind = 10

	s = strings.Replace(s, "-", "", -1)

	if len(s) != 10 {
		return false
	}

	for _, i := range s {
		if int(i) >= 48 && int(i) <= 57 {
			c, _ := strconv.Atoi(string(i))
			sum += (c * ind)
			ind--
		} else {
			if int(i) == 88 {
				if strings.IndexRune(s, i) == 0 || strings.IndexRune(s, i) == 9 {
					sum += (10 * ind)
					ind--
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	if sum%11 == 0 {
		return true
	}

	return false

}
