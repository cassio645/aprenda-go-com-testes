package inteiros

import (
	"fmt"
	"testing"
)

// TestAdicionador testa a função Adiciona passando dois valores e compara com o resultado esperado.
func TestAdicionador(t *testing.T) {
	resultado := Adiciona(2, 2)
	esperado := 4

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}

// Example de como usar a função Adiciona
func ExampleAdiciona() {
	resultado := Adiciona(10, 5)
	fmt.Println(resultado)
	// Output: 15

}
