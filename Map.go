package main

import "fmt"

func main() {

	//Mapas basicamente seriam o que são os dicionários em python;
	//Eles podem guardar informacoes importantes sobre qualquer coisa do usuario;

	mp := map[string]int{"a": 1}
	ma1p := make(map[string]int)
	ma2p := map[string][]int{"a": {1, 2, 3}}	//Ha 3 formas de associar;

	delete(mp, "a")			//Deletamos o map com isso;

	//Fazendo um programa abaixo que calcula quantos numeros sao divisiveis por 1, 2, 3, 4, e 5.
	
	ma3p := map[uint]uint{}
	n := uint(1000)

	for number := uint(1); number <= n; number++ {
		for d := uint(1); d <= 7; d++ {
			if number % d == 0 {
				ma3p[d]++
			}
		}
	}

	
	fmt.Println(mp)
	fmt.Println(ma1p)
	fmt.Println(ma2p)
	fmt.Println(ma3p)
}