package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(s string) Frequency {
	c := make(Frequency)
	s = strings.ToLower(s)
	s = strings.Replace(s, ",", " ", -1)

	reg, _ := regexp.Compile("[^a-zA-Z0-9' ]+")
	s = reg.ReplaceAllString(s, "")

	words := strings.Fields(s)
	if len(words) == 1 {
		c[words[0]] = 1
		return c
	}

	for _, word := range words {
		if strings.Count(word, "'") == 2 {
			word = strings.Replace(word, "'", "", -1)
			c[word] += 1
		} else if strings.Count(word, "'") <= 1 {
			c[word] += 1
		}
	}

	return c
}
