package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
4. Write a function in a language of your choice which, when passed a decimal digit X, returns the
value of X+XX+XXX+XXXX. E.g. if the supplied digit is 3 it should return 3702
(3+33+333+3333).
*/

func XFormula(n int) (int, error) {
	if n < 0 || n > 9 {
		return 0, errors.New("input should be a decimal digit")
	}

	values := make([]string, 4)

	s := fmt.Sprintf("%d", n)
	v := s

	for i := 0; i < 4; i++ {
		values[i] = v
		v += s
	}

	res := 0
	for _, sval := range values {
		ival, err := strconv.Atoi(sval)
		if err != nil {
			return 0, err
		}
		res += ival
	}
	return res, nil
}
