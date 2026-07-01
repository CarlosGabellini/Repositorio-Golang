package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================================
// CONCORRÊNCIA EM GO PARA DESENVOLVIMENTO DE JOGOS
// ============================================================================
// Se você quer criar jogos eficientes, precisa entender que concorrência NÃO é 
// paralelismo. Concorrência é a composição estrutural de processos que rodam
// de forma independente. 
// No Ebitengine, o Game Loop (Update/Draw) roda a 60 FPS. Se você travar a 
// thread principal com um carregamento de arquivo ou requisição de rede, 
// o frame rate cai para zero. É aqui que entra a concorrência.
// ============================================================================

func main() {
	fmt.Println("=== 1. GOROUTINES (O Básico) ===")
	// Uma goroutine é uma thread leve gerenciada pelo Go runtime.
	// Basta colocar a palavra-chave 'go' antes de uma chamada de função.
	go carregarAssetInseguro("Textura_Player.png")
	go carregarAssetInseguro("Sons_Ambiente.mp3")

	// ERRO CLÁSSICO: Se o escopo principal (main) terminar, o programa fecha
	// e mata todas as goroutines em background, mesmo que não tenham terminado.
	// Por enquanto, usamos um Sleep tolo apenas para demonstração visual.
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n=== 2. CHANNELS (Canais de Comunicação) ===")
	// Goroutines não devem compartilhar memória sem proteção. Canais são os
	// canos que permitem enviar e receber dados entre elas com segurança.
	// Criamos um canal que trafega strings:
	canalEventos := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		// O operador '<-' envia dados para o canal
		canalEventos <- "Inimigo Gerado no background!"
	}()

	// O código abaixo BLOQUEIA a execução do main até receber algo do canal.
	// Isso é útil para sincronizar eventos no início do frame do jogo.
	evento := <-canalEventos
	fmt.Println("Recebido no Game Loop:", evento)

	fmt.Println("\n=== 3. WAITGROUP (Sincronização Profissional) ===")
	// Gambiarra com time.Sleep não existe em código de produção.
	// Usamos sync.WaitGroup para esperar que um grupo de tarefas termine.
	var wg sync.WaitGroup

	// Dizemos ao WaitGroup que vamos esperar por 2 tarefas
	wg.Add(2)

	go func() {
		defer wg.Done() // Garante que decrementa o contador ao terminar
		fmt.Println("[Tarefa] Física do mapa inicializada.")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("[Tarefa] Shaders de iluminação compilados.")
	}()

	// Bloqueia aqui até que o contador interno do WaitGroup chegue a ZERO.
	wg.Wait()
	fmt.Println("-> Todas as dependências carregadas. Jogo pronto!")

	fmt.Println("\n=== 4. MUTEX (Proteção contra Data Race) ===")
	// Em jogos, múltiplas goroutines tentando alterar o mesmo dado (ex: a pontuação
	// ou a posição de um personagem) geram um "Data Race" (corrida de dados). Isso
	// corrompe a memória do jogo.
	// O sync.Mutex garante exclusão mútua: apenas UMA goroutine acessa o dado por vez.
	
	var mu sync.Mutex
	pontuacaoGlobal := 0
	var wgProgresso sync.WaitGroup

	// Vamos simular 100 inimigos morrendo ao mesmo tempo e atualizando o score
	InimigosMortos := 100
	wgProgresso.Add(InimigosMortos)

	for i := 0; i < InimigosMortos; i++ {
		go func() {
			defer wgProgresso.Done()
			
			// Se você tirar o Lock/Unlock e rodar com 'go run -race', o Go vai gritar.
			mu.Lock()   // Tranca o acesso para as outras goroutines
			pontuacaoGlobal++
			mu.Unlock() // Libera o acesso
		}()
	}

	wgProgresso.Wait()
	fmt.Printf("Pontuação Final Segura: %d\n", pontuacaoGlobal)
}

func carregarAssetInseguro(nome string) {
	fmt.Printf("[Background] Carregando: %s\n", nome)
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("[Background] Concluído: %s\n", nome)
}