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
	retangulo := Retangulo{12.0, 6.0}
	resultado := Area(retangulo)
	esperado := 72.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}

}

func ExampleArea() {
	retangulo := Retangulo{12.0, 6.0}
	resultado := Area(retangulo)
	fmt.Println(resultado)
	// Output: 72
}
