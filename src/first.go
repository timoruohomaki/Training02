package main

import "fmt"

// OPTION 1: PASSING ARRAY POINTERS

/*

func foo(x *[3]int) int {
	(*x)[0] = (*x)[0] + 1
	return x[0]
}

func main() {
	a := [3]int{1, 2, 3}
	foo(&a)
	fmt.Println(a)

}

*/

// OPTION 2: PASSING SLICES INSTEAD

func foo(sli []int) int {
	sli[0] = sli[0] + 1
	return sli[0]
}

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Println(a)
}
