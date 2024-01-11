package iteracao

import (
	"fmt"
	"testing"
)

// Testando a função Repetir, enviando uma string e verificando se ela retorna repetindo 5x.
func TestRepetir(t *testing.T) {
	resposta := Repetir("a")
	esperado := "aaaaa"

	if resposta != esperado {
		t.Errorf("resposta %s, esperado %s", resposta, esperado)
	}
}

func ExampleRepetir() {
	resposta := Repetir("Oi")
	fmt.Println(resposta)
	// Output: OiOiOiOiOi
}
