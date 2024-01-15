package dicionario

import (
	"fmt"
	"testing"
)

// Teste da função Busca
func TestBusca(t *testing.T) {

	// Criando o dicionario para os testes
	dicionario := Dicionario{"teste": "isso é um teste"}

	comparaString := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	comparaErro := func(t *testing.T, resultado, esperado error) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %s, erro esperado %s", resultado, esperado)
		}
	}

	t.Run("Busca palavra contida no dicionario", func(t *testing.T) {

		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é um teste"

		comparaString(t, resultado, esperado)

	})

	t.Run("Busca palavra desconhecida no dicionario", func(t *testing.T) {
		_, resultado := dicionario.Busca("cenoura")

		comparaErro(t, resultado, ErrNaoEncontrado)

	})

}

// Example de como usar a função Busca
func ExampleDicionario_Busca() {
	dicionario := Dicionario{"amora": "Fruta", "banco": "Objeto", "cachorro": "Animal"}

	resultado, err := dicionario.Busca("cachorro")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultado)
	// Output: Animal
}

// Teste da função Adiciona
func TestAdiciona(t *testing.T) {

	comparaDefinicao := func(t *testing.T, dicionario Dicionario, chave, valor string) {
		t.Helper()

		resultado, err := dicionario.Busca(chave)

		if err != nil {
			t.Fatal("Erro.", err)
		}

		if resultado != valor {
			t.Errorf("resultado %s, esperado %s", resultado, valor)
		}
	}

	comparaErro := func(t *testing.T, resultado, esperado error) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Adicionando palavra nova", func(t *testing.T) {
		// Criando dicionario vazio e adicionando uma nova chave: valor
		dicionario := Dicionario{}
		chave := "amora"
		valor := "Fruta"
		err := dicionario.Adiciona(chave, valor)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, chave, valor)

	})

	// O dicionario esta sendo criado já com os valores, e a função adiciona tenta adicionar na mesma chave.
	t.Run("Adicionando palavra já existente", func(t *testing.T) {

		chave := "amora"
		valor := "Fruta"
		dicionario := Dicionario{chave: valor}

		err := dicionario.Adiciona(chave, "Outro valor")
		comparaErro(t, err, ErrPalavraJaExistente)
		comparaDefinicao(t, dicionario, chave, valor)
	})

}

func ExampleDicionario_Adiciona() {
	dicionario := Dicionario{}
	err := dicionario.Adiciona("cachorro", "Animal")

	if err != nil {
		fmt.Println(err)
	}

	resultado, err := dicionario.Busca("cachorro")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultado)
	// Output: Animal
}
