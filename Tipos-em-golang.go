package main

import "fmt"

// =============================================================================
// CONSTANTES — valor fixo, definido em tempo de compilação, não pode ser alterado
// =============================================================================

const Pi = 3.14159
const AppName = "MeuPrograma"
const MaxTentativas = 5

// Bloco de constantes (mais limpo que declarar uma por uma)
const (
	StatusOK    = 200
	StatusErro  = 500
	Versao      = "1.0.0"
)

// =============================================================================
// TIPOS BÁSICOS
// =============================================================================
//
// INTEIROS
//   int      → tamanho depende da arquitetura (32 ou 64 bits). Use esse na maioria dos casos.
//   int8     → -128 a 127
//   int16    → -32.768 a 32.767
//   int32    → -2.147.483.648 a 2.147.483.647
//   int64    → -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807
//
// INTEIROS SEM SINAL (não aceitam negativos)
//   uint     → 0 a ~4 bilhões (32-bit) ou muito mais (64-bit)
//   uint8    → 0 a 255  (também chamado de byte)
//   uint16   → 0 a 65.535
//   uint32   → 0 a 4.294.967.295
//   uint64   → 0 a 18.446.744.073.709.551.615
//
// PONTO FLUTUANTE (números com casas decimais)
//   float32  → precisão ~6-7 casas decimais
//   float64  → precisão ~15-16 casas decimais. Use esse, é o padrão do Go.
//
// COMPLEXOS (raramente usados fora de matemática/ciência)
//   complex64
//   complex128
//
// TEXTO
//   string   → sequência de bytes UTF-8, imutável
//   rune     → alias para int32, representa um caractere Unicode
//   byte     → alias para uint8, representa um byte bruto
//
// BOOLEANO
//   bool     → true ou false
//
// =============================================================================

func main() {

	// -------------------------------------------------------------------------
	// VAR — declaração explícita de variável
	// O Go inicializa com "zero value" se você não atribuir nada:
	//   int   → 0
	//   float → 0.0
	//   bool  → false
	//   string → ""
	// -------------------------------------------------------------------------

	var nome string = "Carlos"
	var idade int = 22
	var altura float64 = 1.75
	var ativo bool = true

	// Declaração sem valor inicial (zero value)
	var contador int    // vale 0
	var mensagem string // vale ""
	var ligado bool     // vale false

	// -------------------------------------------------------------------------
	// := — declaração curta (inferência de tipo)
	// Só funciona DENTRO de funções. O Go deduz o tipo pelo valor.
	// -------------------------------------------------------------------------

	cidade := "São Paulo"         // string
	peso := 80.5                  // float64
	filhos := 0                   // int
	casado := false               // bool
	letra := 'A'                  // rune (int32)
	b := byte('Z')                // byte (uint8)

	// -------------------------------------------------------------------------
	// STRINGS — detalhes importantes
	// -------------------------------------------------------------------------

	saudacao := "Olá, mundo!"
	multiLinha := `Isso é uma
	string com múltiplas
	linhas (raw string literal).
	Aceita "aspas" e 'tudo' sem escape.`

	comprimento := len(saudacao) // len() retorna número de BYTES, não caracteres!

	// Concatenação
	nomeCompleto := "João" + " " + "Silva"

	// -------------------------------------------------------------------------
	// SAÍDA DE DADOS — fmt
	// -------------------------------------------------------------------------

	// fmt.Println → imprime com espaço entre args e newline no final
	fmt.Println("=== SAÍDA COM Println ===")
	fmt.Println(nome, idade, altura, ativo)

	// fmt.Print → imprime SEM newline automático
	fmt.Print("Sem quebra de linha ")
	fmt.Print("aqui também\n") // \n manual

	// fmt.Printf → formatação com verbos (o mais usado)
	fmt.Println("\n=== SAÍDA COM Printf ===")

	// Verbos principais:
	// %v  → valor no formato padrão (funciona para qualquer tipo)
	// %T  → tipo da variável
	// %d  → inteiro decimal
	// %f  → float  (%0.2f = 2 casas decimais)
	// %s  → string
	// %q  → string com aspas
	// %t  → bool
	// %c  → caractere (rune/byte)
	// %b  → binário
	// %x  → hexadecimal
	// %p  → ponteiro (endereço de memória)

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)
	fmt.Printf("Altura: %.2f m\n", altura)
	fmt.Printf("Ativo: %t\n", ativo)
	fmt.Printf("Tipo de 'peso': %T\n", peso)
	fmt.Printf("Valor padrão (%%v): %v\n", altura)
	fmt.Printf("Letra: %c (valor int: %d)\n", letra, letra)
	fmt.Printf("Byte: %c (valor uint8: %d)\n", b, b)

	// fmt.Sprintf → formata e RETORNA uma string (não imprime)
	fmt.Println("\n=== SAÍDA COM Sprintf ===")
	resultado := fmt.Sprintf("Olá, %s! Você tem %d anos.", nome, idade)
	fmt.Println(resultado)

	// -------------------------------------------------------------------------
	// EXIBINDO AS VARIÁVEIS RESTANTES (para o compilador não reclamar de unused)
	// -------------------------------------------------------------------------

	fmt.Println("\n=== DEMAIS VARIÁVEIS ===")
	fmt.Printf("cidade=%s | peso=%.1f | filhos=%d | casado=%t\n", cidade, peso, filhos, casado)
	fmt.Printf("contador=%d | mensagem=%q | ligado=%t\n", contador, mensagem, ligado)
	fmt.Printf("Comprimento de %q em bytes: %d\n", saudacao, comprimento)
	fmt.Printf("Nome completo: %s\n", nomeCompleto)
	fmt.Println(multiLinha)

	// -------------------------------------------------------------------------
	// CONSTANTES em uso
	// -------------------------------------------------------------------------
	fmt.Println("\n=== CONSTANTES ===")
	fmt.Printf("Pi=%.5f | App=%s | MaxTentativas=%d\n", Pi, AppName, MaxTentativas)
	fmt.Printf("HTTP OK=%d | HTTP Erro=%d | Versão=%s\n", StatusOK, StatusErro, Versao)

	// -------------------------------------------------------------------------
	// OBSERVAÇÃO: CONVERSÃO DE TIPOS
	// Go NÃO faz conversão implícita. Você precisa ser explícito.
	// -------------------------------------------------------------------------
	fmt.Println("\n=== CONVERSÃO DE TIPOS ===")
	var x int = 10
	var y float64 = float64(x) // int → float64 explícito
	var z int = int(y)          // float64 → int (trunca, não arredonda)
	fmt.Printf("x(int)=%d | y(float64)=%.1f | z(int de float)=%d\n", x, y, z)

	numStr := fmt.Sprintf("%d", x) // int → string via Sprintf
	fmt.Printf("numStr=%q (tipo: %T)\n", numStr, numStr)
}