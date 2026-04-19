package main

import "fmt"

func main() {
	// original array (fixed size = 6)
	arr := [6]string{"this", "is", "a", "go", "interview", "question"}
	fmt.Println("original arr:", arr)

	// slicing from index 1 to 4 (excluding 4)
	// s -> ["is", "a", "go"]
	// ptr = &arr[1]
	// length = 4 - 1 = 3
	// capacity = from index 1 to end of array = 6 - 1 = 5
	s := arr[1:4]
	fmt.Println("s:", s)

	// slicing s from index 1 to 2 (excluding 2)
	// s1 -> ["a"]
	// ptr = &arr[2] (because s starts from arr[1])
	// length = 2 - 1 = 1
	// capacity = from index 2 to end of array = 6 - 2 = 4
	s1 := s[1:2]
	fmt.Println("s1:", s1)

	// printing slice properties
	fmt.Println("S1 Slice Length:", len(s1))   // 1
	fmt.Println("S1 Slice Capacity:", cap(s1)) // 4

	// -------------------------------
	// slice literal (dynamic array)
	// underlying array is created automatically
	// length = 3, capacity = 3
	a := []int{1, 2, 3}
	fmt.Println("slice:", a, "Length:", len(a), "capacity:", cap(a))

	// -------------------------------
	// slice using make()
	// make(type, length)
	// creates slice with zero values
	m := make([]int, 3) // [0 0 0]
	m[0] = 5
	fmt.Println("Slice", m, "Length:", len(m), "capacity:", cap(m)) // cap = len = 3

	// make(type, length, capacity)
	// length = 3, capacity = 5 (extra space reserved)
	mc := make([]int, 3, 5)
	fmt.Println("Slice", mc, "Length:", len(mc), "capacity:", cap(mc))

	// -------------------------------
	// nil (empty) slice
	// initially length = 0, capacity = 0, no underlying array
	var x []int

	fmt.Println("append creates underlying array automatically")
	x = append(x, 1) // [1]
	fmt.Println("slice x:", x, "Length:", len(x), "capacity:", cap(x))

	x = append(x, 2) // [1,2]
	fmt.Println("slice x:", x, "Length:", len(x), "capacity:", cap(x))

	// capacity growth happens here (usually doubles)
	x = append(x, 3) // [1,2,3]
	fmt.Println("slice x:", x, "Length:", len(x), "capacity:", cap(x))

	// y points to SAME underlying array as x (no copy)
	y := x
	fmt.Println("slice y:", y, "Length:", len(y), "capacity:", cap(y))

	// append may create NEW array if capacity exceeded
	x = append(x, 4) // [1,2,3,4]
	fmt.Println("slice x:", x, "Length:", len(x), "capacity:", cap(x))

	// y still refers to OLD array
	// so this append works on old backing array
	y = append(y, 5) // [1,2,3,5]
	fmt.Println("slice y:", y, "Length:", len(y), "capacity:", cap(y))

	// final result shows they are now different
	fmt.Println("x:", x) // [1,2,3,4]
	fmt.Println("y:", y) // [1,2,3,5]
	//slice new underling array rule=>1024->100%  increase ,1024<25% increase

	p := []int{1, 2, 3, 4, 5}
	p = append(p, 6) //1 2 3 4 5 6
	p = append(p, 7)
	q := p[4:]
	o := changeslice(q)
	fmt.Println(p)
	fmt.Println(o)

}
func changeslice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

//p=10,6,7,11
