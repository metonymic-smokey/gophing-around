package queenattack

import (
	"fmt"
	"strconv"
	"errors"
	"math"
)

//return (attack,ok)
func CanQueenAttack(w, b string) (bool, error) {

	if w == "" || b=="" ||w == b || w[0] > 'h' || b[0] > 'h' {
		return false, errors.New("invalid")
	}

	wn,we := strconv.Atoi(string(w[1]))
	bn,be :=strconv.Atoi(string(b[1]))

	if we != nil || be != nil {
		return false,errors.New("empty string")
	}

	if wn < 1 || wn > 8 || bn < 1 || bn > 8 {
		return false,errors.New("invalid")
	}
	
	if wn == bn || w[0] == b[0] {
		return true,nil
	}

	if math.Abs(float64(int(w[0]) - int(b[0]))) == math.Abs(float64(wn - bn)) {
		return true,nil
	}

	return false,nil
	
}

}
