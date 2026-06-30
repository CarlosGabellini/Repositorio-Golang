package main
import "fmt"

type inteiro64 = int64
type flutuante64 = float64

func main() {

	var number1 flutuante64 = -21534
	var number2 flutuante64 = 437
	
	fmt.Printf("%.3f", number2 / number1)
}