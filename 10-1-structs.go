package main

import "fmt"

//Fazendo uma continuacao de structs para colocar;
//Podemos combinar e herdar diferentes tipos de struct;

type Sport struct {
	name string
	position string
}

type Person struct {
	name string
	age uint8
	sports []Sport
	favSport Sport
}

func main() {
	p1 := Person{age: 21, name: "Tim", favSport: Sport{"Soccer", "D"}}
	fmt.Println(p1)
	fmt.Println(p1.favSport.position)
}