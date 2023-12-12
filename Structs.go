package main

type person struct {
	name string
	age  int
}

func newperson(name string) *person {
	p := person{name: name}
	p.age = 55
	return &p
}

type rectangle struct {
	height, width float64

	calc struct {
		area, perimeter float64
	}
}
