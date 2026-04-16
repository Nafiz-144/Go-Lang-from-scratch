package main

import (
	"fmt"
	"nafiz/mathlib"
)

/*
go mod init nafiz
*/

func main() {
	fmt.Println("Showing Package Scope...")
	mathlib.Add(5,5)
	

}