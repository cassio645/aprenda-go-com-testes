// Pacote do primeiro capítulo do curso.
package main

// Criando uma constante com prefixo em português
const prefixoPtBr = "Olá "

// Função que retorna a string 'Olá [NOME]'
func Ola(nome string) string {

	// Se o nome for vazio, nome recebe 'mundo!'
	if nome == "" {
		nome = "mundo!"
	}
	return prefixoPtBr + nome
}

// Função main
func main() {

}
