package main

import (
	"fmt"
	"math"
)

// =============================================================================
// POR QUE INTERFACES EXISTEM? (leia isso antes de qualquer código)
// =============================================================================
//
// Imagine que você está fazendo seu jogo e tem Jogador, Monstro, Chefe, Totem...
// Todos eles podem receber dano. Se você quiser criar uma função "atacar", você
// seria forçado a criar uma versão separada pra cada tipo:
//
//   func atacarJogador(j *Jogador, dano int)  { ... }
//   func atacarMonstro(m *Monstro, dano int)  { ... }
//   func atacarChefe(c *Chefe, dano int)       { ... }
//
// Isso é código duplicado, difícil de manter, e quebra toda vez que você
// adiciona um novo tipo. Com interfaces, você escreve UMA função:
//
//   func atacar(alvo Combatente, dano int) { ... }
//
// E qualquer coisa que "saiba receber dano" pode ser passada pra ela.
// Interface é um CONTRATO: "se você tiver esses métodos, pode entrar aqui".
//
// Em resumo, interfaces servem pra:
//   1. Evitar repetição de código
//   2. Fazer funções que aceitam tipos diferentes mas com comportamento parecido
//   3. Deixar o código mais fácil de expandir sem quebrar o que já existe
//   4. Separar O QUE algo faz de COMO ele faz

// =============================================================================
// PARTE 1 — O QUE É UMA INTERFACE
// =============================================================================
//
// Uma interface define um conjunto de ASSINATURAS de métodos.
// Assinatura = nome do método + parâmetros + o que ele retorna.
// A interface não tem nenhum código dentro dela, só as assinaturas.
//
// SINTAXE:
//
//   type NomeDaInterface interface {
//       Metodo1(param tipo) retorno
//       Metodo2() retorno
//   }
//
// Qualquer tipo que tiver TODOS os métodos listados satisfaz a interface.
// Você não escreve "implements" em nenhum lugar. É automático.

// Interface com um único método:
type Combatente interface {
	ReceberDano(dano int)
}

// Interface com vários métodos:
type Entidade interface {
	ReceberDano(dano int)
	Curar(quantidade int)
	EstaVivo() bool
	Nome() string
}

// =============================================================================
// PARTE 2 — IMPLEMENTANDO UMA INTERFACE (sem declarar nada)
// =============================================================================
//
// Um tipo satisfaz uma interface automaticamente se tiver todos os métodos.
// Não tem "implements", não tem "extends", não tem nenhuma palavra-chave extra.
// Go verifica isso sozinho em tempo de compilação.

// --- Tipo 1: Jogador ---
type Jogador struct {
	nome string
	hp   int
	vivo bool
}

// Jogador tem todos os 4 métodos que Entidade exige → satisfaz Entidade
func (j *Jogador) ReceberDano(dano int) {
	j.hp -= dano
	if j.hp <= 0 {
		j.hp = 0
		j.vivo = false
	}
}

func (j *Jogador) Curar(quantidade int) {
	if j.vivo {
		j.hp += quantidade
	}
}

func (j *Jogador) EstaVivo() bool {
	return j.vivo
}

func (j *Jogador) Nome() string {
	return j.nome
}

// --- Tipo 2: Monstro ---
type Monstro struct {
	nome string
	hp   int
}

// Monstro também satisfaz Entidade
func (m *Monstro) ReceberDano(dano int) {
	m.hp -= dano
	if m.hp < 0 {
		m.hp = 0
	}
}

func (m *Monstro) Curar(quantidade int) {
	m.hp += quantidade
}

func (m *Monstro) EstaVivo() bool {
	return m.hp > 0
}

func (m *Monstro) Nome() string {
	return m.nome
}

// --- Tipo 3: Totem (objeto destrutível do cenário, não cura) ---
type Totem struct {
	nome string
	hp   int
}

func (t *Totem) ReceberDano(dano int) {
	t.hp -= dano
	if t.hp < 0 {
		t.hp = 0
	}
}

func (t *Totem) Curar(quantidade int) {
	// Totem não cura. O método existe pra satisfazer a interface, mas não faz nada.
	// Isso é válido em Go. A interface só exige que o método EXISTA.
}

