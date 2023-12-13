package main

import "fmt"

type BookStore interface {
	bookPrice() float64
	discountPrice(float64) float64
}

type calc1 struct {
	tag, dis float64
}

// each book price is calculated with  $6, eachId * 6
func (t1 calc1) bookPrice() float64 {
	eachId := t1.tag * 6.0
	return eachId
}

// / discounted price get 10%, passing parameter
func (t1 calc1) discountPrice(t2 float64) float64 {
	ans2 := t1.tag * 0.1
	return ans2 * t2
}

func recitt(p BookStore) {
	fmt.Println(p)
	vv := p.bookPrice()

	fmt.Println(p.bookPrice())
	fmt.Println(p.discountPrice(vv))

}
