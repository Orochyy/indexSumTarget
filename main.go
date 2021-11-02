package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := []int{}
	for e1, n1 := range nums {
		for e2, n2 := range nums {
			if n1+n2 == target {
				result = []int{e2, e1}
			}
		}
	}
	return result
}

func main() {
	mass := []int{1, 5, 3, 6}
	target := 7

	twoSum := twoSum(mass, target)
	fmt.Println(twoSum)

}
