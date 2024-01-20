package sync

import "testing"

// Função para verificar se o valor do contador é o que esperamos
func verificaContador(t *testing.T, contador Contador, esperado int) {
	t.Helper()

	if contador.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", contador.Valor(), esperado)
	}

}
func TestContador(t *testing.T) {

	t.Run("Contador incrementa 3x e o valor do contador deve ser 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementar()
		contador.Incrementar()
		contador.Incrementar()

		verificaContador(t, contador, 3)

	})

}
