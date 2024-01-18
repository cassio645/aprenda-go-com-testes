// Pacote usado para testar se um website esta ativo ou não
package concorrencia

// Tipo que verifica um website recebendo uma string e retorna um bool
type VerificadorWebsite func(string) bool

// Função VerificaWebsites recebe um tipo verificador e um slice de urls, retorna um map com a chave=url e valor=true ou false
func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		resultados[url] = vw(url)
	}

	return resultados
}
