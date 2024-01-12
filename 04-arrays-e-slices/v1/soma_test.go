package soma

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {
	numeros := [5]int{2, 3, 4, 5, 6}

	resultado := Soma(numeros)
	esperado := 20

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d, dados os numeros %v", resultado, esperado, numeros)
	}
}

func ExampleSoma(){
	resultado := Soma([5]int{10, 20, 30, 40})
	fmt.Println(resultado)
	// Output: 100
}