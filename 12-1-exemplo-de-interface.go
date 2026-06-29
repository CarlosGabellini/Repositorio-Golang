package main

import "fmt"

// 1. A Interface original (o "contrato")
type MetodoPagamento interface {
	Pagar(valor float64)
}

// 2. A função que processa a compra (que permanece INTACTA)
func ProcessarCompra(total float64, m MetodoPagamento) {
	// Ela não quer saber COMO o pagamento é feito, apenas manda pagar
	m.Pagar(total) 
}

// ==========================================================
// NOVA STRUCT: "Crypto-Pagamento" adicionada depois
// ==========================================================

type CryptoPagamento struct {
	CarteiraDestino string
	Moeda           string // ex: "BTC", "ETH"
}

// Criando o método para que ela assine o contrato da interface
func (cp CryptoPagamento) Pagar(valor float64) {
	fmt.Printf("Pagamento de R$ %.2f realizado com sucesso!\n", valor)
	fmt.Printf("Enviando fundos em %s para a carteira: %s\n", cp.Moeda, cp.CarteiraDestino)
}

// ==========================================================

func main() {
	// O valor total da compra do cliente
	totalCarrinho := 250.50

	// Instanciamos a nova struct com os dados necessários
	pagamentoBitcoin := CryptoPagamento{
		CarteiraDestino: "0x71C...3a9",
		Moeda:           "BTC",
	}

	// Passamos a nova struct para a função antiga.
	// Funciona perfeitamente sem termos alterado uma única linha de ProcessarCompra!
	ProcessarCompra(totalCarrinho, pagamentoBitcoin)
}