package main

import "fmt"
var (

	a=10
)
func add(x int ,y int){
	z:=x+y
	fmt.Println(z)
	
}


func init()  {
		fmt.Println("Hello") 
		
	}
func main() {

add(5,4)
add(a,3)
	
}

//code segment(all function)->data segment(global memory)
// ->stack ->heap-GC(garbage collector)