package main

import (
	"fmt"
	"time"
)

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
	var name string
	fmt.Scanf("%s", &name)
	if name == "qwe" {
		fmt.Print(name)
	}

	ch := make(chan int)
	println("Hello World")
	go wiatInerraptByUser(ch)
	time.Sleep(10 * time.Second)
	readCh := <-ch
	if readCh == 1 {
		println("Goodbye world")
	}
}
