package main

import (
	"banco/contas"
	"banco/contas/clientes"
	"fmt"
)

func main() {
	contaLindemberg := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Lindemberg", CPF: "123.123.123-12"}}

	fmt.Println(contaLindemberg)
	fmt.Println(contaLindemberg.ObterSaldo())
	contaLindemberg.Depositar(100)
	fmt.Println(contaLindemberg.ObterSaldo())

}
