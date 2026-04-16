package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter the Number:")
	fmt.Scan(&number)
	if number%2==0{
		fmt.Print("Number Is Even...")
	} else {
		fmt.Print("Number Is Odd...")
	}

}