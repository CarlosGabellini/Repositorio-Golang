package main
import "fmt"

func main() {

	//Eu não posso mudar o tamanho dos arrays uma vez que eles foram postos!
	var array[10][10] float64

	fmt.Printf("Type array: %T\n\n", array)

	for i := 0; i < 10; i++ {
		for a := 0; a < 10; a++ {
			array[i][a] = (float64(a) + 1) * (float64(i) + 1)
			fmt.Printf("[%.2f]\t", array[i][a])
		}
		
		fmt.Printf("\n")
	}

	fmt.Printf("O valor que deu eh: %d\n", dobrar(2, 3))
}