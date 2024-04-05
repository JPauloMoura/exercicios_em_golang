package main

import "fmt"

/*
Você está escrevendo uma parte do nosso sistema anti-fraude de pagamentos.

Uma das formas de detectarmos fraudes é quando o mesmo cliente faz mais de um pagamento numa janela muito curta de tempo.

Seu programa receberá os clientes que estão efetuando pagamentos através de um stream, e deve identificar clientes repetidos que apareçam próximos.

Escreva um algoritmo que leia um stream de clientes, onde cada cliente é identificado por um número inteiro,
 e escreva na saída cada vez que um cliente aparecer repetido numa janela de tamanho J.

**Exemplo:**
J = 4

Stream = [7, 10, 5, 10, 8, 3, 1, 4, 3, 3, 5, 1]

Saída: 10, 3, 3
*/

func main() {
	stream := []int{7, 10, 5, 10, 8, 3, 1, 4, 3, 3, 5, 1}
	transactions := NewTransactionsByClient(4)

	for _, client := range stream {
		if !transactions.ProcessTransaction(client) {
			fmt.Println(client)
		}
	}

}

type clientId int
type transactionRepeted int

// TransactionsByClient é a estrutura responsável por lidar com as transações de clientes
// para identifica transações repetidas em um espaço de tempo
type TransactionsByClient struct {
	// maxComparableTransactions é a quantidade maxima de transações anteriores que vamos comparar com a atual.
	maxComparableTransactions int

	// duplicationsByClient registra a quantidade de transações repetidas por cliente
	duplicationsByClient map[clientId]transactionRepeted
	// transactions registra a ultimos transações dos clientes
	transactions []clientId
}

func NewTransactionsByClient(maxComparableTransactions int) TransactionsByClient {
	return TransactionsByClient{
		maxComparableTransactions: maxComparableTransactions,
		duplicationsByClient:      make(map[clientId]transactionRepeted),
		transactions:              make([]clientId, 0),
	}
}

// ProcessTransaction realiza a analise da transação do cliente, retornando false caso essa transação
// já tenha sido realizada em um curto periodo de tempo
func (t *TransactionsByClient) ProcessTransaction(client int) (sucess bool) {
	if t.maxLimit() {
		t.removeOlder()
	}

	return t.add(clientId(client))
}

func (v TransactionsByClient) maxLimit() bool {
	return len(v.transactions) >= v.maxComparableTransactions
}

func (v *TransactionsByClient) add(client clientId) bool {
	v.transactions = append(v.transactions, client)
	sucess := true

	if v.isDuplicated(client) {
		v.duplicationsByClient[client] += 1
		return !sucess
	}

	v.duplicationsByClient[client] = 0
	return sucess
}

func (v TransactionsByClient) isDuplicated(client clientId) (repedated bool) {
	_, repedated = v.duplicationsByClient[client]
	return
}

func (v *TransactionsByClient) removeOlder() {
	client := v.removeFirstTransaction()

	numberOfRepetitions := v.duplicationsByClient[client]

	if numberOfRepetitions < 1 {
		delete(v.duplicationsByClient, client)
		return
	}

	v.duplicationsByClient[client] -= 1
}

func (v *TransactionsByClient) removeFirstTransaction() (clientRemoved clientId) {
	clientRemoved = v.transactions[0]
	v.transactions = v.transactions[1:]
	return
}
