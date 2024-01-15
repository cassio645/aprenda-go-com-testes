package dicionario

import "testing"

func TestBusca(t *testing.T) {

	comparaString := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	dicionario := map[string]string{"teste": "isso é um teste"}

	resultado := Busca(dicionario, "teste")
	esperado := "isso é um teste"

	comparaString(t, resultado, esperado)

}
