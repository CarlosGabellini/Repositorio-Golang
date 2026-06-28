package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// =========================================================================
	// ARRAY — tamanho FIXO, definido em tempo de compilação
	// =========================================================================
	//
	// Sintaxe: var nome [tamanho]tipo
	//
	// Regras críticas:
	//   - O tamanho FAZ PARTE do tipo. [3]int e [5]int são tipos DIFERENTES.
	//   - Tamanho não pode mudar depois de declarado. Nunca.
	//   - Arrays são copiados por valor — atribuir um array a outro copia tudo.
	//   - Zero value: cada elemento recebe o zero value do tipo (0, "", false)
	//
	// Na prática: arrays puros são pouco usados em Go.
	// Quase sempre você vai usar SLICE (próxima seção).
	//
	// =========================================================================

	fmt.Println("=== ARRAY ===")

	// Declaração com tamanho fixo
	var notas [5]float64
	notas[0] = 8.5
	notas[1] = 7.0
	notas[2] = 9.2
	notas[3] = 6.8
	notas[4] = 7.5

	fmt.Println("Notas:", notas)
	fmt.Printf("Tamanho: %d\n", len(notas))
	fmt.Printf("Primeira nota: %.1f\n", notas[0])
	fmt.Printf("Última nota: %.1f\n", notas[len(notas)-1])

	// Declaração com inicialização direta
	primos := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Primos:", primos)

	// O compilador conta os elementos automaticamente com ...
	vogais := [...]string{"a", "e", "i", "o", "u"}
	fmt.Println("Vogais:", vogais)
	fmt.Printf("Tipo: %T\n", vogais) // [5]string — tamanho faz parte do tipo

	// Array multidimensional (matriz)
	fmt.Println("\n=== ARRAY MULTIDIMENSIONAL ===")

	var matriz [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = i*3 + j + 1
		}
	}

	for _, linha := range matriz {
		fmt.Println(linha)
	}

	// Arrays são copiados por valor — cuidado
	fmt.Println("\n=== ARRAY CÓPIA POR VALOR ===")

	original := [3]int{1, 2, 3}
	copia := original  // copia tudo — são independentes
	copia[0] = 99

	fmt.Println("Original:", original) // [1 2 3] — não foi alterado
	fmt.Println("Cópia:   ", copia)    // [99 2 3]

	// =========================================================================
	// SLICE — tamanho DINÂMICO (o que você vai usar quase sempre)
	// =========================================================================
	//
	// Slice não é um array. É uma "visão" sobre um array subjacente.
	// Internamente tem três campos:
	//   - ponteiro para o array base
	//   - len (comprimento atual)
	//   - cap (capacidade — quanto pode crescer sem realocar)
	//
	// Sintaxe: var nome []tipo  (sem número entre os colchetes)
	//
	// =========================================================================

	fmt.Println("\n=== SLICE ===")

	// Declaração vazia — valor zero de slice é nil
	var s []int
	fmt.Printf("Slice nil: %v | len=%d | nil=%t\n", s, len(s), s == nil)

	// Criação com literal
	frutas := []string{"maçã", "banana", "laranja"}
	fmt.Println("Frutas:", frutas)
	fmt.Printf("len=%d | cap=%d\n", len(frutas), cap(frutas))

	// -------------------------------------------------------------------------
	// append — adiciona elementos ao slice
	// IMPORTANTE: append pode retornar um NOVO slice se precisar realocar.
	// Sempre reatribua: frutas = append(frutas, ...)
	// -------------------------------------------------------------------------

	fmt.Println("\n=== APPEND ===")

	numeros := []int{1, 2, 3}
	fmt.Printf("Antes:  %v | len=%d | cap=%d\n", numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 4)
	numeros = append(numeros, 5, 6) // pode adicionar múltiplos de uma vez
	fmt.Printf("Depois: %v | len=%d | cap=%d\n", numeros, len(numeros), cap(numeros))

	// Append de um slice em outro com ...
	extras := []int{7, 8, 9}
	numeros = append(numeros, extras...)
	fmt.Println("Com extras:", numeros)

	// -------------------------------------------------------------------------
	// make — cria slice com len e cap definidos
	// Usado quando você sabe o tamanho antecipado (evita realocações)
	// -------------------------------------------------------------------------

	fmt.Println("\n=== MAKE ===")

	// make([]tipo, len, cap)
	s2 := make([]int, 3, 6) // 3 elementos (zeros), capacidade para 6
	fmt.Printf("make: %v | len=%d | cap=%d\n", s2, len(s2), cap(s2))

	s2[0] = 10
	s2[1] = 20
	s2[2] = 30
	fmt.Println("Preenchido:", s2)

	// -------------------------------------------------------------------------
	// FATIAMENTO — slice de slice
	// sintaxe: slice[inicio:fim]
	// fim é EXCLUSIVO — slice[1:4] pega índices 1, 2, 3
	// -------------------------------------------------------------------------

	fmt.Println("\n=== FATIAMENTO (slicing) ===")

	letras := []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println("Original:     ", letras)
	fmt.Println("letras[1:4]   ", letras[1:4]) // b c d
	fmt.Println("letras[:3]    ", letras[:3])   // a b c (início omitido = 0)
	fmt.Println("letras[3:]    ", letras[3:])   // d e f (fim omitido = len)
	fmt.Println("letras[:]     ", letras[:])    // cópia completa

	// ATENÇÃO: slice fatiado compartilha o array base
	// Modificar um afeta o outro
	fmt.Println("\n=== SLICE COMPARTILHA MEMÓRIA ===")

	base := []int{1, 2, 3, 4, 5}
	parte := base[1:4] // [2 3 4]

	parte[0] = 99 // modifica base[1] também!

	fmt.Println("base: ", base)  // [1 99 3 4 5] — foi alterado!
	fmt.Println("parte:", parte) // [99 3 4]

	// Para evitar isso, use copy
	fmt.Println("\n=== COPY — cópia independente ===")

	origem := []int{10, 20, 30, 40, 50}
	destino := make([]int, len(origem))
	copiados := copy(destino, origem)

	destino[0] = 999
	fmt.Println("origem:  ", origem)   // não foi alterado
	fmt.Println("destino: ", destino)  // [999 20 30 40 50]
	fmt.Printf("elementos copiados: %d\n", copiados)

	// -------------------------------------------------------------------------
	// REMOVENDO ELEMENTOS DE UM SLICE
	// Go não tem função built-in de remoção. Usa fatiamento + append.
	// -------------------------------------------------------------------------

	fmt.Println("\n=== REMOVER ELEMENTO ===")

	itens := []string{"a", "b", "c", "d", "e"}
	remover := 2 // índice a remover (o "c")

	itens = append(itens[:remover], itens[remover+1:]...)
	fmt.Println("Após remover índice 2:", itens) // [a b d e]

	// -------------------------------------------------------------------------
	// ITERANDO SOBRE SLICE
	// -------------------------------------------------------------------------

	fmt.Println("\n=== ITERAÇÃO ===")

	cores := []string{"vermelho", "verde", "azul"}

	for i, cor := range cores {
		fmt.Printf("[%d] %s\n", i, cor)
	}

	// =========================================================================
	// STRINGS — detalhes importantes
	// =========================================================================
	//
	// String em Go é uma sequência IMUTÁVEL de bytes (não caracteres).
	// Internamente é um slice de bytes ([]byte), mas não pode ser modificada.
	//
	// Consequência: "café" tem 5 bytes, mas 4 caracteres (runes).
	// O 'é' ocupa 2 bytes em UTF-8.
	//
	// =========================================================================

	fmt.Println("\n=== STRINGS ===")

	texto := "Olá, Go!"

	fmt.Printf("String: %s\n", texto)
	fmt.Printf("len() em bytes: %d\n", len(texto))                        // bytes
	fmt.Printf("Caracteres (runes): %d\n", utf8.RuneCountInString(texto)) // runes

	// Acessar por índice retorna BYTE, não caractere
	fmt.Printf("texto[0] = %d (byte) = %c (char)\n", texto[0], texto[0])

	// Para iterar por caractere (rune), use range
	fmt.Println("\n--- range em string ---")
	for i, r := range texto {
		fmt.Printf("byte[%d] = %c (rune %d)\n", i, r, r)
	}

	// Converter string ↔ []byte ↔ []rune
	fmt.Println("\n=== CONVERSÕES DE STRING ===")

	s3 := "hello"
	bytes3 := []byte(s3)    // string → []byte (pode modificar)
	runes3 := []rune(s3)    // string → []rune (correto para Unicode)

	bytes3[0] = 'H'         // modificar o []byte não afeta s3
	deVolta := string(bytes3)

	fmt.Println("Original:", s3)
	fmt.Println("[]byte modificado → string:", deVolta)
	fmt.Printf("[]rune: %v\n", runes3)

	// -------------------------------------------------------------------------
	// PACOTE strings — operações comuns
	// -------------------------------------------------------------------------

	fmt.Println("\n=== PACOTE strings ===")

	frase := "  Go é uma linguagem incrível  "

	fmt.Printf("TrimSpace:   %q\n", strings.TrimSpace(frase))
	fmt.Printf("ToUpper:     %s\n", strings.ToUpper(frase))
	fmt.Printf("ToLower:     %s\n", strings.ToLower(frase))
	fmt.Printf("Contains:    %t\n", strings.Contains(frase, "linguagem"))
	fmt.Printf("HasPrefix:   %t\n", strings.HasPrefix("Golang", "Go"))
	fmt.Printf("HasSuffix:   %t\n", strings.HasSuffix("main.go", ".go"))
	fmt.Printf("Replace:     %s\n", strings.Replace(frase, "Go", "Rust", 1))
	fmt.Printf("ReplaceAll:  %s\n", strings.ReplaceAll("aababab", "ab", "X"))
	fmt.Printf("Count:       %d\n", strings.Count("banana", "a"))
	fmt.Printf("Index:       %d\n", strings.Index("golang", "lang")) // posição ou -1
	fmt.Printf("Repeat:      %s\n", strings.Repeat("Go! ", 3))

	// Split e Join
	fmt.Println("\n--- Split e Join ---")

	csv := "maçã,banana,laranja,uva"
	partes := strings.Split(csv, ",")
	fmt.Println("Split:", partes)

	junto := strings.Join(partes, " | ")
	fmt.Println("Join:", junto)

	// Fields — divide por qualquer espaço em branco (espaços, tabs, newlines)
	palavras := strings.Fields("  um   dois   três  ")
	fmt.Println("Fields:", palavras)

	// -------------------------------------------------------------------------
	// STRINGS SÃO IMUTÁVEIS — para construir strings dinamicamente use strings.Builder
	// Concatenar com + em loop cria uma nova string a cada iteração (ineficiente)
	// -------------------------------------------------------------------------

	fmt.Println("\n=== strings.Builder (concatenação eficiente) ===")

	var sb strings.Builder

	for i := 1; i <= 5; i++ {
		fmt.Fprintf(&sb, "item%d ", i)
	}

	resultado := sb.String()
	fmt.Println(resultado)

	// -------------------------------------------------------------------------
	// SLICE DE STRINGS — padrão muito comum
	// -------------------------------------------------------------------------

	fmt.Println("\n=== SLICE DE STRINGS ===")

	nomes := []string{"Ana", "Bruno", "Carlos", "Diana"}

	// Buscar elemento
	busca := "Carlos"
	encontrado := false
	for _, n := range nomes {
		if n == busca {
			encontrado = true
			break
		}
	}
	fmt.Printf("'%s' encontrado: %t\n", busca, encontrado)

	// Filtrar elementos (cria novo slice)
	longos := []string{}
	for _, n := range nomes {
		if len(n) > 4 {
			longos = append(longos, n)
		}
	}
	fmt.Println("Nomes com mais de 4 letras:", longos)
}