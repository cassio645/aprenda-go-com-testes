// Pacote com a função Ola() usando uma função para o prefixo
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
	return prefixoSaudacao(idioma) + nome

}

// Função que verifica o idioma recebido e retorna o prefixo correspondente.
// Começa com letra minuscula pois é uma função privada.
func prefixoSaudacao(idioma string) (prefixo string) {

	// Verifica se o idioma é espanhol ou frances, caso contrario responde em português
	switch idioma {
	case espanhol:
		prefixo = prefixoEs
	case frances:
		prefixo = prefixoFr
	default:
		prefixo = prefixoPt
	}
	return
}

func main() {

}
