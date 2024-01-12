// Pacote formas com as funções Perimetro e Area recebendo um Struct de Retangulo.
package formas

// Struct Retangulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Perimetro recebe um struct do tipo Retangulo e retorna o valor do perimetro.
func Perimetro(r Retangulo) float64 {
	return 2 * (r.Altura + r.Largura)
}

// Area recebe um struct do tipo Retangulo e retorna a área total do retangulo.
func Area(r Retangulo) float64 {
	return r.Altura * r.Largura
}
