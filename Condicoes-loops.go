package main

import "fmt"

func main() {

	// =========================================================================
	// IF / ELSE IF / ELSE
	// =========================================================================
	//
	// Regras obrigatórias:
	//   1. Sem parênteses na condição — if x > 0 {}, NÃO if (x > 0) {}
	//   2. Chaves { } são obrigatórias SEMPRE, mesmo com uma linha só
	//   3. O { de abertura fica NA MESMA LINHA do if/else (padrão Go)
	//
	// =========================================================================

	fmt.Println("=== IF / ELSE IF / ELSE ===")

	idade := 17

	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else if idade >= 16 {
		fmt.Println("Pode votar, mas não é maior de idade")
	} else {
		fmt.Println("Menor de idade, sem direitos políticos")
	}

	// -------------------------------------------------------------------------
	// IF COM INICIALIZAÇÃO — exclusivo do Go
	// Você pode declarar uma variável dentro do próprio if.
	// Essa variável só existe dentro do bloco if/else. Fora dele, não existe.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== IF COM INICIALIZAÇÃO ===")

	if nota := 7.5; nota >= 7.0 {
		fmt.Printf("Aprovado com nota %.1f\n", nota)
	} else if nota >= 5.0 {
		fmt.Printf("Recuperação com nota %.1f\n", nota)
	} else {
		fmt.Printf("Reprovado com nota %.1f\n", nota)
	}
	// fmt.Println(nota) → ERRO: nota não existe aqui fora

	// -------------------------------------------------------------------------
	// OPERADORES DE COMPARAÇÃO E LÓGICOS
	// -------------------------------------------------------------------------
	//
	//   ==   igual
	//   !=   diferente
	//   >    maior que
	//   <    menor que
	//   >=   maior ou igual
	//   <=   menor ou igual
	//
	//   &&   E lógico (AND)
	//   ||   OU lógico (OR)
	//   !    NÃO lógico (NOT)
	//
	// -------------------------------------------------------------------------

	fmt.Println("\n=== OPERADORES LÓGICOS ===")

	temperatura := 28
	chovendo := false

	if temperatura > 25 && !chovendo {
		fmt.Println("Bom dia para sair")
	}

	saldo := 0.0
	limite := 500.0

	if saldo > 0 || limite > 0 {
		fmt.Println("Tem fundos disponíveis")
	}

	// =========================================================================
	// FOR — O ÚNICO LOOP DO GO
	// =========================================================================
	//
	// Go não tem while, do-while, foreach.
	// Tudo isso é feito com for, em formatos diferentes.
	//
	// =========================================================================

	// -------------------------------------------------------------------------
	// FORMATO 1: for clássico (igual ao C/Java)
	// for inicialização; condição; pós-execução { }
	// -------------------------------------------------------------------------

	fmt.Println("\n=== FOR CLÁSSICO ===")

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// Contando de trás pra frente
	for i := 5; i > 0; i-- {
		fmt.Printf("contagem regressiva: %d\n", i)
	}

	// -------------------------------------------------------------------------
	// FORMATO 2: for como WHILE
	// Omite inicialização e pós-execução, fica só com a condição.
	// Comportamento idêntico ao while de outras linguagens.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== FOR COMO WHILE ===")

	tentativas := 0

	for tentativas < 3 {
		fmt.Printf("Tentativa %d\n", tentativas+1)
		tentativas++
	}

	// -------------------------------------------------------------------------
	// FORMATO 3: for infinito
	// Sem nenhuma condição. Roda para sempre até um break ou return.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== FOR INFINITO (com break) ===")

	contador := 0

	for {
		if contador >= 4 {
			break // sai do loop imediatamente
		}
		fmt.Printf("contador = %d\n", contador)
		contador++
	}

	// -------------------------------------------------------------------------
	// BREAK e CONTINUE
	//
	// break    → sai do loop imediatamente
	// continue → pula para a próxima iteração, ignora o restante do bloco
	// -------------------------------------------------------------------------

	fmt.Println("\n=== CONTINUE (pula pares) ===")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // pula números pares
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n=== BREAK (para no 6) ===")

	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// FORMATO 4: for range
	// Usado para iterar sobre arrays, slices, strings, maps e channels.
	// Retorna índice e valor a cada iteração.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== FOR RANGE (slice) ===")

	frutas := []string{"maçã", "banana", "laranja", "uva"}

	for indice, fruta := range frutas {
		fmt.Printf("[%d] = %s\n", indice, fruta)
	}

	// Se não precisar do índice, use _ para descartar (Go obriga uso de variáveis)
	fmt.Println("\n=== FOR RANGE sem índice ===")

	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// Se não precisar do valor, só o índice:
	fmt.Println("\n=== FOR RANGE só índice ===")

	for i := range frutas {
		fmt.Printf("posição %d\n", i)
	}

	// FOR RANGE em string — itera por RUNE (caractere Unicode), não por byte
	fmt.Println("\n=== FOR RANGE em string ===")

	palavra := "Go!"

	for i, caractere := range palavra {
		fmt.Printf("posição %d → %c (rune: %d)\n", i, caractere, caractere)
	}

	// -------------------------------------------------------------------------
	// LOOPS ANINHADOS
	// -------------------------------------------------------------------------

	fmt.Println("\n=== LOOPS ANINHADOS ===")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d×%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// LABELS — break/continue em loops aninhados
	// Pouco usados, mas existem. Permitem sair de um loop externo a partir
	// de dentro de um loop interno.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== BREAK COM LABEL ===")

externo:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i+j == 4 {
				fmt.Printf("Parou em i=%d, j=%d\n", i, j)
				break externo // sai do loop EXTERNO, não só do interno
			}
			fmt.Printf("i=%d j=%d | ", i, j)
		}
	}
	fmt.Println()

	// =========================================================================
	// SWITCH — alternativa ao if-else em cadeia
	// =========================================================================
	//
	// Em Go, o switch:
	//   - Não precisa de break (cada case já para automaticamente)
	//   - Aceita qualquer tipo comparável, não só inteiros
	//   - Pode ter múltiplos valores por case
	//   - Pode rodar sem expressão (funciona como if-else)
	//
	// =========================================================================

	fmt.Println("\n=== SWITCH ===")

	diaSemana := 3

	switch diaSemana {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	case 6, 7: // múltiplos valores no mesmo case
		fmt.Println("Final de semana")
	default:
		fmt.Println("Dia inválido")
	}

	// Switch sem expressão — equivalente a if-else encadeado
	fmt.Println("\n=== SWITCH SEM EXPRESSÃO ===")

	pontos := 85

	switch {
	case pontos >= 90:
		fmt.Println("Conceito A")
	case pontos >= 70:
		fmt.Println("Conceito B")
	case pontos >= 50:
		fmt.Println("Conceito C")
	default:
		fmt.Println("Conceito D — reprovado")
	}

	// fallthrough — força a execução do próximo case (comportamento inverso ao padrão)
	fmt.Println("\n=== FALLTHROUGH ===")

	x := 1

	switch x {
	case 1:
		fmt.Println("case 1")
		fallthrough // força execução do case 2 também
	case 2:
		fmt.Println("case 2 (executou por fallthrough)")
	case 3:
		fmt.Println("case 3")
	}
}