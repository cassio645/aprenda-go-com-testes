package main

import "testing"

// Usando t.Run para implementar subtestes
func TestOla(t *testing.T) {

	// Usa t.Helper() para fazer a verificação do resultado com o esperado(evita reescrita repetida de codigo)
	verificaMsgCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	// Testando Ola(arg) com argumento
	t.Run("Diz 'Olá [Nome]' quando recebe argumento", func(t *testing.T) {
		resultado := Ola("Cássio")
		esperado := "Olá Cássio"

		
		verificaMsgCorreta(t, resultado, esperado)
		// Substituido pelo uso do t.Helper()
		//if resultado != esperado {t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)}
	})

	// Testando Ola("") sem argumento
	t.Run("Diz 'Olá mundo' quando não recebe argumento", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá mundo!"

		verificaMsgCorreta(t, resultado, esperado)
	})
}
