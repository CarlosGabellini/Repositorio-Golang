package main

import (
	"errors"
	"fmt"
	"strconv"
)

// =============================================================================
// PARTE 1 — ERROR HANDLING (tratamento de erros)
// =============================================================================
//
// Em Go não existe try/catch. Erros são VALORES normais, retornados pela função.
// Você trata o erro ali mesmo, na linha seguinte à chamada.
//
// Por que assim? Porque erros escondidos em exceções são difíceis de rastrear.
// Go te força a lidar com o erro na hora, explicitamente. Chato no começo,
// mas muito mais claro quando o código cresce.
//
// O tipo de erro padrão é a interface "error":
//
//   type error interface {
//       Error() string
//   }
//
// Qualquer tipo que tenha o método Error() string satisfaz a interface error.

// =============================================================================
// 1.1 — RETORNANDO ERROS COM fmt.Errorf E errors.New
// =============================================================================
//
// Duas formas de criar um erro simples:
//
//   errors.New("mensagem")       → erro estático, mensagem fixa
//   fmt.Errorf("valor: %d", x)  → erro com formatação, como um Printf

var ErrDivisaoPorZero = errors.New("divisão por zero")
// "var Err..." com maiúscula é a convenção para erros que serão comparados fora do pacote.
// Chamados de "sentinel errors" (erros sentinela).

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero // devolve zero + o erro
	}
	return a / b, nil // nil = sem erro
}

func raizQuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("raiz quadrada de número negativo: %.2f", n)
	}
	// math.Sqrt existe, mas aqui vamos calcular via Newton só pra exemplificar
	resultado := n
	for i := 0; i < 50; i++ {
		resultado = (resultado + n/resultado) / 2
	}
	return resultado, nil
}

// =============================================================================
// 1.2 — ERRO CUSTOMIZADO (struct que implementa a interface error)
// =============================================================================
//
// Quando você precisa carregar mais informação no erro além de uma string,
// crie um tipo próprio implementando Error() string.
//
// Útil quando quem chamou a função precisa inspecionar os detalhes do erro.

type ErrConversao struct {
	Valor   string // o valor que falhou
	TipoAlvo string // o tipo que tentamos converter
	Causa   error  // o erro original
}

func (e *ErrConversao) Error() string {
	return fmt.Sprintf("não foi possível converter '%s' para %s: %v",
		e.Valor, e.TipoAlvo, e.Causa)
}

// Unwrap expõe o erro interno, necessário para errors.Is / errors.As funcionar
// com erros empacotados dentro de structs.
func (e *ErrConversao) Unwrap() error {
	return e.Causa
}

func converterParaInt(s string) (int, error) {
	n, err := strconv.Atoi(s) // tenta converter string → int
	if err != nil {
		// empacota o erro original dentro do nosso tipo customizado
		return 0, &ErrConversao{
			Valor:    s,
			TipoAlvo: "int",
			Causa:    err,
		}
	}
	return n, nil
}

// =============================================================================
// 1.3 — errors.Is e errors.As
// =============================================================================
//
// errors.Is  → verifica se um erro (ou qualquer erro que ele embrulhe) É um
//              valor específico. Usado com erros sentinela.
//
// errors.As  → verifica se um erro (ou embrulho) É de um tipo específico, e se
//              for, preenche uma variável com ele. Usado com erros customizados.
//
// Por que não comparar com "=="?
// Porque erros podem ser embrulhados em camadas (wrapping). errors.Is/As
// atravessam essas camadas. Comparar com == só verifica a camada de cima.

func demonstrarIsAs() {
	fmt.Println("--- errors.Is e errors.As ---")

	// errors.Is com sentinel error:
	_, err := dividir(10, 0)
	if errors.Is(err, ErrDivisaoPorZero) {
		fmt.Println("errors.Is: confirmado, é ErrDivisaoPorZero")
	}

	// errors.As com erro customizado:
	_, err2 := converterParaInt("abc")
	var errConv *ErrConversao
	if errors.As(err2, &errConv) {
		// errConv agora tem os campos do nosso ErrConversao
		fmt.Printf("errors.As: valor='%s', tipo='%s'\n", errConv.Valor, errConv.TipoAlvo)
	}
}

