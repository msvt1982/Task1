// package main

// import (
// 	"fmt"
// 	"time"
// )

// func wiatInerruptByUser(ch chan int, now time.Time) {
// 	var name string
// 	fmt.Scanf("%s", &name)
// 	if name == "qwe" {
// 		now1 := time.Now()
// 		now2 := int(now1.Sub(now).Seconds())
// 		ch <- now2
// 	} else {
// 		ch <- 0
// 	}
// }

// func main() {
// 	ch := make(chan int) //как глобальную?..
// 	println("Hello World")
// 	now := time.Now()
// 	go wiatInerruptByUser(ch, now)
// 	time.Sleep(10 * time.Second)
// 	readCh := <-ch
// 	if readCh != 0 {
// 		fmt.Printf("Stopped by the user after %v seconds\n", readCh)
// 		// fmt.Printf("Stopped by the user after seconds")
// 	}
// 	println("Goodbye world")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {

// 	sigs := make(chan os.Signal, 1)
// 	progTimer := make(chan int)
// 	startProg := time.Now()
// 	var sig os.Signal

// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	fmt.Println("Hello World")

// 	go func() {
// 		sig := <-sigs
// 		// fmt.Println()
// 		fmt.Println(sig)
// 		closeProg := time.Now()
// 		closeProgAfter := int(closeProg.Sub(startProg).Seconds())
// 		fmt.Printf("\nStopped by the user after %v seconds\n", closeProgAfter)
// 		progTimer <- closeProgAfter
// 	}()

//		time.Sleep(3 * time.Second)
//		readCh := <-progTimer
//		fmt.Println("1")
//		if sig == syscall.SIGINT {
//			fmt.Println("2")
//			fmt.Println("3")
//			if readCh > 0 {
//				fmt.Println("4")
//				fmt.Printf("Stopped by the user after %v seconds\n", readCh)
//			}
//		}
//		fmt.Println("Goodbye world")
//	}
//
// ----------------------
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
		// time.Sleep(3 * time.Second)
		fmt.Println("Goodbye world")
	}
}
