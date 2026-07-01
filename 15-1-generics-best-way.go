package main

import "fmt"

type Number interface {
	int | float64 | uint |
	float32
}

func add[T Number] (x T, y T) T {
	var sum T = x + y
	return sum
}

func main() {
	value := add(32.453, -24)
	value1 := add(43.54, -215)

	fmt.Println(value)
	fmt.Println(value1)
}