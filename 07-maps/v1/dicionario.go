// Pacote dicionario com a função Busca, busca um valor no map atraves de uma chave.
package dicionario

// Busca recebe um map de strings e uma chave, e retorna o valor referente àquela chave
func Busca(dicionario map[string]string, chave string) string {
	return dicionario[chave]
}