// =============================================================================
// 1.4 — WRAPPING (embrulhar erros)
// =============================================================================
//
// Quando você propaga um erro adicionando contexto, usa %w no fmt.Errorf.
// Isso embrulha o erro original dentro do novo, mantendo a cadeia.
// errors.Is e errors.As conseguem navegar por essa cadeia.
//
//   fmt.Errorf("contexto: %w", errOriginal)   → embrulha (use %w)
//   fmt.Errorf("contexto: %v", errOriginal)   → só formata como string, NÃO embrulha

func buscarUsuario(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("buscarUsuario: %w", ErrDivisaoPorZero) // só pra exemplificar wrapping
	}
	if id == 999 {
		return "", fmt.Errorf("buscarUsuario id=%d: usuário não encontrado", id)
	}
	return "Carlos", nil
}

func carregarPerfil(id int) (string, error) {
	usuario, err := buscarUsuario(id)
	if err != nil {
		// adicionamos contexto ao erro e o embrulhamos com %w
		return "", fmt.Errorf("carregarPerfil: %w", err)
	}
	return "Perfil de " + usuario, nil
}

// =============================================================================
// 1.5 — PADRÕES COMUNS NO DIA A DIA
// =============================================================================

// Padrão 1: retorno antecipado (early return)
// Trate o erro logo e retorne. Não aninha ifs desnecessariamente.
func processarDados(input string) (int, error) {
	n, err := converterParaInt(input)
	if err != nil {
		return 0, fmt.Errorf("processarDados: %w", err) // retorno antecipado
	}
	// a partir daqui, n é válido
	return n * 2, nil
}

// Padrão 2: ignorar erro intencionalmente com _
// Use SOMENTE quando você tem certeza que o erro não pode ocorrer.
// Em produção, ignorar erros sem motivo é um bug esperando acontecer.
func exemploIgnorarErro() {
	resultado, _ := dividir(10, 2) // sabemos que 2 != 0, então ok ignorar
	fmt.Println("resultado:", resultado)
}

// Padrão 3: erro fatal com panic (use com moderação)
// panic para o programa. Só use quando o estado é irrecuperável.
// Em aplicações web/servidor, quase nunca use panic.
func deveExistir(valor string) string {
	if valor == "" {
		panic("valor obrigatório não pode ser vazio") // estado irrecuperável
	}
	return valor
}

// =============================================================================
// PARTE 2 — GENERICS
// =============================================================================
//
// Generics (Go 1.18+) permite escrever funções e structs que funcionam com
// VÁRIOS tipos sem repetir código, mantendo a verificação de tipos do compilador.
//
// Antes de generics, você tinha duas opções ruins:
//   1. Escrever uma função pra cada tipo (repetição)
//   2. Usar "any" (interface vazia) e perder a segurança de tipos
//
// Generics resolve os dois problemas.
//
// SINTAXE:
//
//   func Nome[T Restricao](param T) T { ... }
//
//   T          → parâmetro de tipo (placeholder). Você escolhe o nome (T, K, V, etc.)
//   Restricao  → quais tipos T pode ser (define o "contrato" do tipo genérico)
//   [T Restricao] → lista de parâmetros de tipo, fica entre colchetes após o nome

// =============================================================================
// 2.1 — RESTRIÇÕES (constraints)
// =============================================================================
//
// Uma restrição é uma interface que limita quais tipos podem ser usados.
// Go já tem algumas prontas no pacote "constraints" (ou você cria a sua).

// Restrição customizada: qualquer tipo numérico inteiro
type Inteiro interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

// Restrição customizada: qualquer tipo numérico (inteiro ou float)
type Numero interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// "comparable" é uma restrição embutida no Go.
// Tipos comparable podem ser comparados com == e !=
// (int, string, bool, structs de campos comparáveis, etc.)

// =============================================================================
// 2.2 — FUNÇÕES GENÉRICAS
// =============================================================================

// Sem generics: você precisaria de somarInts, somarFloats, somarInt32...
// Com generics: uma única função funciona pra todos os tipos numéricos.
func Somar[T Numero](a, b T) T {
	return a + b
}

