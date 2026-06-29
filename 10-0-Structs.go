package main

import "fmt"

//Vamos tentar aprender um pouco de structs em golang;

type Person struct {
	Name string
	Age uint8		//Prescisa ser em letra maiuscula para outros pacotes;
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) SetName(newName string) {
	p.name = newName
	fmt.Println(p)
}

func main() {
	var p1 Person = Person{age: 24, name: "Tim"}
	p1.setName("Joey")

	fmt.Println(p1)
}