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


// Benchmark para função Repetir, calcula a média de tempo em nano segundos para a função executar.
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("Oi")
	}
}
