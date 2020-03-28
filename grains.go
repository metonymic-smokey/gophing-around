package grains


import (
	"fmt"
	"errors"
	"math"
	)

func Square(n int) (uint64,error) {
//	var err error
	if(n > 64 || n <= 0){
		return 0,errors.New("Value exceeded")
	} else {
		return uint64(math.Pow(float64(2), float64(n-1))),nil
	}
}

func Total() uint64 {
	
	return 18446744073709551615
}

func grains() {
	var n int
	var t uint64	

	fmt.Println("Enter n: ")	
	fmt.Scanln(&n)

	x,_ := Square(n)
	fmt.Println(x)	

	t = Total()
	fmt.Println(t)
}
