package main

//linked list
import "fmt"

func calculate() (result int) {
	fmt.Println("first:", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer:", result)

	}
	defer show()
	result = 5
	fmt.Println("Second", result)
	return

}

func calc() int {
	result := 0
	fmt.Println("first:", result)
	show := func() {
		result = result + 10
		fmt.Println("defer:", result)
	}
	defer show()
	result = 5
	fmt.Println("Second:", result)
	return result

}

func main() {
	a := calculate()
	fmt.Println("main first:", a)
	b := calc()
	fmt.Println("main second:", b)

}
