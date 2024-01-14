// Pacote carteira com funções de depositar e consultar saldo.
package carteira

// Struct carteira com atributo privado saldo
type Carteira struct {
	saldo int
}

// Método de carteira para depositar, usando ponteiro para carteira
func (c *Carteira) Depositar(quantidade int) {
	c.saldo += quantidade
}

// Método de carteira para conferir o saldo, usando ponteiro
func (c *Carteira) Saldo() int {
	return c.saldo
}
