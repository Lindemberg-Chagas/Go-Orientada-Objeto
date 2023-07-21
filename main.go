package main

import (
	"banco/contas"
	"banco/contas/clientes"
	"fmt"
)

func PagarBoleto(conta VerficaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerficaConta interface {
	Sacar(valor float64) string
}

func main() {
	contaLindemberg := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Lindemberg", CPF: "123.123.123-12"}}

	fmt.Println(contaLindemberg)
	fmt.Println(contaLindemberg.ObterSaldo())
	contaLindemberg.Depositar(100)
	PagarBoleto(&contaLindemberg, 40)

}
