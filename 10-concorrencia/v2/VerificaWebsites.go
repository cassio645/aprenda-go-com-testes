// Pacote usado para testar se um website esta ativo ou não usando goroutine
package concorrencia

import "time"

// Tipo que verifica um website recebendo uma string e retorna um bool
type VerificadorWebsite func(string) bool

// Função VerificaWebsites recebe um tipo verificador e um slice de urls usando goroutines, retorna um map com a chave=url e valor=true ou false
func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		// Criando goroutine com uma função anonima para ser executada concorrentemente
		go func(u string){
		resultados[u] = vw(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return resultados
}
