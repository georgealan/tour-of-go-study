package main

import "fmt"

/*
The var statement declare a list of variables, as in function argument lists, the
type is last, it can be at package or function level.
*/
var c, python, java bool
var s, j int = 1, 2 // Here type can be omitted
// d := 3 // This can't happen, because the := assigment can't be used outside a function

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var verify, test, think = true, false, "More"
	fmt.Println(s, j, verify, test, think)

	d := 3 // Inside a function the := short assigment can be used in place of a var declaration
	fmt.Println(d)

	fmt.Println("Variables declared without an explicit initial value are given their zero value")
	var sw int
	var fl float64
	var tg bool
	var name string
	fmt.Printf("%v %v %v %q\n", sw, fl, tg, name)

	/*
		Type inferences
		When declaring a variable without specifying an explicit type (either by using the := syntax
		or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	*/
	qw := 42
	qe := 3.1415
	qr := 0.867 + 0.5i
	fmt.Printf("Type of qw: %T, Type of qe: %T, Type of qr: %T\n", qw, qe, qr)

}
