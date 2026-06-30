package main

import (
	"errors"
	"fmt"
)

/*
================================================================================
GUIA RÁPIDO: ERROR HANDLING, DEFER E PANIC EM GO
================================================================================

Em Go, o tratamento de erros é explícito. Não usamos blocos try/catch tradicionais.
Em vez disso, funções que podem falhar retornam o erro como o último valor.

1. ERROR: Usado para situações esperadas (ex: arquivo não encontrado, falha na API).
2. DEFER: Adia a execução de uma função até que a função atual termine.
3. PANIC: Usado para erros catastróficos que o programa não consegue recuperar sozinho.
================================================================================
*/

func main() {
	// --- 1. EXEMPLO DE TRATAMENTO DE ERRO PADRÃO ---
	fmt.Println("--- 1. Tratamento de Erro Comum ---")
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro encontrado:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}

	// Forçando um erro comum
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Tratamento correto do erro:", err)
	}

	fmt.Println("--- 2. Exemplo de Defer ---")
	exemploDefer()

	fmt.Println("--- 3. Exemplo de Panic com Recover ---")
	exemploPanic()
	
	fmt.Println("O programa continuou rodando com sucesso graças ao recover!")
}

// --- FUNÇÃO COM ERRO TRADICIONAL ---
// Retorna o resultado e um "error". Se não houver erro, retorna "nil".
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Criando e retornando um erro simples
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil
}

// --- EXPLICAÇÃO DO DEFER ---
// O 'defer' empilha as funções e as executa na ordem LIFO (Last In, First Out)
// no momento em que a função atual (exemploDefer) está finalizando.
func exemploDefer() {
	// Muito usado para fechar arquivos, conexões de banco de dados ou liberar travas.
	defer fmt.Println("-> [DEFER 1] Eu fui adiado primeiro, então serei executado por último.")
	defer fmt.Println("-> [DEFER 2] Eu fui adiado por último, então serei executado primeiro.")

	fmt.Println("Executando o corpo da função exemploDefer...")
	fmt.Println("A função exemploDefer está quase terminando...")
}

// --- EXPLICAÇÃO DO PANIC E RECOVER ---
// 'panic' interrompe o fluxo normal. 'recover' captura o panic e impede o programa de quebrar.
func exemploPanic() {
	// O recover DEVE sempre estar dentro de uma função com 'defer'
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("-> [RECOVER] Recuperamos o programa de um desastre!")
			fmt.Println("-> Motivo do pânico:", r)
		}
	}()

	fmt.Println("Iniciando uma operação perigosa...")
	
	// Simulando um erro catastrófico (ex: índice fora do array ou falta de memória)
	causarPanico := true
	if causarPanico {
		panic("ERRO CRÍTICO: Banco de dados explodiu!")
	}

	// Esta linha NUNCA será executada porque o panic acontece antes
	fmt.Println("Esta linha nunca vai aparecer.")
}