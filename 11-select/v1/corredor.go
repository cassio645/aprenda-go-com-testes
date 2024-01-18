// Pacote corredor com a função corredor() que compara qual url é mais rapida.
package corredor

import (
	"net/http"
	"time"
)

// Função que recebe duas url e compara o tempo parece receber alguma resposta, retorna a url que demorou menos tempo.
func Corredor(a, b string) (vencedor string) {
	inicioA := time.Now()
	http.Get(a)
	duracaoA := time.Since(inicioA)

	inicioB := time.Now()
	http.Get(b)
	duracaoB := time.Since(inicioB)

	if duracaoA < duracaoB {
		return a
	}

	return b
}
