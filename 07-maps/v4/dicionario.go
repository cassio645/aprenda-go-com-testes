// Pacote dicionario com as funções Busca, Adiciona, Atualiza e Deleta
package dicionario

// Mensagem de erro quando a palavra não for encontrada
const (
	ErrNaoEncontrado      = ErrDicionario("Palavra não encontrada.")
	ErrPalavraJaExistente = ErrDicionario("Esta palavra já existe.")
	ErrPalavraInexistente = ErrDicionario("Esta palavra não foi encontrada.")
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

// Função Adiciona uma nova chave: valor dentro do Dicionario
func (d Dicionario) Adiciona(chave, valor string) error {
	_, err := d.Busca(chave)

	// Caso Busca não encontrar a palavra é pq ela nao existe, então é adicionada
	// Caso retornar nil é pq achou a palavra, portante ela ja existe e não é adicionada
	switch err {
	case ErrNaoEncontrado:
		d[chave] = valor
	case nil:
		return ErrPalavraJaExistente
	}
	return nil
}

// Função Atualiza o valor de uma chave já existente
func (d Dicionario) Atualiza(chave, valor string) error {
	_, err := d.Busca(chave)

	// Caso Busca não encontrar a palavra é pq ela nao existe e não deveCasor adicionada com este método
	// Caso retornar nil é pq achou a palavra, portante ela ja existe então será atualizada
	switch err {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[chave] = valor
	}
	return nil
}

// Função Deleta uma chave:valor do dicionario
func (d Dicionario) Deleta(chave string){
	delete(d, chave)
}