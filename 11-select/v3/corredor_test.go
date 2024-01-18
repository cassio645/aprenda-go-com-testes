package corredor

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Cria um servidor com um valor de atraso.
func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}

// Teste usando servidores criados a partir de net/http/httptest
func TestCorredor(t *testing.T) {

	t.Run("Testando url mais rápida", func(t *testing.T) {
		// Criando um servidor lento com atraso de 20milisegundos e outro sem atraso.
		servidorLento := criarServidorComAtraso(20 * time.Millisecond)
		servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		// defer é usado para fechar os servidores após tudo na função ser executado.
		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, err := Corredor(URLLenta, URLRapida)

		if err != nil {
			t.Fatalf("Não esperava erro, mas obteve %v", err)
		}

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Testando erro quando ambas demoram mais de 10s", func(t *testing.T) {
		servidor := criarServidorComAtraso(11 * time.Second)

		defer servidor.Close()

		// Definindo o tempo limite para 20 milisegundos(usado apenas para testar o tempo limite).
		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})

}


func ExampleCorredor() {
	vencedor, err := Corredor("http://google.com", "http://example.com")

	// Verifica se não há erro e exibe o vencedor.
	if err == nil {
		fmt.Printf("O vencedor é: %s\n", vencedor)
	} else {
		fmt.Printf("Erro: %v\n", err)
	}

	// Output:
	// O vencedor é: http://google.com
}