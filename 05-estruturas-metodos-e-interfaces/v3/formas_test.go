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

	validacao := func(t *testing.T, resultado, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}

	}

	// Testando o método Area do retangulo
	t.Run("Teste Area de um Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0

		validacao(t, resultado, esperado)
	})

	// Testando o método Area do circulo
	t.Run("Teste Area de um Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := 314.1592653589793

		validacao(t, resultado, esperado)
	})

}

func ExampleArea() {
	retangulo := Retangulo{12.0, 6.0}
	circulo := Circulo{10}

	fmt.Println(retangulo.Area())
	fmt.Println(circulo.Area())
	// Output:
	// 72
	// 314.1592653589793
}
