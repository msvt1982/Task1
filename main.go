package main

import (
	"fmt"
	"time"
)

var ch chan int

func wiatInerraptByUser(ch chan int) {
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
	go wiatInerraptByUser(ch)
	time.Sleep(10 * time.Second)
	readCh := <-ch
	if readCh == 1 {
		println("Goodbye world")
	}
}
