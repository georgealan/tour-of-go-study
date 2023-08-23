package packagesandmore

import "fmt"

func add2(x, y int) int {
	return x + y
}

func SumTwoNumbers2() {
	// When two or more consecutive named function parameters share a type, you can omit the type from
	// all but the last.
	fmt.Println(add2(42, 13))
}
