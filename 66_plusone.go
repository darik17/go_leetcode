package main

import "fmt"

func main() {
	var digits = []int{8, 9, 9}
	fmt.Print(plusOne(digits))
}

func plusOne(digits []int) []int {
	last := digits[len(digits)-1]
	if last < 9 {
		digits[len(digits)-1] = last + 1
	} else {
		el := 1
		for i := len(digits) - 1; i != -1; i-- {
			if digits[i]+el > 9 {
				digits[i] = 0
				el = 1
			} else {
				digits[i] += 1
				el = 0
				return digits
			}
		}
		if el == 1 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
