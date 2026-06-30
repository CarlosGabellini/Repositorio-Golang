package main

import "fmt"

// ==============================================================================
// O QUE SÃO PONTEIROS?
// ==============================================================================
// Imagine que uma variável (ex: 'idade') é uma CAIXA que guarda um valor (ex: '30').
// O ponteiro não guarda o valor '30', ele guarda o ENDEREÇO de onde essa caixa 
// está guardada na memória do seu computador (ex: Rua da Memória, número 0x14000).
//
// Operadores mágicos:
// & (E comercial) -> Pega o ENDEREÇO de uma variável. Lê-se "endereço de".
// * (Asterisco)   -> Acessa ou altera o VALOR que está no endereço. Lê-se "valor apontado por".
// ==============================================================================

func main() {
	fmt.Println("=== Entendendo Ponteiros em Go ===")

	// ---------------------------------------------------------
	// Exemplo 1: O Básico (Brincando com a caixa e o endereço)
	// ---------------------------------------------------------
	fmt.Println("\n--- Exemplo 1: O Básico ---")
	
	idade := 30
	// Criamos um ponteiro que aponta para um int. 
	// Atribuímos a ele o endereço da variável 'idade' usando o &
	var ponteiroParaIdade *int = &idade 

	fmt.Printf("A variável 'idade' guarda o valor: %d\n", idade)
	fmt.Printf("O endereço da 'idade' na memória é: %p\n", &idade)
	fmt.Printf("O ponteiro guarda exatamente esse endereço: %p\n", ponteiroParaIdade)
	
	// Para ver o valor que está no endereço, usamos o *
	fmt.Printf("O valor que está dentro do endereço apontado é: %d\n", *ponteiroParaIdade)

	// Podemos MUDAR a variável original usando o ponteiro!
	*ponteiroParaIdade = 31 // "Vá até o endereço e mude o valor lá dentro para 31"
	
	fmt.Println("\nMudamos o valor usando '*ponteiroParaIdade = 31'")
	fmt.Printf("Novo valor da variável 'idade': %d (Ela mudou!)\n", idade)


	// ---------------------------------------------------------
	// Exemplo 2: Por que usamos ponteiros? (Em Funções)
	// ---------------------------------------------------------
	// Em Go, quando você passa uma variável para uma função, o Go faz uma CÓPIA dela.
	// Se a função altera a variável, ela altera a cópia, não a original.
	fmt.Println("\n--- Exemplo 2: Funções ---")
	
	saldo := 100

	fmt.Println("Saldo ANTES das funções:", saldo)

	// Tentando mudar sem ponteiro (passagem por valor/cópia)
	tentaMudarSaldo(saldo)
	fmt.Println("Saldo DEPOIS de tentaMudarSaldo (sem ponteiro):", saldo) // Continua 100!

	// Mudando com ponteiro (passagem por referência/endereço)
	// Passamos o ENDEREÇO usando &
	mudaSaldoDeVerdade(&saldo) 
	fmt.Println("Saldo DEPOIS de mudaSaldoDeVerdade (com ponteiro):", saldo) // Agora é 500!


	// ---------------------------------------------------------
	// Exemplo 3: Ponteiros com Structs (Objetos)
	// ---------------------------------------------------------
	// É muito comum usar ponteiros em Structs para não ter que ficar
	// copiando dados pesados toda vez, e para poder alterar os dados originais.
	fmt.Println("\n--- Exemplo 3: Structs ---")
	
	// Criamos um usuário
	usuario := Usuario{Nome: "Carlos", Idade: 25}
	fmt.Println("Usuário antes do aniversário:", usuario)

	// Chamamos o método que tem um "receiver" de ponteiro
	usuario.FazerAniversario()
	fmt.Println("Usuário depois do aniversário:", usuario)
}

// ==============================================================================
// FUNÇÕES AUXILIARES PARA OS EXEMPLOS
// ==============================================================================

// Recebe um 'int'. O Go cria uma cópia da variável aqui dentro.
func tentaMudarSaldo(valor int) {
	valor = 500 // Altera apenas a cópia. A original fica intacta.
}

// Recebe um '*int' (Ponteiro para int). 
func mudaSaldoDeVerdade(valor *int) {
	*valor = 500 // Vai até o endereço de memória original e altera para 500.
}

// Exemplo de Struct
type Usuario struct {
	Nome  string
	Idade int
}

// Método do struct Usuario. 
// O receiver '(u *Usuario)' tem um ponteiro! 
// Isso significa que vamos alterar o struct original, e não uma cópia dele.
func (u *Usuario) FazerAniversario() {
	// Dica do Go: Quando usamos structs com ponteiros, não precisamos fazer "(*u).Idade".
	// O Go é inteligente e já entende que queremos alterar a propriedade no endereço apontado.
	u.Idade += 1 
}