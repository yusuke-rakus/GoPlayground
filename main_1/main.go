package main

import "fmt"

func addOne_p(x *int) {
	*x += 1
}

func pointer_do() {
	x := 10
	addOne_p(&x)
	fmt.Println("pointer_do: ", x)
}

func addOne(x int) {
	x += 1
}

func do() {
	x := 10
	addOne_p(&x)
	fmt.Println("do: ", x)
}

func main() {
	pointer_do()
	do()
}
