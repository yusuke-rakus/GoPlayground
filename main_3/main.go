package main

import "fmt"

type Person struct {
	name *string
	age  *int
}

func (p *Person) addAge(num int) *Person {
	*p.age += num
	return p
}

func main() {
	name := "Liza"
	age := 18
	Liza := &Person{name: &name, age: &age}
	fmt.Println(*Liza)

	ageNum := 10
	Liza.addAge(ageNum)
	fmt.Println(*Liza)
}
