package main

import (
	"log"
	"strconv"
	"strings"
)


/*
3. Write a function in a language of your choice which, when passed a positive integer returns
true if the decimal representation of that integer contains no odd digits and otherwise returns
false.
 */

func DecimalRepresentation(n int64) bool {
	s := strconv.FormatInt(n, 10)
	sNumbers := strings.Split(s, "")

	for _, v := range sNumbers {
		val, err := strconv.Atoi(v)

		if err != nil {
			log.Print(err)
			continue
		}

		if !isEven(int64(val)) {
			return false
		}
	}
	return true
}
