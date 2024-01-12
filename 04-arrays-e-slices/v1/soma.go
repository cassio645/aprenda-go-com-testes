// Pacote com a função Soma recebendo arrays
package soma

// Soma recebe um array de inteiros e retorna soma de todos os elementos.
func Soma(numeros [5]int) (total int) {
	for _, valor := range numeros {
		total += valor
	}
	return
}
