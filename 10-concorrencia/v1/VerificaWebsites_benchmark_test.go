package concorrencia

import (
	"testing"
	"time"
)

// Função de simulação que espera 20 milissegundos antes de retornar true.
// Isso é usado no benchmark para simular um verificador de websites com latência.
func slowStubVerificadorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificaWebsites(b *testing.B) {
	// Cria um slice com 100 lugares
	urls := make([]string, 100)

	// Preenche todos lugares com a mesma string
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}

	for i := 0; i < b.N; i++ {
		VerificaWebsites(slowStubVerificadorWebsite, urls)
	}
}

// nano segundos: 3210887400
