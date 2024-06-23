package main

import "fmt"

type Person struct {
	name   string
	tall   int
	weight int
}

func (s *Person) change_name(new_name string) {
	s.name = new_name
}

func addOnea(s *Person) {
	s.tall++
	s.weight++
}

func main_1() {
	p1 := Person{name: "matsumoto", tall: 10, weight: 20}

	p2 := p1 // 値渡し
	p2.name = "yamada"
	fmt.Println(p1) // matsumoto
	fmt.Println(p2) // yamada

	p3 := &p1 // 参照渡し
	(*p3).name = "tanaka"
	fmt.Println(p1)  // tanaka
	fmt.Println(*p3) // tanaka
}

func main_2() {
	name := "jack"

	name_2 := &name // ポインタ変数を渡す

	fmt.Println(name_2)  // 0xc000014070
	fmt.Println(*name_2) // jack
}

func main_3() {
	name := "太郎"
	namePoint := &name

	*namePoint = "二郎"
	/*
		namePoint = "二郎" // 'namePoint'はポインタ変数(アドレス)のため、意味わからないことしてエラー

		namePoint := name
		namePoint = "二郎" // 値渡しからの値変更ならOK
	*/

	fmt.Println(name)
	fmt.Println(*namePoint)
}

func main() {
	main_1()
	// main_2()
	// main_3()
}
