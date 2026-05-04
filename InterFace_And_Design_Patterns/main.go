package main

import "fmt"

type user struct {
	Name   string
	Age    int
	Salary float64
}

//only signature of the function
type People interface {
}

func (usr user) PrintDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Salary:", usr.Salary)

}

func main() {
	usr1 := user{
		Name:   "Nafiz",
		Age:    24,
		Salary: 0.1,
	}
	usr1.PrintDetails()

}
