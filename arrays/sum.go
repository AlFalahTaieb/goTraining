package main

func sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}
