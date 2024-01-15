// Pacote dicionario com as funções Busca e Adiciona num map de strings
package dicionario


// Mensagem de erro quando a palavra não for encontrada
const (
	ErrNaoEncontrado      = ErrDicionario("Palavra não encontrada.")
	ErrPalavraJaExistente = ErrDicionario("Esta palavra já existe.")
)

// Criando um tipo erro para o dicionario
type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

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

// Função Adiciona coloca uma nova chave: valor dentro do Dicionario
func (d Dicionario) Adiciona(chave, valor string) error {
	_, err := d.Busca(chave)

	// Se Busca não encontrar a palavra é pq ela nao existe
	// Se retornar nil é pq achou a palavra, portante ela ja existe
	switch err {
	case ErrNaoEncontrado:
		d[chave] = valor
	case nil:
		return ErrPalavraJaExistente
	}
	return nil
}
