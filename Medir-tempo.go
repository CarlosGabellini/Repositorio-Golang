package main

import (
	"fmt"
	"time"
)

// Esta função calcula a diferença de tempo. 
// Ela recebe o tempo de início e imprime a duração total.
func medirTempo(inicio time.Time, nomeDaFuncao string) {
	duracao := time.Since(inicio)
	fmt.Printf("⏱️ [Métrica] A função '%s' demorou %v para ser concluída.\n", nomeDaFuncao, duracao)
}

func operacaoDemorada() {
	// O truque está aqui: 
	// time.Now() é executado IMEDIATAMENTE quando a linha é lida (no início da função).
	// Mas a função 'medirTempo' só será EXECUTADA no final de 'operacaoDemorada'.
	defer medirTempo(time.Now(), "operacaoDemorada")

	fmt.Println("Iniciando uma tarefa pesada...")
	
	// Simula um processo que demora 2 segundos (ex: download ou busca no banco)
	time.Sleep(12 * time.Second) 
	
	fmt.Println("Tarefa pesada concluída!")
}

func main() {
	fmt.Println("--- Início do Programa ---")
	
	operacaoDemorada()
	
	fmt.Println("--- Fim do Programa ---")
}