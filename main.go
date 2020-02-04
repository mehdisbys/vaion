package main

import (
	"fmt"
)

func main() {
	fmt.Println(BigEvenFibonnacci(100)) // see also EvenFibonacci()

	fmt.Println(CommonValues([]int{3, 4, 5, 6, 7, 124, 887, 1221}, []int{0, 4, 5, 12, 887}))

	fmt.Println(DecimalRepresentation(222))

	fmt.Println(XFormula(3))
}
