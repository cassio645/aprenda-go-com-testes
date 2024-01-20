package sincroniza

import (
	"fmt"
	"sync"
	"testing"
)

// Função para verificar se o valor do contador é o que esperamos
func verificaContador(t *testing.T, contador *Contador, esperado int) {
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

		verificaContador(t, &contador, 3)

	})

	t.Run("Contador usado de maneira concorrente", func(t *testing.T) {
		contagemEsperada := 1000
		contador := Contador{}

		// Criando o waitgroup que vai esperar todas goroutines
		var wg sync.WaitGroup
		wg.Add(contagemEsperada) // Informando a quantidade de goroutines

		for i := 0; i < contagemEsperada; i++ {
			// Goroutine incrementando o contador de maneira concorrente
			go func(w *sync.WaitGroup) {
				contador.Incrementar()
				w.Done() //  Avisando que a goroutine já foi executada
			}(&wg)
		}
		wg.Wait() // Esperando todas goroutines finalizarem

		verificaContador(t, &contador, contagemEsperada)
	})

}


func ExampleContador_Incrementar() {
	contador := NovoContador()

	// Chama a função Incrementar várias vezes
	for i := 0; i < 5; i++ {
		contador.Incrementar()
	}
	fmt.Println(contador.Valor())
	// Output: 5
}

func ExampleContador_Valor() {
	contador := NovoContador()

	// Incrementa o Contador algumas vezes
	contador.Incrementar()
	contador.Incrementar()

	// Obtém e exibe o valor atual do Contador
	valor := contador.Valor()
	fmt.Println(valor)
	// Output: 2
}


