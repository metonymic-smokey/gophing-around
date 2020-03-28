package isogram

import (
	"strings"
	)


func IsIsogram(word string) bool {

	freq := make(map[rune]int)
	word= strings.Replace(strings.ToLower(word), "-", "", -1)
	word = strings.Replace(strings.ToLower(word), " ", "", -1)

	for _, l := range word {
		freq[l]++

		if freq[l] > 1 {
			return false
		}
	}
	
	return true
}

