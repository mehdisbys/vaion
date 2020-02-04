package main

import (
	"fmt"
	"math/big"
)

/*
1. The Fibonacci sequence is defined as a sequence of integers starting with 1 and 1, where each
subsequent value is the sum of the preceding two. I.e.
f(0) = 1
f(1) = 1
f(n) = f(n-1) + f(n-2) where n >= 2
Write a program in a language of your choice to calculate the sum of the first 100 even-valued
Fibonacci numbers
*/

// Note : stops working after n=30
// Golang highest available integer is 9223372036854775807
// See BigEvenFibonnacci function below
func EvenFibonacci(n int) []int64 {
	f := []int64{1, 1}
	found := 0
	evenFib := make([]int64, n)

	for i := 2; found < n; i++ {
		val := f[i-1] + f[i-2]
		f = append(f, val)

		if isEven(val) {
			evenFib[found] = val
			found++
		}
	}
	return evenFib
}

func isEven(n int64) bool {
	return n%2 == 0
}

func isBigEven(n *big.Int) bool {
	if n == nil {
		return false
	}
	return n.Bit(0) == 0
}

//From :  https://github.com/golang/go/issues/30943
func BigEvenFibonnacci(n int) []*big.Int {
	f := []*big.Int{big.NewInt(1), big.NewInt(1)}
	found := 0
	evenFib := make([]*big.Int, n)

	for i := 2; found < n; i++ {

		val := fib(i)
		f = append(f, val)

		if isBigEven(val) {
			evenFib[found] = val
			found++
		}
	}
	return evenFib
}

func fib(n int) *big.Int {

	left := big.NewInt(1)
	right := big.NewInt(1)

	helper1 := big.NewInt(0)
	helper2 := big.NewInt(0)

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1.Mul(big.NewInt(2), right)
		helper1.Sub(helper1, left)
		helper1.Mul(left, helper1)

		helper2.Mul(right, right)
		left.Mul(left, left)
		helper2.Add(helper2, left)

		if bin[i] == '0' {
			left.Set(helper1)
			right.Set(helper2)
		} else {
			left.Set(helper2)
			right.Add(helper1, helper2)
		}
	}
	return left
}
