// Pacote com a função Soma recebendo um Slice
package soma

// Soma recebe um slice de inteiros e retorna soma de todos os elementos.
func Soma(numeros []int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}
