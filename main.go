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
	now := time.Now()
	go wiatInerruptByUser(ch)
	time.Sleep(10 * time.Second)
	readCh := <-ch
	if readCh == 1 {
		now1 := time.Now().Sub(now).Seconds()
		fmt.Println("Stopped by the user after %v seconds", now1)
	}
	println("Goodbye world")
}
