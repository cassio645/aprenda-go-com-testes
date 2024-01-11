// Pacote com a função Ola() recebendo idioma.
package ola

// Criando constantes de idiomas
const espanhol = "es"
const frances = "fr"

// Criando constantes com prefixo 'Olá" em alguns idiomas
const prefixoPt = "Olá "
const prefixoEs = "Hola "
const prefixoFr = "Bonjour "

// Função que retorna a string 'Olá [NOME]' no idioma informado
func Ola(nome, idioma string) string {

	// Se o nome for vazio, nome recebe 'mundo!'
	if nome == "" {
		nome = "mundo!"
	}

	var prefixo string

	// Verifica se o idioma é espanhol ou frances, caso contrario responde em português
	switch idioma {
	case espanhol:
		prefixo = prefixoEs
	case frances:
		prefixo = prefixoFr
	default:
		prefixo = prefixoPt
	}
	return prefixo + nome

}

func main() {

}
