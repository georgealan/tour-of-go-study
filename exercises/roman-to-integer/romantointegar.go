package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	answer := 0
	number := 0

	word := strings.Split(s, "")

	for i := len(word) - 1; i >= 0; i-- {
		switch word[i] {
		case "I":
			number = 1
		case "V":
			number = 5
		case "X":
			number = 10
		case "L":
			number = 50
		case "C":
			number = 100
		case "D":
			number = 500
		case "M":
			number = 1000
		}

		if 4*number < answer {
			answer -= number
		} else {
			answer += number
		}
	}
	return answer
}
