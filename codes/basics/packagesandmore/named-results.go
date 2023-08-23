package packagesandmore

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func PrintSplit() {
	fmt.Println(split(17))
}
