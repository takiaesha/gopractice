package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	//boolean
	fmt.Println(true || true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false || true)

	var p int = 100
	fmt.Println(p)

	// Numbers
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)
	fmt.Println("200 - 100 = ", 200-100)
	fmt.Println("200 * 100  = ", 200*100)
	fmt.Println("22.5 / 3 = ", 22.50/3)

	//Strings
	fmt.Print("Coffee:")
	fmt.Print(`Hi there`)
	fmt.Println("Hi There")

	fmt.Println(len("Hello World"))
	var d = len("TakiaEsha")
	fmt.Println("size of second string = ", d)
	fmt.Println("Hello" + "World")

	fmt.Println("Hello World"[1]) // e = 101
	fmt.Println("Hello World"[2]) //l  = 108

	/// variables
	var q int = 12345
	var i = 1234000005 //same meaning
	o := 1234567       //best way to use

	fmt.Println(q)
	fmt.Println(i)
	fmt.Println(o)

	ch := "AppsCode"
	l := len(ch)

	fmt.Println("Name = ", ch)
	fmt.Println("Size  = ", l)
	fmt.Println(ch[4]) /// ASCII value C = 67 print
	//equal operators
	s1 := "AAAAAAA"
	s2 := "AAa"
	fmt.Println(s1 == s2)

	//scanf can be used
	/*var number float64
	fmt.Scanf("%f", &number)
	output := number
	fmt.Println("Using Scanf, output wil be = ", output)

	//increment opertor
	k := 10
	k += 1
	fmt.Println(k)*/

	main1()

	///file handling example
	a := 12
	b := 10

	sum := add(a, b)
	dif := difference(a, b)

	fmt.Println("Sum is  : ", sum)
	fmt.Println("Difference is : ", dif)

	//control structure
	//controlS()
	//p3()

	//arrays
	//array()

	slice()

}