func (t *Totem) EstaVivo() bool {
	return t.hp > 0
}

func (t *Totem) Nome() string {
	return t.nome
}

// =============================================================================
// PARTE 3 — USANDO A INTERFACE EM FUNÇÕES
// =============================================================================
//
// Aqui está o ponto principal: agora que Jogador, Monstro e Totem satisfazem
// Entidade, você pode passar qualquer um deles pra uma função que aceita Entidade.

// Uma função que aceita qualquer Entidade:
func atacar(atacante string, alvo Entidade, dano int) {
	if !alvo.EstaVivo() {
		fmt.Printf("%s já está destruído/morto.\n", alvo.Nome())
		return
	}
	alvo.ReceberDano(dano)
	if alvo.EstaVivo() {
		fmt.Printf("%s causou %d de dano em %s.\n", atacante, dano, alvo.Nome())
	} else {
		fmt.Printf("%s causou %d de dano em %s. %s foi eliminado!\n",
			atacante, dano, alvo.Nome(), alvo.Nome())
	}
}

// Função que cura qualquer Entidade viva:
func curar(alvo Entidade, quantidade int) {
	if !alvo.EstaVivo() {
		fmt.Printf("Não é possível curar %s, está morto.\n", alvo.Nome())
		return
	}
	alvo.Curar(quantidade)
	fmt.Printf("%s foi curado em %d.\n", alvo.Nome(), quantidade)
}

// Função que imprime o status de uma lista de entidades.
// "[]Entidade" = slice de Entidade. Pode misturar Jogador, Monstro, Totem...
func mostrarStatus(entidades []Entidade) {
	fmt.Println("--- STATUS ---")
	for _, e := range entidades {
		status := "vivo"
		if !e.EstaVivo() {
			status = "morto/destruído"
		}
		fmt.Printf("  %-10s → %s\n", e.Nome(), status)
	}
}

// =============================================================================
// PARTE 4 — INTERFACE VAZIA: interface{}  /  any
// =============================================================================
//
// "interface{}" (ou "any", que é um apelido) é uma interface sem nenhum método.
// Todo tipo em Go satisfaz ela automaticamente.
// Use quando você genuinamente não sabe que tipo vai receber.
//
// CUIDADO: não use "any" pra tudo só porque é conveniente. Você perde a
// verificação de tipos do compilador e o código fica difícil de entender.

func imprimirQualquerCoisa(valor any) {
	fmt.Printf("Valor: %v | Tipo: %T\n", valor, valor)
}

// =============================================================================
// PARTE 5 — TYPE ASSERTION (afirmar o tipo real)
// =============================================================================
//
// Quando você tem uma interface, o compilador não sabe o tipo concreto dentro dela.
// Type assertion é como você diz: "eu sei que isso é um *Jogador, me dê o acesso".
//
// SINTAXE:
//   concreto, ok := variavel.(TipoConcreto)
//
// "ok" é um bool: true se a conversão funcionou, false se não.
// SEMPRE use a forma com "ok". Se não usar e o tipo for errado, o programa vai travar.

func verificarSeJogador(e Entidade) {
	jogador, ok := e.(*Jogador) // tenta afirmar que e é *Jogador
	if ok {
		fmt.Printf("%s é um Jogador com HP: %d\n", jogador.nome, jogador.hp)
	} else {
		fmt.Printf("%s não é um Jogador.\n", e.Nome())
	}
}

// =============================================================================
// PARTE 6 — TYPE SWITCH (verificar vários tipos de uma vez)
// =============================================================================
//
// Quando você pode receber tipos diferentes e quer agir diferente pra cada um,
// use um type switch. É como um switch normal, mas testando tipos.

func descreverEntidade(e Entidade) {
	switch v := e.(type) { // v recebe o valor já convertido pro tipo correto
	case *Jogador:
		fmt.Printf("[Jogador] %s — HP: %d\n", v.nome, v.hp)
	case *Monstro:
		fmt.Printf("[Monstro] %s — HP: %d\n", v.nome, v.hp)
	case *Totem:
		fmt.Printf("[Totem]   %s — HP: %d\n", v.nome, v.hp)
	default:
		fmt.Printf("[Desconhecido] %s\n", e.Nome())
	}
}

