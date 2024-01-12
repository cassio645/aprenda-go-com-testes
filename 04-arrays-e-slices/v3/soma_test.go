package soma

import (
	"fmt"
	"reflect"
	"testing"
)

// Teste para função Soma normal, que recebe um slice.
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

// Teste para função SomaTudo que recebe vários slices.
func TestSomarTudo(t *testing.T) {
	resultado := SomarTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	// No Go não podemos comparar slices com == !=
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}

}

func ExampleSomarTudo(){
	resultado := SomarTudo([]int{10, 50}, []int{2, 4, 6})
	fmt.Println(resultado)
	// Output: [60 12]
}