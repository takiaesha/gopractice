package main

import "fmt"

func slice() {
	var x []int

	fmt.Println(x) //null slice

	//built-in function
	y := make([]int, 5)
	fmt.Println(y) // [0 0 0 0 0]

	z := make([]int, 5, 10)
	fmt.Println(z)

	//low high
	ab := [7]float64{1, 2, 3, 4.7, 5, 6.9, 7} //array
	fmt.Println(ab)

	p := ab[0:5]
	fmt.Println(p) //1 2 3 4.7 5
	q := ab[:]
	fmt.Println(q) //whole array
	r := ab[5:len(ab)]
	fmt.Println(r) /// 6.9 7

	//slice append & copy
	slice1 := []int{123, 100, 560, 670}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, "\n", slice2)

	copy(slice1, slice2)
	fmt.Println("Slice1 = ", slice1)
	fmt.Println("Slice2 = ", slice2)
}
