package main

import "fmt"

// Pointer -> stores the address of a variable in memory (RAM)

// Function that takes a pointer to an array
// Instead of copying the whole array, it works with the original array
func print(numbers *[3]int) {
	fmt.Println(numbers) // prints the address/reference (default behavior)
	// To print actual values, you could use: fmt.Println(*numbers)
}

// -------- Struct Definition --------
// User is a custom data type that groups related data
type User struct {
	name string // user's name
	age  int    // user's age
}

// -------- Pass by Value --------
// This function receives a copy of the User struct
// Any modification here will NOT affect the original data
func printDetails(u User) {
	fmt.Println("Name:", u.name)
	fmt.Println("Age:", u.age)
}

// -------- Pass by Reference (Pointer) --------
// This function receives a pointer to User
// It works with the original object in memory
func pointerview(p *User) {
	fmt.Println("Name:", p.name) // access using pointer
	fmt.Println("Age:", p.age)
}

func main() {

	// -------- Pointer Basic Example --------
	// x := 20
	// p := &x          // p stores address of x
	// fmt.Println("x:", x)
	// *p = 30          // change value at address
	// fmt.Println("x:", x) // x is now updated (30)
	// fmt.Println("Address:", p)
	// fmt.Println("Value at address:", *p)

	// -------- Array Pointer Example --------
	arr := [3]int{1, 2, 3}

	// Passing address of array (not copying full array)
	print(&arr)

	// -------- Struct Example --------
	user1 := User{
		name: "nk",
		age:  6,
	}

	// Pass by value (copy)
	printDetails(user1)

	// Pass by reference (pointer)
	pointerview(&user1)
}
