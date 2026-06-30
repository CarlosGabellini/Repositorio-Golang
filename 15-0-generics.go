package main
import "fmt"

//Vamos aprender sobre generics, na qual permitem um codigo mais flexivel e engajado;

//func add(x int, y int) int {
// return x + y}
// 
// O problema de fazer isso daqui é que nos limitamos somente ao int;

//Colocando o tipo T, a funcao pode se tornar bem mais dinamica;

func add[T int | float64 | uint] (x T, y T) T {
	var sum T = x + y
	return sum
}

func main() {

	value := add(12.54, 56)
	value1 := add(2, 3)
	value2 := add(34.65, 32.12)
	
	fmt.Println(value)
	fmt.Println(value1)
	fmt.Println(value2)
}