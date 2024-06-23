package main

import (
	"fmt"
	"strconv"
)

func parser(str_list <-chan string) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error)

	go func() {
		for str := range str_list {
			fmt.Println("str: ", str)
			i, err := strconv.Atoi(str)
			if err != nil {
				errCh <- err
			}
			ch <- i
		}
	}()
	return ch, errCh
}

func case_1() {
	ch := make(chan string)
	resultCh, errCh := parser(ch)

	str_list := []string{"111", "2", "3", "4", "hej"}

	for _, s := range str_list {
		ch <- s
		select {
		case result := <-resultCh:
			fmt.Printf("%d(%T)", result, result)
		case err := <-errCh:
			fmt.Println("error!: ", err)
			return
		}
	}
	close(ch)
}

func main() {
	case_1()

}
