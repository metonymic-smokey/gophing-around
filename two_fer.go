package twofer
//package main

import (
	"fmt"
)

func ShareWith(name string) string {
	if name=="" {
		return "One for you, one for me."
	}	

	var res string
	
	res = "One for "+name+", one for me."
	return res
}

func twofer() {
	fmt.Println(ShareWith("Zaphod"))

}
