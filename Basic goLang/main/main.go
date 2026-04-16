package main

import "fmt"

func main() {
	const country="Bangladesh" 
	var name, batch, section string
	var id, gpa int
	fmt.Print("Enter Your Details:\n")
	fmt.Print("Name:")
	fmt.Scan(&name)
	fmt.Print("Batch:")
	fmt.Scan(&batch)
	fmt.Print("Section:")
	fmt.Scan(&section)
	fmt.Print("ID:")
	fmt.Scan(&id)
	fmt.Print("GPA:")
	fmt.Scan(&gpa)
	

    fmt.Printf("Student Name:%q from:%s\n" ,name,country)
	fmt.Printf("Student Batch:%s\n",batch)
	fmt.Println("Student Section:",section)
	fmt.Println("Student ID:",id)
	fmt.Println("Student GPA:",gpa)
	
	fmt.Println("GPA Number conversion")
	fmt.Printf("Binary Number:%b\n",gpa)
	fmt.Printf("Octal Number:%o\n",gpa)
  


}
