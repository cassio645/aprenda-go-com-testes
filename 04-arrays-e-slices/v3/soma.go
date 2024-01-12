// Pacote com a função Soma e SomaTudo que somam os elementos contidos nos slices.
package soma

// Soma recebe um slice de inteiros e retorna soma de todos os elementos.
func Soma(numeros []int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}

// SomaTudo recebe vários slices e devolve a soma de cada um deles.
func SomarTudo(valores ...[]int) (total []int) {
	// Pega todos os slices contidos em valores e usa a função soma para calcular o resultado de cada slice passado.
	for _, valor := range valores {
		total = append(total, Soma(valor))
	}
	return
}
