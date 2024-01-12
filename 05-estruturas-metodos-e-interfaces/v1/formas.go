// Pacote formas com as funções Perimetro e Area
package formas

// Perimetro recebe dois float64 com altura e largura do retangulo e retorna o valor do perimetro.
func Perimetro(largura, altura float64) float64 {
	return 2 * (largura + altura)
}

// Area recebe dois valores float64 altura e largura e retorna a área total do retangulo.
func Area(largura, altura float64) float64 {
	return largura * altura
}
