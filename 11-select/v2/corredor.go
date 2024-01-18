// Pacote corredor com a função corredor() que compara qual url é mais rapida.
package corredor

import (
	"net/http"
	"time"
)

// Função que recebe duas url e retorna a url que demorou menos tempo para responder.
func Corredor(a, b string) (vencedor string) {
	duracaoA := medirTempoDeResposta(a)
	duracaoB := medirTempoDeResposta(b)

	if duracaoA < duracaoB {
		return a
	}
	return b
}

// Recebe uma url e retorna o tempo de resposta dela.
func medirTempoDeResposta(url string) time.Duration {
	inicio := time.Now()
	http.Get(url)
	duracao := time.Since(inicio)
	return duracao
}
