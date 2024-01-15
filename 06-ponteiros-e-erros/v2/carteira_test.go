package carteira

import (
	"fmt"
	"testing"
)

// Testando as funções de depositar e retirar
func TestCarteira(t *testing.T) {
	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()

		saldo := carteira.Saldo()

		if saldo != esperado {
			t.Errorf("saldo %s, esperado %s", saldo, esperado)
		}
	}

	// Quando queremos que aconteça um erro
	confirmaErro := func(t *testing.T, resultado error, esperado error) {
		t.Helper()
		if resultado == nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu.")
		}

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	// Quando não queremos um erro, mas ele acontece
	confirmaErroInexistnete := func(t *testing.T, resultado error) {
		t.Helper()
		if resultado != nil {
			t.Fatal("Erro inesperado recebido.")
		}
	}

	t.Run("Teste Depositar 15", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(15))

		confirmaSaldo(t, carteira, Bitcoin(15))
	})

	t.Run("Teste Retirar 5", func(t *testing.T) {

		// Criando uma nova carteira com saldo 15
		carteira := Carteira{saldo: Bitcoin(15)}

		erro := carteira.Retirar(5)

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistnete(t, erro) // não queremos erro
	})

	t.Run("Teste Retirar mais do que pode", func(t *testing.T) {
		saldoInicial := Bitcoin(10)
		carteira := Carteira{saldoInicial}

		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente) // queremos erro
	})

}

func ExampleCarteira_Depositar() {
	carteira := Carteira{}
	carteira.Depositar(Bitcoin(20))
	fmt.Println(carteira.Saldo())
	// Output: 20 BTC
}

func ExampleCarteira_Saldo() {
	carteira := Carteira{saldo: Bitcoin(10)}
	fmt.Println(carteira.Saldo())
	// Output: 10 BTC
}

func ExampleCarteira_Retirar() {
	carteira := Carteira{saldo: Bitcoin(30)}
	err := carteira.Retirar(Bitcoin(15))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carteira.Saldo())
	// Output: 15 BTC
}
