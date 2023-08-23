package packagesandmore

import (
	"fmt"
	"math"
)

func TestImportName() {
	/*
		In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi,
		which is exported from the math package.
	*/
	fmt.Println(math.Pi)
}
