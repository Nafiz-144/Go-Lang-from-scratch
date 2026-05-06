package main

import (
	"fmt"
	"os"
)

// only signature of the function
type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}
type bankuser interface {
	WithdrowMoney(amount float64) float64
}

type user struct {
	Name   string
	Age    int
	Salary float64
}

func (usr user) PrintDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Salary:", usr.Salary)

}
func (obj user) ReceiveMoney(amount float64) float64 {

	obj.Salary = obj.Salary + amount
	return obj.Salary
}

func (obj user) WithdrowMoney(amount float64) float64 {

	obj.Salary = obj.Salary - amount
	return obj.Salary
}

func main() {
	var usr1 People
	usr1 = user{
		Name:   "Nafiz",
		Age:    24,
		Salary: 0.1,
	}
	usr1.PrintDetails()
	usr1.ReceiveMoney(45)

	var usr2 bankuser
	usr2 = user{
		Name:   "karim",
		Age:    655,
		Salary: 6401,
	}
	usr2.WithdrowMoney(84)
	odj, err := usr2.(user)
	if !err {
		fmt.Println("sorry")
		os.Exit(1)
	}
	odj.WithdrowMoney(84)

}
