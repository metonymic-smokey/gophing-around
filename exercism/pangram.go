package pangram

import (
	"fmt"
	"strings"
)

func IsPangram(s string) bool {
	s = strings.ToLower(s)

	var i rune
	var letters []rune

	for i = 'a'; i <= 'z'; i++ {
		letters = append(letters, (i))
	}
	//	fmt.Println(letters)

	c:=0

	for i='a';i <='z';i++ {
		if strings.Contains(s,string(i)) == true {
			c++
		}	
	}

	if c == 26 {
		return true
	}

	return false
}

func pangram() {
	fmt.Println(IsPangram("The quick brown fox jumps over the lazy dog"))

}