// =============================================================================
// PARTE 7 — COMPOSIÇÃO DE INTERFACES
// =============================================================================
//
// Interfaces podem ser combinadas em interfaces maiores.
// Isso permite criar interfaces específicas reutilizando partes menores.
// É o mesmo conceito de composição que vimos em structs.

type Atacante interface {
	Atacar(alvo Entidade, dano int)
}

type Movivel interface {
	Mover(x, y float64)
}

// EntidadeCompleta exige tudo: Entidade + Atacante + Movivel
type EntidadeCompleta interface {
	Entidade
	Atacante
	Movivel
}

// =============================================================================
// PARTE 8 — EXEMPLO MAIS REALISTA: formas geométricas
// =============================================================================
//
// Um exemplo clássico que mostra bem o poder das interfaces.
// Cada forma tem sua própria fórmula de área, mas a interface unifica tudo.

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

type Triangulo struct {
	Base, Altura, LadoA, LadoB, LadoC float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	return t.LadoA + t.LadoB + t.LadoC
}

// Uma única função que funciona pra qualquer forma:
func imprimirForma(f Forma) {
	fmt.Printf("  Área: %.2f | Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// =============================================================================
// MAIN
// =============================================================================

func main() {
	fmt.Println("======================================================")
	fmt.Println("  PARTE 2 — SATISFAZENDO A INTERFACE")
	fmt.Println("======================================================")

	// Criando valores concretos:
	j := &Jogador{nome: "Herói", hp: 100, vivo: true}
	m := &Monstro{nome: "Goblin", hp: 50}
	t := &Totem{nome: "Totem da Cura", hp: 30}

	// Todos podem ser usados onde Entidade é esperado:
	var e1 Entidade = j
	var e2 Entidade = m
	var e3 Entidade = t

	descreverEntidade(e1)
	descreverEntidade(e2)
	descreverEntidade(e3)

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  PARTE 3 — FUNÇÕES COM INTERFACE")
	fmt.Println("======================================================")

	atacar("Arqueiro", m, 20)
	atacar("Mago", j, 40)
	atacar("Guerreiro", t, 35)
	curar(j, 25)

	fmt.Println()
	// Slice misturando tipos diferentes — só funciona porque todos são Entidade:
	todos := []Entidade{j, m, t}
	mostrarStatus(todos)

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  PARTE 4 — INTERFACE VAZIA")
	fmt.Println("======================================================")

	imprimirQualquerCoisa(42)
	imprimirQualquerCoisa("texto")
	imprimirQualquerCoisa(true)
	imprimirQualquerCoisa(j) // até uma struct funciona

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  PARTE 5 — TYPE ASSERTION")
	fmt.Println("======================================================")

	verificarSeJogador(j)
	verificarSeJogador(m)

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  PARTE 6 — TYPE SWITCH")
	fmt.Println("======================================================")

	entidades := []Entidade{j, m, t}
	for _, e := range entidades {
		descreverEntidade(e)
	}

	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println("  PARTE 8 — FORMAS GEOMÉTRICAS")
	fmt.Println("======================================================")

	formas := []Forma{
		Circulo{Raio: 5},
		Retangulo{Largura: 4, Altura: 6},
		Triangulo{Base: 3, Altura: 4, LadoA: 3, LadoB: 4, LadoC: 5},
	}

	nomes := []string{"Círculo", "Retângulo", "Triângulo"}
	for i, f := range formas {
		fmt.Printf("%s:\n", nomes[i])
		imprimirForma(f)
	}
}

// =============================================================================
// SIMULAÇÃO DE PACOTE SEPARADO
// =============================================================================
//
// Na prática, structs e interfaces costumam ficar em pacotes separados.
// Abaixo está um comentário mostrando como ficaria a estrutura de arquivos
// e como a main usaria esse pacote.
//
// -----------------------------------------------------------------------------
// ESTRUTURA DE ARQUIVOS:
//
//   meujogo/
//   ├── main.go
//   └── entidade/
//       ├── entidade.go    ← define a interface
//       └── jogador.go     ← define a struct e seus métodos
//
// -----------------------------------------------------------------------------
// ARQUIVO: entidade/entidade.go
//
//   package entidade
//
//   // A interface fica neste pacote.
//   // Letra MAIÚSCULA = exportado (visível fora do pacote). Regra fundamental em Go.
//   type Entidade interface {
//       ReceberDano(dano int)
//       Curar(quantidade int)
//       EstaVivo() bool
//       Nome() string
//   }
//
// -----------------------------------------------------------------------------
// ARQUIVO: entidade/jogador.go
//
//   package entidade
//
//   // Struct exportada (maiúscula).
//   type Jogador struct {
//       nome string   // campo em minúsculo = privado (só acessível dentro do pacote)
//       hp   int
//       vivo bool
//   }
//
//   // Construtor exportado. Convenção: "New" + NomeDoTipo.
//   func NewJogador(nome string, hp int) *Jogador {
//       return &Jogador{nome: nome, hp: hp, vivo: true}
//   }
//
//   // Métodos que satisfazem a interface Entidade:
//   func (j *Jogador) ReceberDano(dano int) { ... }
//   func (j *Jogador) Curar(quantidade int)  { ... }
//   func (j *Jogador) EstaVivo() bool        { return j.vivo }
//   func (j *Jogador) Nome() string          { return j.nome }
//
// -----------------------------------------------------------------------------
// ARQUIVO: main.go
//
//   package main
//
//   import (
//       "fmt"
//       "meujogo/entidade"   // importando o pacote pelo caminho do módulo
//   )
//
//   func atacar(alvo entidade.Entidade, dano int) {
//       // "entidade.Entidade" = tipo Entidade dentro do pacote entidade
//       alvo.ReceberDano(dano)
//   }
//
//   func main() {
//       // Usando o construtor do pacote pra criar um *Jogador:
//       j := entidade.NewJogador("Herói", 100)
//
//       // j é *entidade.Jogador, que satisfaz entidade.Entidade automaticamente.
//       // Então pode ser passado diretamente pra função atacar:
//       atacar(j, 30)
//
//       fmt.Println(j.EstaVivo()) // true
//   }
//
// -----------------------------------------------------------------------------
// RESUMO DO QUE ACONTECE:
//
//   1. entidade.NewJogador(...)  → retorna *entidade.Jogador
//   2. *entidade.Jogador tem ReceberDano, Curar, EstaVivo, Nome
//   3. entidade.Entidade exige ReceberDano, Curar, EstaVivo, Nome
//   4. Go verifica: todos os métodos batem? → sim → Jogador satisfaz Entidade
//   5. Você passa j pra atacar(alvo entidade.Entidade, ...) sem nenhum cast
//
// Essa separação em pacotes é o padrão real em projetos Go.
// Cada entidade do seu jogo (Jogador, Monstro, Chefe) pode ter seu próprio arquivo
// dentro do pacote, e todos satisfazem a mesma interface sem depender uns dos outros.

// =============================================================================
// RESUMO GERAL DE INTERFACES
// =============================================================================
//
//   DEFINIÇÃO:
//     type MinhaInterface interface { Metodo() retorno }
//
//   IMPLEMENTAÇÃO (automática, sem palavra-chave):
//     func (t *MeuTipo) Metodo() retorno { ... }
//
//   USO EM FUNÇÃO:
//     func fazer(x MinhaInterface) { x.Metodo() }
//
//   SLICE DE INTERFACE (mistura de tipos):
//     itens := []MinhaInterface{&TipoA{}, &TipoB{}}
//
//   TYPE ASSERTION (recuperar tipo concreto):
//     concreto, ok := x.(*MeuTipo)
//
//   TYPE SWITCH (vários tipos de uma vez):
//     switch v := x.(type) { case *TipoA: ... case *TipoB: ... }
//
//   INTERFACE VAZIA (aceita tudo):
//     func f(v any) { ... }   // use com cautela
//
//   COMPOSIÇÃO DE INTERFACES:
//     type Grande interface { InterfaceA; InterfaceB }
//
//   PACOTES (convenção de nomenclatura):
//     - Tipo/Interface exportado → começa com MAIÚSCULA: Entidade, Jogador
//     - Tipo/campo privado      → começa com minúscula: nome, hp
//     - Construtor              → NewNomeDoTipo(...)