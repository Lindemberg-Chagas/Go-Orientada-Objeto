package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *contaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Efetivado"
	} else {
		return "Saldo Insuficente"
	}
}
func (c *contaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito Efetivado"
	} else {
		return "Valor de Deposito não permitido!"
	}
}

func (c *contaCorrente) Transfere(valorDaTransfere float64, contaDestino *contaCorrente) string {
	if valorDaTransfere > 0 && valorDaTransfere <= c.saldo {
		c.saldo -= valorDaTransfere
		contaDestino.Depositar(valorDaTransfere)
		return "Transferencia Efetivada"

	} else {
		return "Não foi possivel realizar a tranferencia"
	}
}

func main() {
	contaLindemberg := contaCorrente{}
	contaLindemberg.titular = "Lindemberg"
	contaLindemberg.saldo = 150.50

	contaBerg := contaCorrente{}
	contaBerg.titular = "Berg"
	contaBerg.saldo = 50.0

	//fmt.Println(contaLindemberg.saldo)
	//fmt.Println("Saque 100.00")
	//fmt.Println(contaLindemberg.Sacar(100))
	//fmt.Println("O saldo atualizado é")
	//fmt.Println(contaLindemberg.saldo)
	//fmt.Println("Deposite 200.00")
	//fmt.Println(contaLindemberg.Depositar(200))
	//fmt.Println("O saldo atualizado é")
	//fmt.Println(contaLindemberg.saldo)

	fmt.Println("Saldo em Conta é")
	fmt.Println(contaLindemberg.saldo)
	fmt.Println("Realizar Transferencia")
	fmt.Println(contaLindemberg.Transfere(100, &contaBerg))
	fmt.Println("O saldo atualizado é")
	fmt.Println(contaLindemberg.saldo)
}
