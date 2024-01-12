package formas

import (
	"fmt"
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{Largura: 10.0, Altura: 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

func ExamplePerimetro() {
	retangulo := Retangulo{Largura: 10.0, Altura: 5.0}
	resultado := Perimetro(retangulo)
	fmt.Println(resultado)
	// Output: 30
}

func TestArea(t *testing.T) {

	validacao := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()

		// Chamando o método Area da interface Forma
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("Teste Area de um Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		esperado := 72.0
		validacao(t, retangulo, esperado)
	})

	t.Run("Teste Area de um Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		esperado := 314.1592653589793
		validacao(t, circulo, esperado)
	})

}

func ExampleArea(){
	retangulo := Retangulo{12.0, 6.0}
	circulo := Circulo{10}

	fmt.Println(retangulo.Area())
	fmt.Println(circulo.Area())
	// Output:
	// 72
	// 314.1592653589793
}