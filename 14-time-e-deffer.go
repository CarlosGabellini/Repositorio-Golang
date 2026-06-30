package main

import (
	"fmt"
	"time"
)

// 1. Crie essa função auxiliar uma única vez no seu projeto
// Vou utilizar esta funcao para ver quanto tempo uma funcao está demorando;
func Rastrear(nome string) func() {
	inicio := time.Now() // Captura o tempo no momento em que entra na função
	
	// Retorna uma função que será executada pelo defer no final
	return func() {
		fmt.Printf("⏱️ [%s] finalizou em %v\n", nome, time.Since(inicio))
	}
}

// 2. Agora, veja como fica limpo aplicar nas suas funções:

func processarPagamento() {
	// Repare nos parênteses no final: defer Rastrear(...)()
	// Nao esquecer dos parenteses
	defer Rastrear("processarPagamento")() 

	fmt.Println("Conectando com a operadora de cartão...")
	time.Sleep(800 * time.Millisecond)
}

func enviarEmailDeConfirmacao() {
	defer Rastrear("enviarEmailDeConfirmacao")()

	fmt.Println("Enviando e-mail...")
	time.Sleep(150 * time.Millisecond)
}

func main() {
	processarPagamento()
	enviarEmailDeConfirmacao()
}