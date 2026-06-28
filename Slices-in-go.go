package main

import "fmt"

//Vamos ver como funciona Slices em golang, visto que devem ser melhores para lidar com arrays;

func main() {

	//Existem 3 propriedades em golang para arrays;
	
	// Pointer	-> comeco do seu array;
	// length	-> Significa medida em ingles, entao é quantos elementos eu tenho no array;
	// capacity	-> Capacidade total do array;

	array := [5]int {1, 2, 3, 4, 5}
	sl := array[1:]

	sl = sl[:4]			//Agora adicionamos o quarto elemento do array original com slice.
	
	sl[0] = 100

	fmt.Println(sl[3], array)
	fmt.Println(sl, len(sl), cap(sl))


	//Abaixo, estamos usando Slice sem presi
	
	slice = []string{"Hello", "World", "Family"}
	fmt.
}