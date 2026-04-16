package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 ,user2 User
	user1 = User{
		Name: "Nafiz",
		Age:  19,
	}
	user2 = User{
		Name: "Habib",
		Age:  30,
	}

	fmt.Println("Name:",user1.Name)
	fmt.Println("Age:",user1.Age)
	fmt.Println("Name:",user2.Name)
	fmt.Println("Age:",user2.Age)

}
