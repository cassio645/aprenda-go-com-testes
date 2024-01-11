// Package iteração com a função de Repetir() uma string.
package iteracao

// Repetir recebe uma string e um int, e repete a string o número de vezes especificado.
func Repetir(s string, vezes int) (repetido string) {
	if vezes <= 0 {
		vezes = 5
	}

	for i := 0; i < vezes; i++ {
		repetido += s
	}
	return
}
