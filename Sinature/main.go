package main

import "fmt"

type nafiz func(a int)

func printvalue(p nafiz) {
	p(20)
}
func main() {

	logFn := func(num int) {
		fmt.Println(num)

	}
	printvalue(logFn)
}
