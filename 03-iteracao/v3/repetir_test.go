package iteracao

import (
	"fmt"
	"testing"
)

// Usando subtestes para testar a função repetir informando a quantidade de vezes.
func TestRepetir(t *testing.T) {

	verificaValor := func(t *testing.T, resposta, esperado string){
		t.Helper()
		if resposta != esperado {
			t.Errorf("resposta %s, esperado %s", resposta, esperado)
		}
	}

	// Com argumento deve repretir 10x
	t.Run("Repetindo 10x a string", func(t *testing.T) {
		resultado := Repetir("Oi", 10)
		esperado := "OiOiOiOiOiOiOiOiOiOi"

		verificaValor(t, resultado, esperado)
	})

	// Sem argumento deve usar o valor default, e repetir 5x
	t.Run("Repetindo 10x a string", func(t *testing.T) {
		resultado := Repetir("Oi", 0)
		esperado := "OiOiOiOiOi"

		verificaValor(t, resultado, esperado)
	})


}

func ExampleRepetir() {
	fmt.Println(Repetir("Oi", 0))
	fmt.Println(Repetir("a", 10))
	// Output:
	// OiOiOiOiOi
	// aaaaaaaaaa	
}


// Benchmark para função Repetir, calcula a média de tempo em nano segundos para a função executar.
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("Oi", 0)
	}
}
