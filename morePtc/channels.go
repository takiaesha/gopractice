package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("Message is : ", msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

//usage of select