// Retorna o maior de dois valores (qualquer tipo ordenável)
func Maior[T Numero](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Retorna o menor valor de um slice (qualquer tipo numérico)
func Minimo[T Numero](valores []T) (T, error) {
	var zero T // zero value do tipo T
	if len(valores) == 0 {
		return zero, errors.New("slice vazio")
	}
	min := valores[0]
	for _, v := range valores[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// Verifica se um slice contém um valor (qualquer tipo comparable)
func Contem[T comparable](slice []T, valor T) bool {
	for _, v := range slice {
		if v == valor {
			return true
		}
	}
	return false
}

// Filtra um slice mantendo só os elementos onde a função predicado retorna true.
// Dois parâmetros de tipo: T é o tipo dos elementos.
func Filtrar[T any](slice []T, predicado func(T) bool) []T {
	resultado := []T{}
	for _, v := range slice {
		if predicado(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

// Transforma cada elemento de um slice (map/transform clássico).
// Dois parâmetros de tipo: T (tipo de entrada), U (tipo de saída).
func Mapear[T any, U any](slice []T, transformar func(T) U) []U {
	resultado := make([]U, len(slice))
	for i, v := range slice {
		resultado[i] = transformar(v)
	}
	return resultado
}

// =============================================================================
// 2.3 — STRUCTS GENÉRICAS
// =============================================================================
//
// Structs também podem ser genéricas. Útil para coleções e contêineres.

// Par genérico: guarda dois valores de tipos possivelmente diferentes.
type Par[K, V any] struct {
	Chave  K
	Valor  V
}

func (p Par[K, V]) String() string {
	return fmt.Sprintf("(%v → %v)", p.Chave, p.Valor)
}

// Pilha genérica (stack): last in, first out.
// Funciona com qualquer tipo.
type Pilha[T any] struct {
	itens []T
}

func (p *Pilha[T]) Empilhar(v T) {
	p.itens = append(p.itens, v)
}

func (p *Pilha[T]) Desempilhar() (T, error) {
	var zero T
	if len(p.itens) == 0 {
		return zero, errors.New("pilha vazia")
	}
	topo := p.itens[len(p.itens)-1]
	p.itens = p.itens[:len(p.itens)-1]
	return topo, nil
}

func (p *Pilha[T]) Tamanho() int {
	return len(p.itens)
}

// =============================================================================
// 2.4 — COMBINANDO GENERICS COM ERROR HANDLING
// =============================================================================
//
// Um tipo Result[T] que representa "ou um valor T, ou um erro".
// Padrão comum em outras linguagens (Rust, Kotlin), implementado via generics em Go.

type Result[T any] struct {
	valor T
	err   error
}

func Ok[T any](v T) Result[T] {
	return Result[T]{valor: v}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) Desempacotar() (T, error) {
	return r.valor, r.err
}

func (r Result[T]) EhErro() bool {
	return r.err != nil
}

// Função que usa Result[T] como retorno:
func dividirSeguro(a, b float64) Result[float64] {
	if b == 0 {
		return Err[float64](ErrDivisaoPorZero)
	}
	return Ok(a / b)
}

// =============================================================================
// MAIN
// =============================================================================

func main() {

	// -------------------------------------------------------------------------
	fmt.Println("======================================================")
	fmt.Println("  ERROR HANDLING")
	fmt.Println("======================================================")

	// Caso normal:
	if r, err := dividir(10, 3); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", r)
	}

	// Erro sentinela:
	_, err := dividir(10, 0)
	fmt.Println("Erro divisão:", err)
	fmt.Println("É ErrDivisaoPorZero?", errors.Is(err, ErrDivisaoPorZero))

	// Erro customizado:
	_, err2 := converterParaInt("xyz")
	fmt.Println("Erro conversão:", err2)

	demonstrarIsAs()

	// Wrapping:
	fmt.Println()
	fmt.Println("--- Wrapping ---")
	_, err3 := carregarPerfil(999)
	fmt.Println("Erro:", err3)

	_, err4 := carregarPerfil(-1)
	fmt.Println("Erro embrulhado:", err4)
	// errors.Is ainda consegue encontrar o sentinela mesmo embrulhado:
	fmt.Println("Contém ErrDivisaoPorZero?", errors.Is(err4, ErrDivisaoPorZero))

	// Early return:
	fmt.Println()
	fmt.Println("--- processarDados ---")
	if n, err := processarDados("21"); err == nil {
		fmt.Println("processarDados('21'):", n)
	}
	if _, err := processarDados("abc"); err != nil {
		fmt.Println("processarDados('abc'):", err)
	}

	// -------------------------------------------------------------------------
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  GENERICS")
	fmt.Println("======================================================")

	// Funções genéricas — Go infere o tipo automaticamente na maioria dos casos:
	fmt.Println("--- Somar ---")
	fmt.Println("int:    ", Somar(3, 4))
	fmt.Println("float64:", Somar(1.5, 2.5))
	fmt.Println("int64:  ", Somar(int64(100), int64(200)))

	fmt.Println()
	fmt.Println("--- Maior ---")
	fmt.Println("Maior(7, 3):", Maior(7, 3))
	fmt.Println("Maior(2.9, 3.1):", Maior(2.9, 3.1))

	fmt.Println()
	fmt.Println("--- Minimo ---")
	ints := []int{5, 2, 8, 1, 9}
	if min, err := Minimo(ints); err == nil {
		fmt.Println("Minimo de", ints, "→", min)
	}
	floats := []float64{3.3, 1.1, 2.2}
	if min, err := Minimo(floats); err == nil {
		fmt.Printf("Minimo de %v → %.1f\n", floats, min)
	}

	fmt.Println()
	fmt.Println("--- Contem ---")
	nomes := []string{"Ana", "Bruno", "Carlos"}
	fmt.Println("Contem 'Carlos':", Contem(nomes, "Carlos"))
	fmt.Println("Contem 'Diego':", Contem(nomes, "Diego"))
	numeros := []int{10, 20, 30}
	fmt.Println("Contem 20:", Contem(numeros, 20))

	fmt.Println()
	fmt.Println("--- Filtrar ---")
	pares := Filtrar([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println("Pares:", pares)
	longos := Filtrar(nomes, func(s string) bool { return len(s) > 3 })
	fmt.Println("Nomes > 3 letras:", longos)

	fmt.Println()
	fmt.Println("--- Mapear ---")
	dobros := Mapear([]int{1, 2, 3, 4}, func(n int) int { return n * 2 })
	fmt.Println("Dobros:", dobros)
	// Transforma []int em []string:
	strs := Mapear([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("item%d", n) })
	fmt.Println("Como strings:", strs)

	fmt.Println()
	fmt.Println("--- Par (struct genérica) ---")
	p1 := Par[string, int]{Chave: "idade", Valor: 17}
	p2 := Par[int, bool]{Chave: 42, Valor: true}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println()
	fmt.Println("--- Pilha (struct genérica) ---")
	pilha := &Pilha[string]{}
	pilha.Empilhar("primeiro")
	pilha.Empilhar("segundo")
	pilha.Empilhar("terceiro")
	fmt.Println("Tamanho:", pilha.Tamanho())
	for pilha.Tamanho() > 0 {
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilhou:", v)
	}
	_, err5 := pilha.Desempilhar()
	fmt.Println("Desempilhar vazia:", err5)

	fmt.Println()
	fmt.Println("--- Result[T] (generics + error handling) ---")
	r1 := dividirSeguro(10, 4)
	if v, err := r1.Desempacotar(); err == nil {
		fmt.Printf("10 / 4 = %.2f\n", v)
	}
	r2 := dividirSeguro(10, 0)
	if r2.EhErro() {
		_, err := r2.Desempacotar()
		fmt.Println("Erro:", err)
	}
}

// =============================================================================
// RESUMO
// =============================================================================
//
// ERROR HANDLING:
//   - Erros são valores, retornados como último valor da função
//   - nil = sem erro
//   - errors.New("msg")              → erro simples
//   - fmt.Errorf("ctx: %w", err)     → erro com contexto + wrapping (%w)
//   - tipo customizado + Error()     → erro com campos inspecionáveis
//   - errors.Is(err, Sentinela)      → compara atravessando wrapping
//   - errors.As(err, &tipoConcreto)  → extrai tipo concreto atravessando wrapping
//   - panic                          → só para estados irrecuperáveis
//
// GENERICS:
//   - func Nome[T Restricao](p T) T  → função genérica
//   - type X[T any] struct { ... }   → struct genérica
//   - Restricao = interface com tipos permitidos (int | float64 | ...)
//   - comparable                     → tipos que suportam == e !=
//   - any                            → qualquer tipo (sem restrição)
//   - Go infere T automaticamente na maioria das chamadas
//   - Use generics quando você repetiria código mudando só o tipo
//   - Não use generics onde interfaces comuns já resolvem