// Pacote dicionario com o tipo Dicionario map de strings com a função Busca
package dicionario

import (
	"errors"
)

// Mensagem de erro quando a palavra não for encontrada
var ErrNaoEncontrado = errors.New("Palavra não encontrada.")

// Tipo Dicionario é um map de strings
type Dicionario map[string]string

// Função Busca retorna um valor referente a chave, caso a chave nao existe retorna um erro.
func (d Dicionario) Busca(chave string) (string, error) {

	// Se a chave não for encontrada no dicionario existe recebe false
	valor, existe := d[chave]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return valor, nil
}
