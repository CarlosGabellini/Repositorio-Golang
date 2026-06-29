package main

import "fmt"

// =============================================================================
// FUNÇÕES EM GO
// =============================================================================
//
// Uma função é um bloco de código com um nome, que você pode "chamar" (executar)
// quando quiser. Evita repetir o mesmo código em vários lugares.
//
// SINTAXE BÁSICA:
//
//   func nomeDaFuncao(parametro1 tipo1, parametro2 tipo2) tipoDeRetorno {
//       // corpo da função
//       return valor
//   }
//
// - "func"           → palavra-chave que inicia toda função
// - "nomeDaFuncao"   → você escolhe o nome
// - parâmetros       → dados que a função recebe pra trabalhar (pode não ter nenhum)
// - tipoDeRetorno    → tipo do valor que a função devolve (pode não devolver nada)
// - "return"         → devolve o resultado pra quem chamou a função

// -----------------------------------------------------------------------------
// 1. Função sem parâmetros e sem retorno
// -----------------------------------------------------------------------------
// Só executa código. Não recebe nada, não devolve nada.

func dizerOla() {
	fmt.Println("Olá, mundo!")
}

// -----------------------------------------------------------------------------
// 2. Função com parâmetros
// -----------------------------------------------------------------------------
// Recebe valores pra usar dentro dela.
// "nome string" → parâmetro chamado "nome" do tipo string

func cumprimentar(nome string) {
	fmt.Println("Olá,", nome)
}

// -----------------------------------------------------------------------------
// 3. Função com retorno
// -----------------------------------------------------------------------------
// Devolve um valor. O tipo do retorno fica depois dos parênteses.
// Quem chamou a função pode guardar esse valor numa variável.

func somar(a int, b int) int {
	return a + b
}

// ATALHO: quando dois parâmetros têm o mesmo tipo, você pode escrever assim:
func subtrair(a, b int) int {
	return a - b
}

// -----------------------------------------------------------------------------
// 4. Retorno múltiplo — Go permite devolver mais de um valor ao mesmo tempo
// -----------------------------------------------------------------------------
// Muito comum em Go para devolver (resultado, erro) juntos.

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil // "nil" significa "sem erro"
}

// -----------------------------------------------------------------------------
// 5. Retorno nomeado
// -----------------------------------------------------------------------------
// Você pode dar nomes aos valores de retorno. Eles viram variáveis locais.
// Ao chamar "return" sozinho (naked return), Go devolve os valores nomeados.
// Útil para clareza, mas não abuse — pode confundir em funções longas.

func calcularAreaPerimetro(largura, altura float64) (area float64, perimetro float64) {
	area = largura * altura
	perimetro = 2 * (largura + altura)
	return // "naked return": devolve area e perimetro como estão
}

// -----------------------------------------------------------------------------
// 6. Variadic — aceita número ilimitado de argumentos do mesmo tipo
// -----------------------------------------------------------------------------
// "nums ...int" significa "aceita zero ou mais ints"
// Dentro da função, "nums" é tratado como um slice (fatia) de ints

