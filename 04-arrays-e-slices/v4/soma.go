// Pacote com a função Soma e SomaTodoOResto que somam os elementos contidos nos slices.
package soma

// Soma recebe um slice de inteiros e retorna soma de todos os elementos.
func Soma(numeros []int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}

// SomaTodoOResto recebe vários slices e retorna o valor da soma de cada um deles EXCETO do primeiro elemento de cada slice.
func SomaTodoOResto(valores ...[]int) (total []int) {
	// Pega todos os slices contidos em valores;
	// Usa a função soma para calcular o resultado de cada slice passado a partir do segundo elemento do slice.
	for _, valor := range valores {
		// Se o slice for vazio, adiciona 0
		if len(valor) == 0 {
			total = append(total, 0)
		} else {
			total = append(total, Soma(valor[1:]))
		}
	}
	return
}
