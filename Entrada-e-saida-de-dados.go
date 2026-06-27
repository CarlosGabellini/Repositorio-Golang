package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// =========================================================================
	// SAÍDA DE DADOS — pacote fmt
	// =========================================================================
	//
	// Já visto antes, mas revisão rápida dos três principais:
	//
	//   fmt.Print()   → imprime sem newline
	//   fmt.Println() → imprime com newline automático
	//   fmt.Printf()  → imprime com formatação (verbos %)
	//   fmt.Sprintf() → formata e RETORNA string, não imprime
	//   fmt.Fprintf() → imprime formatado em um destino (arquivo, stderr, etc.)
	//
	// =========================================================================

	fmt.Println("=== SAÍDA DE DADOS ===")
	fmt.Print("Sem quebra de linha ")
	fmt.Print("— continua aqui\n")
	fmt.Printf("Nome: %s | Idade: %d | Altura: %.2f\n", "Carlos", 22, 1.75)

	msg := fmt.Sprintf("Valor formatado: %d", 42)
	fmt.Println(msg)

	// Saída para stderr (erros, logs — não vai para stdout)
	fmt.Fprintln(os.Stderr, "Isso vai para o stderr, não para o stdout")

	// =========================================================================
	// ENTRADA DE DADOS
	// =========================================================================
	//
	// Go tem duas formas principais de ler entrada do usuário:
	//
	//   1. fmt.Scan / fmt.Scanf / fmt.Scanln  → simples, separado por espaço
	//   2. bufio.Scanner                       → lê linha inteira, mais robusto
	//
	// A diferença crítica:
	//   fmt.Scan para no espaço — "João Silva" lido com Scan vira só "João"
	//   bufio.Scanner lê tudo até o Enter
	//
	// =========================================================================

	// -------------------------------------------------------------------------
	// FORMA 1: fmt.Scan
	// Lê um valor, para no espaço ou newline.
	// O & passa o ENDEREÇO da variável (ponteiro) — Go precisa saber onde gravar.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== fmt.Scan ===")
	fmt.Print("Digite um número inteiro: ")

	var numero int
	fmt.Scan(&numero) // & é obrigatório — sem ele o compilador rejeita

	fmt.Printf("Você digitou: %d\n", numero)

	// -------------------------------------------------------------------------
	// FORMA 2: fmt.Scanf
	// Lê com formato específico, igual ao printf do C.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== fmt.Scanf ===")
	fmt.Print("Digite seu nome e idade (ex: Carlos 22): ")

	var nome string
	var idade int
	fmt.Scanf("%s %d", &nome, &idade) // lê string até espaço, depois inteiro

	fmt.Printf("Nome: %s | Idade: %d\n", nome, idade)

	// -------------------------------------------------------------------------
	// FORMA 3: fmt.Scanln
	// Igual ao Scan, mas para no newline (Enter).
	// Ainda divide por espaço, mas não passa para a linha seguinte.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== fmt.Scanln ===")
	fmt.Print("Digite dois números separados por espaço: ")

	var a, b float64
	fmt.Scanln(&a, &b)

	fmt.Printf("Soma: %.2f\n", a+b)

	// -------------------------------------------------------------------------
	// PROBLEMA do fmt.Scan com strings com espaço
	//
	// Se você fizer:
	//   var nomeCompleto string
	//   fmt.Scan(&nomeCompleto)
	//   → e digitar "João Silva", só "João" é lido
	//
	// Para ler linha inteira com espaços, use bufio.Scanner (abaixo).
	// -------------------------------------------------------------------------

	// -------------------------------------------------------------------------
	// FORMA 4: bufio.Scanner — lê linha inteira (a mais usada na prática)
	// -------------------------------------------------------------------------

	fmt.Println("\n=== bufio.Scanner — linha inteira ===")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite seu nome completo: ")
	scanner.Scan()                      // lê uma linha
	nomeCompleto := scanner.Text()      // pega o texto lido (sem o \n)

	fmt.Printf("Olá, %s!\n", nomeCompleto)

	// Lendo múltiplas linhas com o mesmo scanner
	fmt.Print("Digite sua cidade: ")
	scanner.Scan()
	cidade := scanner.Text()

	fmt.Printf("Cidade: %s\n", cidade)

	// -------------------------------------------------------------------------
	// CONVERSÃO DE TIPO — entrada sempre chega como string
	//
	// Quando você lê com bufio, tudo é string.
	// Para usar como número, precisa converter com o pacote strconv.
	//
	//   strconv.Atoi()        → string → int  (retorna valor e erro)
	//   strconv.ParseFloat()  → string → float64
	//   strconv.ParseBool()   → string → bool
	//   strconv.Itoa()        → int → string
	//
	// -------------------------------------------------------------------------

	fmt.Println("\n=== CONVERSÃO string → número ===")

	fmt.Print("Digite sua idade: ")
	scanner.Scan()
	idadeStr := scanner.Text()

	idadeConvertida, err := strconv.Atoi(idadeStr) // Atoi = ASCII to Integer
	if err != nil {
		fmt.Println("Erro: não é um número válido:", err)
	} else {
		fmt.Printf("Idade como int: %d | Daqui a 10 anos: %d\n", idadeConvertida, idadeConvertida+10)
	}

	// ParseFloat precisa do tamanho em bits (64 para float64)
	fmt.Print("Digite um número decimal: ")
	scanner.Scan()
	floatStr := scanner.Text()

	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Erro: não é um decimal válido:", err)
	} else {
		fmt.Printf("Float64: %.4f\n", floatVal)
	}

	// int → string
	codigo := 1042
	codigoStr := strconv.Itoa(codigo)
	fmt.Printf("Código como string: %q (tipo: %T)\n", codigoStr, codigoStr)

	// -------------------------------------------------------------------------
	// TRATAMENTO DE ESPAÇOS — strings.TrimSpace
	//
	// Ao ler com bufio, se o usuário digitar espaço antes ou depois,
	// esses espaços ficam na string. TrimSpace remove isso.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== TrimSpace ===")

	fmt.Print("Digite algo com espaços nas bordas: ")
	scanner.Scan()
	entrada := scanner.Text()

	limpo := strings.TrimSpace(entrada)
	fmt.Printf("Original: %q\n", entrada)
	fmt.Printf("Limpo:    %q\n", limpo)

	// -------------------------------------------------------------------------
	// LENDO ATÉ O USUÁRIO DIGITAR "sair" — loop com entrada
	// -------------------------------------------------------------------------

	fmt.Println("\n=== LOOP DE ENTRADA (digite 'sair' para parar) ===")

	for {
		fmt.Print("→ ")
		scanner.Scan()
		linha := strings.TrimSpace(scanner.Text())

		if linha == "sair" {
			fmt.Println("Encerrando.")
			break
		}

		if linha == "" {
			fmt.Println("Linha vazia, tente de novo.")
			continue
		}

		fmt.Printf("Você digitou: %q\n", linha)
	}

	// -------------------------------------------------------------------------
	// ARGUMENTOS DE LINHA DE COMANDO — os.Args
	//
	// Quando você executa: go run main.go arg1 arg2
	// os.Args[0] = caminho do programa
	// os.Args[1] = "arg1"
	// os.Args[2] = "arg2"
	// -------------------------------------------------------------------------

	fmt.Println("\n=== os.Args ===")

	fmt.Printf("Total de argumentos: %d\n", len(os.Args))

	for i, arg := range os.Args {
		fmt.Printf("os.Args[%d] = %q\n", i, arg)
	}

	if len(os.Args) > 1 {
		fmt.Printf("Primeiro argumento passado: %s\n", os.Args[1])
	} else {
		fmt.Println("Nenhum argumento extra passado.")
		fmt.Println("Teste com: go run condicionais_loops.go hello mundo")
	}
}