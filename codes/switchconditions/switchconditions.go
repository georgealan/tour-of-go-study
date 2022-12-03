package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	printOsName()

	today := time.Now().Weekday()
	whenIsSaturday(today)

	fmt.Print("Hello George! ")
	greetings()

}

/*
A switch statement is a shorter way to write a sequence of if - else statements.
It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case,
not all the cases that follow. In effect, the break statement that is needed at the end of each case in those
languages is provided automatically in Go. Another important difference is that Go's switch cases need not be
constants, and the values involved need not be integers.
*/
func printOsName() {
	fmt.Print("Go is running on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux ")
	default:
		// freebsd, openbsd, plan9, windows
		fmt.Printf("%s.\n", os)
	}
}

func whenIsSaturday(today time.Weekday) {
	fmt.Print("Is today Saturday? ")

	switch time.Saturday {
	case today + 0:
		fmt.Println("- Today are.")
	case today + 1:
		fmt.Println("- Is Tomorrow.")
	case today + 2:
		fmt.Println("- Will be in two days.")
	case today + 3:
		fmt.Println("- Will be in three days.")
	case today + 4:
		fmt.Println("- Will be in four days")
	case today + 5:
		fmt.Println("- Will be in five days")
	case today + 6:
		fmt.Println("- It was yesterday moron")
	}
}

// Switch without condition
func greetings() {
	actualHour := time.Now()

	switch {
	case actualHour.Hour() < 12:
		fmt.Println("Good morning!\n")
	case actualHour.Hour() < 17:
		fmt.Println("Good afternoon!\n")
	default:
		fmt.Println("Good evening!\n")
	}
}
