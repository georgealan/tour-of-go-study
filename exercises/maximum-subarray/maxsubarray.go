package main

import "fmt"

/*
TODO Given an integer array nums, find the subarray which has the largest sum and return its sum.
*/
func main() {
	s1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	s2 := []int{1}
	s3 := []int{5, 4, -1, 7, 8}

	fmt.Println(maxSubArray(s1))
	fmt.Println(maxSubArray(s2))
	fmt.Println(maxSubArray(s3))
}

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum = currentSum + nums[i]
		if currentSum > maxSoFar {
			maxSoFar = currentSum
		}
	}

	return maxSoFar
}
