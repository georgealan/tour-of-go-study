package main

import "fmt"

func main() {
	arr1 := []int{2, 11, 7, 15}
	arr2 := []int{3, 2, 4}
	arr3 := []int{3, 3}

	fmt.Println(twoSum(arr1, 9))
	fmt.Println(twoSum(arr2, 6))
	fmt.Println(twoSum(arr3, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums { // index and number

		// Basic maths: 2 + 7 = 9 - 7 = 2 + 7 = 9 - 7 = 2
		if v, found := m[target-num]; found {
			return []int{v, i} // map value that we found with this key m[target - num] followed by index
		}

		m[num] = i
	}
	return nil
}
