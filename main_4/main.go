package main

import (
	"fmt"
	"strconv"
)

func atoi(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("関数atoiで失敗!: %w", err)
	}
	fmt.Println(s, i)
	return nil
}

func check(values ...string) {
	for _, s := range values {
		if err := atoi(s); err != nil {
			fmt.Println("checkエラー!:", err)
		}
	}
}

func myFunction(a int, b int) {
	defer fmt.Println(b)
	fmt.Println(a)
}

func main() {

	// fmt.Println("%w")

	// check("10", "9", "19", "sample")

	// defer fmt.Println("hello")
	// fmt.Println("world")

	// defer fmt.Println("6")
	// myFunction(1, 2)
	// defer myFunction(4, 5)
	// fmt.Println("3")

	myfunc("2", "1")
	defer myfunc("6", "5")
	myfunc("4", "3")

}

func myfunc(num1 string, num2 string) {
	defer fmt.Println(num1)
	fmt.Println(num2)
}