func somarTodos(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// -----------------------------------------------------------------------------
// 7. Funções como valores (first-class functions)
// -----------------------------------------------------------------------------
// Em Go, funções são valores como qualquer outro. Você pode:
//   - guardar uma função numa variável
//   - passar uma função como argumento pra outra função
//   - retornar uma função

// Esta função recebe OUTRA FUNÇÃO como parâmetro.
// "operacao func(int, int) int" → parâmetro que é uma função que recebe dois ints e devolve um int
func calcular(a, b int, operacao func(int, int) int) int {
	return operacao(a, b)
}

// -----------------------------------------------------------------------------
// 8. Closure (função anônima que "captura" variáveis do escopo externo)
// -----------------------------------------------------------------------------
// Uma closure é uma função sem nome que "lembra" variáveis do lugar onde foi criada.
// Muito usada quando você precisa de uma função temporária ou customizada.

func criarContador() func() int {
	contagem := 0        // esta variável fica "presa" dentro da closure
	return func() int { // função anônima: sem nome, definida inline
		contagem++
		return contagem
	}
}

// =============================================================================
// STRUCTS EM GO
// =============================================================================
//
// Uma struct é uma forma de agrupar dados relacionados num único tipo.
// Pense como uma "ficha" com campos: uma ficha de personagem tem nome, HP, nível, etc.
//
// Go NÃO tem classes como Java/C++. Structs + métodos cumprem esse papel.
//
// SINTAXE:
//
//   type NomeDaStruct struct {
//       campo1 tipo1
//       campo2 tipo2
//   }

// -----------------------------------------------------------------------------
// 9. Struct simples
// -----------------------------------------------------------------------------

type Jogador struct {
	Nome  string
	HP    int
	Nivel int
	Vivo  bool
}

// -----------------------------------------------------------------------------
// 10. Struct com struct embutida (composição)
// -----------------------------------------------------------------------------
// Go não tem herança. Usa COMPOSIÇÃO: uma struct dentro da outra.
// Isso é chamado de "embedding" (embutir).
// Os campos e métodos da struct embutida ficam acessíveis diretamente.

type Posicao struct {
	X float64
	Y float64
}

type Entidade struct {
	Posicao        // embutida: Entidade.X e Entidade.Y ficam disponíveis diretamente
	Velocidade float64
}

// -----------------------------------------------------------------------------
// 11. Métodos em structs
// -----------------------------------------------------------------------------
// Um método é uma função "ligada" a um tipo. O receptor fica entre "func" e o nome.
//
//   func (receptor TipoReceptor) NomeDoMetodo(params) retorno { ... }
//
// RECEPTOR POR VALOR  → func (j Jogador)  → trabalha em uma CÓPIA da struct
// RECEPTOR POR PONTEIRO → func (j *Jogador) → trabalha na struct ORIGINAL
//
// Regra geral:
//   - Use ponteiro (*T) quando o método precisa MODIFICAR a struct
//   - Use valor  ( T) quando só precisa LER os dados

// Receptor por valor: só lê, não modifica
func (j Jogador) Apresentar() {
	fmt.Printf("Jogador: %s | HP: %d | Nível: %d\n", j.Nome, j.HP, j.Nivel)
}

// Receptor por ponteiro: modifica a struct original
func (j *Jogador) ReceberDano(dano int) {
	j.HP -= dano
	if j.HP <= 0 {
		j.HP = 0
		j.Vivo = false
		fmt.Printf("%s morreu!\n", j.Nome)
	}
}

func (j *Jogador) Curar(quantidade int) {
	if !j.Vivo {
		fmt.Printf("%s está morto e não pode ser curado.\n", j.Nome)
		return
	}
	j.HP += quantidade
	fmt.Printf("%s curou %d de HP. HP atual: %d\n", j.Nome, quantidade, j.HP)
}

func (j *Jogador) SubirNivel() {
	j.Nivel++
	j.HP += 20 // bônus de HP por nível
	fmt.Printf("%s subiu para o nível %d!\n", j.Nome, j.Nivel)
}

// -----------------------------------------------------------------------------
// 12. Construtor (por convenção, não é palavra-chave)
// -----------------------------------------------------------------------------
// Go não tem "new" como em outras linguagens (bem, tem, mas não é usual assim).
// O padrão é criar uma função chamada "NovoX" que devolve uma instância da struct.
// Retornar ponteiro (*Jogador) é comum quando a struct vai ser modificada.

func NovoJogador(nome string, nivel int) *Jogador {
	return &Jogador{
		Nome:  nome,
		HP:    100 + (nivel * 20),
		Nivel: nivel,
		Vivo:  true,
	}
}

// -----------------------------------------------------------------------------
// 13. Interfaces
// -----------------------------------------------------------------------------
// Uma interface define um CONTRATO: "qualquer tipo que tenha esses métodos
// satisfaz essa interface". Go usa duck typing implícito — você não declara
// explicitamente que implementa uma interface.
//
// Se um tipo tem todos os métodos que a interface exige → ele implementa a interface.
// Simples assim.

type Entidade_Combate interface {
	ReceberDano(dano int)
	Apresentar()
}

// Jogador já tem ReceberDano e Apresentar, então satisfaz Entidade_Combate.

// Outra entidade diferente que também satisfaz a mesma interface:
type Monstro struct {
	Nome string
	HP   int
	Dano int
}

func (m *Monstro) ReceberDano(dano int) {
	m.HP -= dano
	if m.HP <= 0 {
		m.HP = 0
		fmt.Printf("Monstro %s foi derrotado!\n", m.Nome)
	}
}

func (m Monstro) Apresentar() {
	fmt.Printf("Monstro: %s | HP: %d | Dano: %d\n", m.Nome, m.HP, m.Dano)
}

// Função que aceita QUALQUER coisa que satisfaça Entidade_Combate
func atacar(atacante string, alvo Entidade_Combate, dano int) {
	fmt.Printf("%s ataca por %d de dano!\n", atacante, dano)
	alvo.ReceberDano(dano)
}

// -----------------------------------------------------------------------------
// 14. Struct com slice de structs (relação "tem vários")
// -----------------------------------------------------------------------------

type Inventario struct {
	Itens []string
	Ouro  int
}

type PersonagemCompleto struct {
	Jogador           // embutindo Jogador dentro de PersonagemCompleto
	Inv     Inventario
}

func (p *PersonagemCompleto) AdicionarItem(item string) {
	p.Inv.Itens = append(p.Inv.Itens, item)
	fmt.Printf("%s adicionou '%s' ao inventário.\n", p.Nome, item)
}

func (p PersonagemCompleto) MostrarInventario() {
	fmt.Printf("Inventário de %s: %v | Ouro: %d\n", p.Nome, p.Inv.Itens, p.Inv.Ouro)
}

// =============================================================================
// MAIN — onde tudo roda
// =============================================================================

func main() {
	fmt.Println("======================================================")
	fmt.Println("  DEMONSTRAÇÃO: FUNÇÕES")
	fmt.Println("======================================================")

	// --- Funções básicas ---
	dizerOla()
	cumprimentar("Carlos")

	resultado := somar(10, 5)
	fmt.Println("10 + 5 =", resultado)

	fmt.Println("10 - 3 =", subtrair(10, 3))

	// --- Retorno múltiplo ---
	quociente, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quociente)
	}

	_, errDiv := dividir(5, 0) // underscore "_" descarta o primeiro retorno
	if errDiv != nil {
		fmt.Println("Erro capturado:", errDiv)
	}

	// --- Retorno nomeado ---
	area, perimetro := calcularAreaPerimetro(4, 6)
	fmt.Printf("Área: %.1f | Perímetro: %.1f\n", area, perimetro)

	// --- Variadic ---
	fmt.Println("Soma variadic (1,2,3,4,5):", somarTodos(1, 2, 3, 4, 5))
	fmt.Println("Soma variadic ():", somarTodos()) // pode chamar sem argumentos

	// --- Função como valor ---
	multiplicar := func(a, b int) int { return a * b } // função anônima guardada em variável
	fmt.Println("calcular(3, 4, somar):", calcular(3, 4, somar))
	fmt.Println("calcular(3, 4, multiplicar):", calcular(3, 4, multiplicar))

	// --- Closure ---
	contador := criarContador()
	fmt.Println("Contador:", contador()) // 1
	fmt.Println("Contador:", contador()) // 2
	fmt.Println("Contador:", contador()) // 3

	outroContador := criarContador() // closure independente, começa do zero
	fmt.Println("Outro contador:", outroContador()) // 1

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  DEMONSTRAÇÃO: STRUCTS")
	fmt.Println("======================================================")

	// --- Criando structs ---

	// Forma 1: literal com nomes dos campos (recomendado, ordem não importa)
	j1 := Jogador{Nome: "Herói", HP: 100, Nivel: 1, Vivo: true}
	j1.Apresentar()

	// Forma 2: usando o construtor (mais seguro, centraliza lógica)
	j2 := NovoJogador("Arqueiro", 3)
	j2.Apresentar()

	// --- Métodos modificando a struct ---
	fmt.Println()
	j2.ReceberDano(30)
	j2.Curar(15)
	j2.SubirNivel()
	j2.Apresentar()

	// --- Cenário de morte ---
	fmt.Println()
	j1.ReceberDano(999)
	j1.Curar(50) // tenta curar morto

	// --- Interface ---
	fmt.Println()
	fmt.Println("--- Combate via interface ---")
	monstro := &Monstro{Nome: "Goblin", HP: 50, Dano: 8}
	monstro.Apresentar()

	atacar("Herói", monstro, 20)
	atacar("Herói", monstro, 20)
	atacar("Herói", monstro, 20)

	// --- Composição e embedding ---
	fmt.Println()
	fmt.Println("--- PersonagemCompleto ---")
	pc := PersonagemCompleto{
		Jogador: Jogador{Nome: "Mago", HP: 80, Nivel: 2, Vivo: true},
		Inv:     Inventario{Ouro: 150},
	}

	// Graças ao embedding, acessamos campos e métodos de Jogador diretamente:
	pc.Apresentar()     // método de Jogador, acessado via PersonagemCompleto
	pc.SubirNivel()     // idem
	pc.AdicionarItem("Poção")
	pc.AdicionarItem("Espada de ferro")
	pc.MostrarInventario()

	// --- Struct com composição (Posicao embutida em Entidade) ---
	fmt.Println()
	e := Entidade{
		Posicao:    Posicao{X: 10.5, Y: 3.2},
		Velocidade: 2.5,
	}
	// Graças ao embedding, X e Y ficam acessíveis diretamente:
	fmt.Printf("Entidade em X=%.1f, Y=%.1f, vel=%.1f\n", e.X, e.Y, e.Velocidade)
	// Também funciona via e.Posicao.X, mas o shorthand e.X é mais limpo

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  FIM DA DEMONSTRAÇÃO")
	fmt.Println("======================================================")
}

// =============================================================================
// RESUMO MENTAL
// =============================================================================
//
// FUNÇÕES:
//   func nome(params) retorno { ... }   → básico
//   func nome() (a, b tipo) { ... }     → retorno múltiplo / nomeado
//   func nome(vals ...tipo) { ... }     → variadic (0 ou mais args)
//   variavel := func(...) { ... }       → função anônima / closure
//
// STRUCTS:
//   type X struct { campo tipo }        → define um novo tipo
//   x := X{campo: valor}               → cria instância
//   func (x X) Metodo() { ... }        → receptor por valor  (só lê)
//   func (x *X) Metodo() { ... }       → receptor por ponteiro (modifica)
//   type Y struct { X; ... }           → embedding (composição, não herança)
//
// INTERFACES:
//   type I interface { Metodo() }      → contrato implícito
//   Qualquer tipo com Metodo() satisfaz I automaticamente
//
// PONTEIROS (&, *):
//   &variavel   → endereço de memória da variável (vira ponteiro)
//   *ponteiro   → valor apontado (desreferência)
//   Use *T nos métodos que precisam modificar a struct original