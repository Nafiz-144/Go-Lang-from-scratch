package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Printf("Enter Three Number:")
	fmt.Scan(&num1,&num2,&num3)
	if(num1>num2)&&(num1>num3){
		fmt.Printf("Number 1....that is %v\n",num1)
	}else if(num2>num1)&&(num2>num3){
		fmt.Printf("Number 2 ! .... that is %v\n",num2)
	} else 	if(num3>num1)&&(num3>num2){
		fmt.Printf("Number 3! .... that is %v\n",num3)
	}else {
		fmt.Printf("They are Equal.0.....")
	}

}