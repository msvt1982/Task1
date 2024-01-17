package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var sig os.Signal
var startProg, closeProg time.Time
var closeProgAfter int

func handle(c chan os.Signal) {
	for {
		sig = <-c
		closeProg = time.Now()
		closeProgAfter = int(closeProg.Sub(startProg).Seconds())
		_ = closeProgAfter
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGINT)
	startProg = time.Now()
	fmt.Println("Hello World")
	go handle(c)

	time.Sleep(5 * time.Second)

	if sig == syscall.SIGINT {

		fmt.Printf("\nStopped by the user after %d seconds\n", closeProgAfter)
	} else {
		fmt.Println("Goodbye world")
	}
}
