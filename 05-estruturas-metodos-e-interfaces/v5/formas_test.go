package formas

import (
	"fmt"
	"reflect"
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

// Testando com table driven tests, uma lista de testes
func TestForma(t *testing.T) {
	// Criando uma lista de structs com a interface Forma
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retangulo", forma: Retangulo{Largura: 12.0, Altura: 6.0}, esperado: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10.0}, esperado: 314.1592653589793},
		{nome: "Retangulo", forma: Retangulo{Largura: 10.0, Altura: 8.0}, esperado: 80.0},
		{nome: "Triangulo", forma: Triangulo{Base: 12.0, Altura: 6.0}, esperado: 36.0},
	}

	// For percorrendo cada elemento da lista de structs
	for _, valor := range testesArea {
		// Usando t.Run() para uma saída mais limpa em caso de falha.
		t.Run(valor.nome, func(t *testing.T) {
			resultado := valor.forma.Area()
			if resultado != valor.esperado {
				t.Errorf("%#v resultado %.2f, esperado %.2f", valor.forma, resultado, valor.esperado)
			}
		})

	}
}

func ExampleForma() {
	// Criando uma lista de structs com a interface Forma
	formas := []Forma{
		Retangulo{Largura: 12.0, Altura: 6.0},
		Circulo{Raio: 10.0},
		Triangulo{Base: 12.0, Altura: 6.0},
	}

	// Iterando sobre as formas e imprimindo o resultado da função Area
	for _, forma := range formas {
		fmt.Printf("%s - Área: %.2f\n", reflect.TypeOf(forma).Name(), forma.Area())
	}

	// Output:
	// Retangulo - Área: 72.00
	// Circulo - Área: 314.16
	// Triangulo - Área: 36.00
}
