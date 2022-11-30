package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 2, 1, 1, 3}                 // true
	arr2 := []int{1, 2}                             // false
	arr3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0} // true
	arr4 := []int{1, 2, 2, 1, 1, 2, 3}              // false
	arr5 := []int{-1050, 2, 2, 2500, 1, 1, 2, 3}    // Shall not pass, error

	fmt.Println(uniqueOccurrences(arr1))
	fmt.Println(uniqueOccurrences(arr2))
	fmt.Println(uniqueOccurrences(arr3))
	fmt.Println(uniqueOccurrences(arr4))
	fmt.Println(uniqueOccurrences(arr5))
}

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int, len(arr))
	for _, num := range arr {
		freq[num]++
	}

	hasSeen := make(map[int]bool, len(freq))
	for _, vFreq := range freq {
		if hasSeen[vFreq] {
			return false
		}
		hasSeen[vFreq] = true
	}
	return true
}
