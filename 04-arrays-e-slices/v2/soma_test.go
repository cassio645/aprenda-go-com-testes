package soma

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {

	validacao := func(t *testing.T, resultado int, esperado int, numeros []int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dados os numeros %v", resultado, esperado, numeros)
		}
	}

	t.Run("Testa um slice de 5 elementos", func(t *testing.T) {
		numeros := []int{2, 3, 4, 5, 6}
		resultado := Soma(numeros)
		esperado := 20

		validacao(t, resultado, esperado, numeros)
	})

	t.Run("Testa slice de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		validacao(t, resultado, esperado, numeros)
	})

}

func ExampleSoma(){
	resultado := Soma([]int{10, 20, 30, 40})
	fmt.Println(resultado)
	// Output: 100
}
