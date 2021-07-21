package main

import "fmt"

type Figure struct {
	a int
	b int
}

func p(rec Figure)(p int){
	p = 2 * (rec.a + rec.b)
	return p
}

func s(rec Figure)(s int){
	s = rec.b * rec.a
	return s

}

func main() {
	var rectangle Figure = Figure{
		a: 2,
		b: 4,
	}
	fmt.Println(rectangle)
	fmt.Println(s(rectangle))
	fmt.Println(p(rectangle))
	}
