package main

import "fmt"

func main() {

	sl := make([]int, 10)			//Capacidade é de 10 inteiros;
	sl[0] = 124

	fmt.Println(sl[0], len(sl))
}