package main

import (
	"fmt"
	"math"
)

/*
Go has only one looping construct, the for loop
*/

func main() {
	sumFromRange(10)
	sumFromRange2(10)

	fmt.Println("Exercise: Loops and Functions")
	fmt.Println("My own function", sqrt(789))
	fmt.Println("Original expected result", math.Sqrt(789))
}

func sumFromRange(x int) {
	sum := 0
	for i := 0; i < x; i++ {
		sum += i
	}
	fmt.Printf("Sum result: %v\n", sum)
}

func sumFromRange2(x int) {
	sum := x
	for sum < 100 { // The init and post statements are optional.
		sum += sum
	}
	fmt.Printf("Sum without init and post statement: %v\n", sum)
	// This form is like a while statement
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		if z == math.Sqrt(x) {
			fmt.Printf("Loop stoped here in indice = %v and value of z is = %v\n", i, z)
			break
		}
		fmt.Printf("In the indice %v, the variable z = %v\n", i, z)
	}
	return z
}
