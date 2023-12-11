package main

import "fmt"

func array() {
	var a [4]int
	fmt.Println(a) // [0 0 0 10]

	//taking input
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < 4; i++ {
		fmt.Print(a[i], " ")
		print(" ") // add another space
	} /// output of the array without bracket-->12 34 56 100

	println()
	a[3] = 10
	fmt.Println(a) // [12 34 56 10]

	//manipulations
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println((sum / len(a)))     //ok sum,len(a) type are SAME
	fmt.Println(float64(sum) / 5.0) // type casting

	//another type of for loop
	arr := [5]int{10, 30, 50, 100, 112}

	total := 0
	for _, v := range arr {
		total += v
	}

	fmt.Println(total)
}
