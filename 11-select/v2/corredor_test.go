package corredor

import (
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

	// Criando um servidor lento com atraso de 20milisegundos e outro sem atraso.
	servidorLento := criarServidorComAtraso(20 * time.Millisecond)
	servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

	// defer é usado para fechar os servidores após tudo na função ser executado.
	defer servidorLento.Close()
	defer servidorRapido.Close()

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
