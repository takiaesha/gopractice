package main

import "fmt"

func main1() {
	ss := "AppsCode"
	ss += " Inc."

	fmt.Println("Another main function calling", ss)

	xx := 100.78
	xx += 100

	fmt.Println("another = ", xx)

}
