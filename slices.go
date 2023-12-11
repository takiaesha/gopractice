package main

import "fmt"

func print1(s []int) {
	s[0] = 12
	s = append(s, 10)
}
func print2(s *[]int) {
	(*s)[0] = 70
	*s = append(*s, 90)
}
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

	//pass by value
	myslice := []int{1, 2, 3, 4, 5}
	print1(myslice)
	fmt.Println(myslice) //12 2 3 4 5

	//pass by reference
	myslice2 := []int{10, 20, 30, 40, 50}

	print2(&myslice2)
	fmt.Println(myslice2) //70 20 30 40 50 90
}
