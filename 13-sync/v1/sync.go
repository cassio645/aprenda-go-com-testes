// Pacote sync possui um tipo contador que incrementa o valor que pode ser consultado.
package sync

// Tipo contador
type Contador struct {
	valor int
}

// Função que incrementa o valor do contador
func (c *Contador) Incrementar() {
	c.valor++
}

// Função que retorna o valor atual do contador
func (c *Contador) Valor() int {
	return c.valor
}
