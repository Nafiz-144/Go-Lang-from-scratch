/*********** Compile Phase ************
- go run main.go  => Compiles and runs (temporary binary)
- go build main.go => Compiles and creates executable file
*************************************/

package main

// fmt package is used for input/output operations
import "fmt"

// User is a struct (custom data type)
// It groups related data together
type User struct {
	Name string // stores user's name
	Age  int    // stores user's age
}

// Normal function that takes a User struct as parameter
func printUserDetails(usr User){
	fmt.Println("Name:", usr.Name) // access Name field
	fmt.Println("Age:", usr.Age)   // access Age field
}

// ******** Receiver Function (Method) ********
// This function is attached to User type
// It can be called using object.method()
func (usr User) printDetails(){
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// Another method with a parameter
// usr1 is the receiver variable (instance of User)
func (usr1 User) call(a int){
	fmt.Println("Name:", usr1.Name) // access Name from struct
	fmt.Println("Value:", a)        // print passed argument (not actual Age)
}

func main() {

	// Declare two variables of type User
	var user1, user2 User

	// Initialize user1
	user1 = User{
		Name: "Nafiz",
		Age:  19,
	}

	// Initialize user2
	user2 = User{
		Name: "Habib",
		Age:  30,
	}

	// Call normal function (pass struct as argument)
	printUserDetails(user1)
	printUserDetails(user2)

	// Call method (receiver function)
	user1.printDetails()

	// Call method with parameter
	user2.call(10)
}