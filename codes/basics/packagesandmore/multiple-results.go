package packagesandmore

import "fmt"

// A function can return any number of results.
// This function below returns two strings
func swap(x, y string) (string, string) {
	return y, x
}

func PrintMultipleResults() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
