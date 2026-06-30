package main

import "fmt"

//Vamos aprender melhor sobre ErrorHandling

func divide(a int, b int) int {
	return a / b
}

func defferedFunc() {
	fmt.Println("defer!")
	r := recover()

	fmt.Println(r)
}

func main() {

	defer defferedFunc()		//Adia uma funcao;
	panic("This caused a crash!")		//Tudo que estiver depois de um panico nao executa!

	fmt.Println("run")
}