package main

import (
	"fmt"
	"math"
)

func main() {

	//Vamos aprender o básico de operadores em golang;
	// *, +, -, /, --, ++, %, até agora somente os básicos que eu já sabia;

	x := uint8(7)
	y := 1000
	z := int(x) / y		//Tome cuidado fazendo isso, a gente perde informacao
						//com bits importantes;

	x++;
						
	fmt.Println(z)
	fmt.Println(x)

	//Podemos importar o pacote de matematica;
	//Agora vamos poder usar potenciacao e rquad;

	var rquad64 float64 = math.Sqrt(64)
	var pot4 float64 = math.Pow(rquad64, 4)

	fmt.Println("Numero que saiu: ", rquad64)
	fmt.Println("Numero da potencia: ", pot4)
}