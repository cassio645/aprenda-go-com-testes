// Pacote concorrencia é usado para testar se um website está ativo ou não usando goroutines e canais.
package concorrencia

// Tipo VerificadorWebsite é uma função que recebe uma string (URL) e retorna um bool indicando se o website está ativo.
type VerificadorWebsite func(string) bool

// O tipo resultado contém a URL e um booleano indicando se está ativo.
type resultado struct {
	URL   string
	Ativo bool
}

// Função recebe um verificador de websites e uma lista de URLs, utiliza goroutines e canais para verificar concorrentemente se os websites estão ativos ou não. Retorna um mapa onde a chave é a URL e o valor é um booleano indicando se o website está ativo.
func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	// Mapa para armazenar os resultados
	resultados := make(map[string]bool)

	// Canal para receber os resultados das goroutines
	canalResultado := make(chan resultado)

	for _, url := range urls {
		// Criando goroutine com uma função anônima para ser executada concorrentemente
		go func(u string) {
			// Enviando o resultado para o canal
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	// Loop para receber os resultados das goroutines
	for i := 0; i < len(urls); i++ {
		// Recebendo resultado do canal e atualizando o mapa
		resultado := <-canalResultado
		resultados[resultado.URL] = resultado.Ativo
	}

	// Fechando o canal, pois não precisamos mais dele
	close(canalResultado)

	return resultados
}
