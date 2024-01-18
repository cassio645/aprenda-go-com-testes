package corredor

import "testing"


// Teste usando sites reais(não recomendado para testes)
func TestCorredor(t *testing.T) {
	URLLenta := "http://www.facebook.com"
	URLRapida := "http://www.quii.co.uk"

	resultado := Corredor(URLLenta, URLRapida)
	esperado := URLRapida

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}
