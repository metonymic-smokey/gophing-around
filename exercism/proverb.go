package proverb

import (
	"fmt"
)

func Proverb(rhyme []string) []string {
	var res []string
	var x string
	var i, l int
	//c -> counter
	l = len(rhyme)
	
	if l == 0 {
		return []string{}
	}


	for i = 0; i < l-1; i++ {
		x = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		res = append(res, x)
	}

	res = append(res, "And all for the want of a "+rhyme[0]+".")

	return res
}

func proverb() {

	r := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	fmt.Println(Proverb(r))
}
