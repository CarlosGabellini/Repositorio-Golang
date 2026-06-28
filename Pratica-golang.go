package main

import "fmt"

func main() {

	const controle int32 = 100
	var variable1 int32 = 21
	var true1 bool = true

	variable1 = variable1 * controle

	if true1 == false {
		fmt.Printf("Deu verdadeiro, valor da variavel: %d\n", variable1)
		fmt.Printf("Valor de controle: %d\n", controle)
	}
}