package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"
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

	//concurrency, goroutines
	go f(0)

	var ch string //without it program will auto terminate, reach at the end in individual
	fmt.Scanln(&ch)

	//channel
	var v chan string = make(chan string)

	go pinger(v)
	go ponger(v)
	go printer(v)

	var input string
	fmt.Scanln(&input)

	//select on channel
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "First stage"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			c2 <- "second stage"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var tt string
	fmt.Scanln(&tt)

	///container
	//push
	var con list.List

	con.PushBack(10)
	con.PushBack(20)
	con.PushBack(30)
	//printing
	for i := con.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(int))
	}

	//Mutex
	m := new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "Start")
			time.Sleep(time.Second * 2)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var ii string
	fmt.Scanln(&ii)

	//HTTP Client
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status : ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println((scanner.Text()))
		//time.Sleep(time.Second * 2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//HTTP Server
	http.HandleFunc("/another way", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":9999", nil)

}
