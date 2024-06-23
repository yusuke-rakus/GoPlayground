package main

import (
	"fmt"
	"time"
)

func do(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go do("a")
	go do("b")
	defer do("c")
	do("d")
}
