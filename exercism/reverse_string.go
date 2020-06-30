package reverse

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

func reverse() {
	var word string

	fmt.Println("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word = scanner.Text()

	fmt.Println(word)

	res := Reverse(word)
	fmt.Println(res)
}
