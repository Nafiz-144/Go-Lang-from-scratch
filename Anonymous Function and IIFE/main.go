package main

import "fmt"

// 1. Fixed: Use 'var' for package-level variables.
// Logic like 'if' moved into a function.
var e = 10
// Standard Function
func add(x int, y int) {
	sum := x + y
	fmt.Println("From standard function:", sum)
}

func main() {
	// Moved the 'if' logic inside main so it can execute
	if e > 5 {
		fmt.Println("Expression:", e)
	}

	add(4, 7)

	// Immediately Invoked Function Expression (IIFE)
	func(a int, b int) {
		res := a + b
		fmt.Println("From anonymous function:", res)
	}(1, 2)
}

func init() {
	fmt.Println("Welcome to the system....!")
}