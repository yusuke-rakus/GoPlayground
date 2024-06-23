package main

import (
	"fmt"
	"time"
)

// パターン1
func send_strCh(strCh chan string, s string) {
	strCh <- s
}
func case_1() {
	strCh := make(chan string)

	go send_strCh(strCh, "a")
	go send_strCh(strCh, "b")
	go send_strCh(strCh, "c")
	fmt.Println(<-strCh, <-strCh, <-strCh)
}

// パターン2
func send_intCh(intCh chan int, i int) {
	intCh <- i
	intCh <- i * i
	intCh <- i * i * i
	close(intCh)
}
func case_2() {
	intCh := make(chan int, 3)
	send_intCh(intCh, 2)
	for i := range intCh {
		fmt.Println(i)
	}
}

func gen(intChan chan int, strChan chan string) {
	intChan <- 1
	strChan <- "abc"
	intChan <- 2
	time.Sleep(100 * time.Microsecond)
	strChan <- "ABC"
	intChan <- 3
	strChan <- "done"
}
func case_3() {
	intCh := make(chan int, 3)
	strCh := make(chan string, 3)

	go gen(intCh, strCh)
	for {
		select {
		case i := <-intCh:
			fmt.Print(i)
		case s := <-strCh:
			fmt.Print(s)
			if s == "done" {
				return
			}
		default:
			fmt.Print(".")
		}
	}
}

func main() {
	// case_1()

	// case_2()

	case_3()
}
