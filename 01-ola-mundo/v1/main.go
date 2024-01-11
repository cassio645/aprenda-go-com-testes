// Pacote do primeiro capítulo do curso.
package ola

import "fmt"

// Função que retorna a string 'Olá mundo!'
func Ola() string {
	return "Olá mundo!"
}

// Função que retorna uma string 'Bom dia [NOME]!'
func BomDia(s string) string {
	return "Bom dia " + s + "!"
}

// Função main chamando a função Ola()
func main() {
	fmt.Println(Ola())
}
