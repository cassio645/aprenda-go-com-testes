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


// Teste para função SomaTodoOResto que recebe vários slices.
func TestSomaTodoOResto(t *testing.T) {
	validacao := func(t *testing.T, resultado, esperado []int) {
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	}

	t.Run("Testando com vários slices", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		validacao(t, resultado, esperado)
	})

	t.Run("Testando com slice vazio", func(t *testing.T){
		resultado := SomaTodoOResto([]int{}, []int{1, 2, 3})
		esperado := []int{0, 5}

		validacao(t, resultado, esperado)
	})

}

func ExampleSomaTodoOResto(){
	resultado := SomaTodoOResto([]int{1, 2, 3}, []int{10, 4, 2})
	fmt.Println(resultado)
	// Output: [5 6]
}