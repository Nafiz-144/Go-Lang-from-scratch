package main

import "fmt"

// Define a custom function type named 'nafiz'
// This type represents any function that takes an int as input and returns nothing
type nafiz func(a int)

// printvalue function takes a parameter 'p' of type nafiz (i.e., a function)
// It calls that function and passes the value 20
func printvalue(p nafiz) {
	p(20) // invoke the function with argument 20
}

func main() {

	// Define an anonymous function and assign it to logFn
	// This function takes an integer and prints it
	logFn := func(num int) {
		fmt.Println(num) // print the received value
	}

	// Pass the function 'logFn' as an argument to printvalue
	// Since logFn matches the 'nafiz' function signature, it's valid
	printvalue(logFn)
}
