// Package iteração com a função de Repetir() uma string.
package iteracao

// Numero de vezes que a iteração será feita
const numeroVezes = 5

// Repetir recebe uma string e retorna ela repetida 5x
func Repetir(s string) (repetido string) {
	for i := 0; i < numeroVezes; i++ {
		repetido += s
	}
	return
}
