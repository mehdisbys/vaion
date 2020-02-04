package main


/* 2. Write a function in a language of your choice which, when passed two sorted arrays of integers
returns an array of any numbers which appear in both. No value should appear in the returned
array more than once.
*/


func CommonValues(a, b []int) []int {
	res := map[int]bool{}

	for _, j := range a {
		for _, y := range b {
			if j == y {
				res[j] = true
			}
		}
	}

	var common []int
	for k := range res {
		common = append(common, k)
	}
	return common
}

