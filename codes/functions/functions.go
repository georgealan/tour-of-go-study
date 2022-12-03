package main

import "fmt"

func main() {
	fmt.Println(add(10, 90))
	a, b := swap("Alan", "George")
	fmt.Println(a, b)
	fmt.Println(split(25))

	/*
		A defer statement defers the execution of a function until the surrounding function returns.
		The deferred call's arguments are evaluated immediately, but the function call is not executed
		until the surrounding function returns.
	*/
	defer fmt.Println("world!") //  Will be executed after function bellow
	fmt.Println("Hello")        // Execute before the defer function above

	stackingDefer()
}

// when two or more consecutive named function parameters share a type you can shorten then
func add(x, y int) int {
	return y + y
}

// A function can return more than a number of results
func swap(x, y string) (string, string) {
	return y, x
}

/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, as with the example shown here.
They can harm readability in longer functions.
*/
func split(sum int) (x, y int) { // Return values named
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
Deferred function calls are pushed onto a stack. When a function returns,
its deferred calls are executed in last-in-first-out order.
*/
func stackingDefer() {
	fmt.Println("1º Before loop with defer statement")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v, ", i) // Will print in descending order
	}
	fmt.Println("2º After loop with defer statement")
}

/*
TODO Continue in Basics: More types: structs, slices, and maps / Pointers
 Link: https://go.dev/tour/moretypes/1
*/
