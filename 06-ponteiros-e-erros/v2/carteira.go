// Pacote carteira de Bitcoin com funções de depositar retirar e consultar saldo.
package carteira

import (
	"errors"
	"fmt"
)

// Criando tipo Bitcoin
type Bitcoin int

// Criando uma variável do tipo erro, com a mensagem 'saldo insuficiente'
var ErroSaldoInsuficiente = errors.New("Saldo insuficiente.")

// Definindo como aparece uma variavel do tipo Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Struct carteira com atributo privado saldo do tipo Bitcoin
type Carteira struct {
	saldo Bitcoin
}

// Método de carteira para depositar, usando ponteiro para carteira
func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

// Método Retira o valor informado do saldo da carteira
// E retorna um erro quando a quantidade é maior que o saldo atual na carteira
func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= quantidade
	return nil
}

// Método de carteira para conferir o saldo, usando ponteiro
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func ExampleDepositar() {
	carteira := Carteira{}
	carteira.Depositar(Bitcoin(20))
	fmt.Println(carteira.Saldo())
	// Output: 20 BTC
}

func ExampleSaldo(){
	carteira := Carteira{saldo: Bitcoin(10)}
	fmt.Println(carteira.Saldo())
	// Output: 10 BTC
}

func ExampleRetirar() {
	carteira := Carteira{saldo: Bitcoin(30)}
	err := carteira.Retirar(Bitcoin(15))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carteira.Saldo())
	// Output: 15 BTC
}