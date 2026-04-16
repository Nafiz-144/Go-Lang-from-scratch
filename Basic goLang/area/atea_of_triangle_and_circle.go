package main

import "fmt"
func main() {
	var height, weight, radius float64
	fmt.Printf("Enter the Height:")
	fmt.Scan(&height)
	fmt.Printf("Enter the Weight:")
	fmt.Scan(&weight)
	fmt.Printf("Enter the Radius:")
	fmt.Scan(&radius)
    pi:=3.1416
	area:=pi*float64(radius)*float64(radius)
	fmt.Printf("Area of circle :%v\n",area)

	areat :=0.5*height*weight
	fmt.Printf("Area of Triangle :%v\n",areat)


}