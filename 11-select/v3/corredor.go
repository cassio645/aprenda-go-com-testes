// Pacote corredor com a função corredor() que compara qual url é mais rapida com limite de 10s.
package corredor

import (
	"fmt"
	"net/http"
	"time"
)

// Limite padrão para a função Corredor()
var limiteDeDezSegundos = 10 * time.Second

// Função compara o tempo de resposta de duas Urls e retorna a mais rápída, com tempo limite de 10s
func Corredor(a, b string) (vencedor string, erro error) {
    return Configuravel(a, b, limiteDeDezSegundos)
}

// Compara os tempos de resposta definindo um tempo limite (para testes)
func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(tempoLimite):
        return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
    }
}


// ping realiza uma requisição HTTP para a URL fornecida e envia um sinal para o canal quando a requisição é concluída.
func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
