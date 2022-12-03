package main

import (
	"fmt"
	"math/cmplx"
)

/*
Go's basic types are:

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point

float32 float64
complex64 complex128
*/

var (
	Test   bool       = false
	MaxInt uint64     = 1<<64 - 1
	Z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Test, Test)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Z, Z)

	/*  Basic Convertions
	The expression T(v) converts the value v to the type T
	Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	*/
	var i int = 96
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("%v %v %v", i, f, u)
}
