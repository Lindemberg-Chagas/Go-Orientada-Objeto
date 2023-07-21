package contas

type ContaPoupanca struct {
	Titular                              string
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Efetivado"
	} else {
		return "saldo Insuficente"
	}
}
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito Efetivado"
	} else {
		return "Valor de Deposito não permitido!"
	}
}

func (c *ContaPoupanca) Transfere(valorDaTransfere float64, contaDestino *ContaCorrente) string {
	if valorDaTransfere > 0 && valorDaTransfere <= c.saldo {
		c.saldo -= valorDaTransfere
		contaDestino.Depositar(valorDaTransfere)
		return "Transferencia Efetivada"

	} else {
		return "Não foi possivel realizar a tranferencia"
	}
}
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
