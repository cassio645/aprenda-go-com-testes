package carteira

import (
	"testing"
)

// Teste depositar e ver saldo da carteira
func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)
	saldo := carteira.Saldo()

	esperado := 10

	if saldo != esperado {
		t.Errorf("saldo %d, esperado %d", saldo, esperado)
	}
}
