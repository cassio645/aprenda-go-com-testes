// Pacote sync possui um tipo Contador que pode ser usado de maneira concorrente.
package sincroniza

import "sync"

// Tipo contador
type Contador struct {
	mu sync.Mutex
	valor int
}

// NovoContador retorna um novo Contador
func NovoContador() *Contador {
	return &Contador{}
}

// Função que incrementa o valor usando Mutex para travar o contador quando executado de maneira concorrente.
func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

// Função que retorna o valor atual do contador
func (c *Contador) Valor() int {
	return c.valor
}
