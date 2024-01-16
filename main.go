package main

import (
	"fmt"
	"time"
)

var ch chan int

func wiatInerruptByUser(ch chan int) {
	var name string
	fmt.Scanf("%s", &name)
	if name == "qwe" {
		ch <- 1
	} else {
		ch <- 0
	}
}

func main() {
	println("Hello World")
	go wiatInerruptByUser(ch)
	time.Sleep(10 * time.Second)
	readCh := <-ch
	if readCh == 1 {
		println("Goodbye world")
	}
}
