// Pacote formas com a funções Perimetro e os Métodos Area contidos nos structs de Retangulo e Circulo.
package formas

import "math"

// Struct Retangulo
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Método Area do struct Retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}


// Struct Circulo
type Circulo struct {
	Raio float64
}

// Método Area do struct Circulo
func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}


// Perimetro recebe um struct do tipo Retangulo e retorna o valor do perimetro.
func Perimetro(r Retangulo) float64 {
	return 2 * (r.Altura + r.Largura)
}
