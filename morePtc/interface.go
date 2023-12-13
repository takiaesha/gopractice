package main

import (
	"fmt"
	"math"
)

// area determination
type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	height, width float64
}

func (va circle) area() float64 {
	ans1 := math.Pi * va.radius * va.radius
	return ans1
}
func (rec rectangle) area() float64 {
	ans2 := rec.height * rec.width
	return ans2
}

func calc(q shape) {
	fmt.Println(q)
	fmt.Println(q.area())
}

func main() {

	r := rectangle{height: 3, width: 4}
	c := circle{radius: 5}

	fmt.Print("Circle area : ")
	calc(c)
	fmt.Print("Rectangle area : ")
	calc(r)

	//bookstore example
	id1 := calc1{tag: 134, dis: 15}
	id2 := calc1{tag: 345, dis: 34}

	recitt(id1)
	fmt.Println("Each Book Actual Price: ", id1)

	recitt(id2)
	fmt.Println("Second ", id2)

}
