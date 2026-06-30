package main

import "fmt"

//A interface trabalha junto com structs para definir uma serie
// de funcoes e parametros;

type Shape interface {
	getPerimeter() uint
	getSides() uint
}

type Square struct {
	largura uint
}

type Triangle struct {
	a uint
	b uint
	c uint
}

func (s Square) getPerimeter() uint {
	return 4 * s.largura
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

func (t Triangle) getSides() []uint {
	return []uint{t.a, t.b, t.c}
}

func main() {

	var s Shape = Triangle{1, 2, 3}
	var a Shape = Square{15}

	fmt.Println(s.getPerimeter())
	fmt.Println(a.getPerimeter())
}