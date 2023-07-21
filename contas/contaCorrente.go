package contas

import (
	"banco/contas/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Efetivado"
	} else {
		return "saldo Insuficente"
	}
}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito Efetivado"
	} else {
		return "Valor de Deposito não permitido!"
	}
}

func (c *ContaCorrente) Transfere(valorDaTransfere float64, contaDestino *ContaCorrente) string {
	if valorDaTransfere > 0 && valorDaTransfere <= c.saldo {
		c.saldo -= valorDaTransfere
		contaDestino.Depositar(valorDaTransfere)
		return "Transferencia Efetivada"

	} else {
		return "Não foi possivel realizar a tranferencia"
	}
}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
