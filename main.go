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

func main() {
	contaLindemberg := contaCorrente{}
	contaLindemberg.titular = "Lindemberg"
	contaLindemberg.saldo = 150.50

	fmt.Println(contaLindemberg.saldo)
	fmt.Println(contaLindemberg.Sacar(100))
	fmt.Println(contaLindemberg.saldo)

}
