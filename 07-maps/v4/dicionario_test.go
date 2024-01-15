package dicionario

import (
	"fmt"
	"testing"
)

// Função usada para comparar o resultado da busca em um Dicionario
func comparaDefinicao(t *testing.T, dicionario Dicionario, chave, valor string) {
	t.Helper()

	resultado, err := dicionario.Busca(chave)

	if err != nil {
		t.Fatal("Erro.", err)
	}

	if resultado != valor {
		t.Errorf("resultado %s, esperado %s", resultado, valor)
	}
}

// Função usada para comparar a mensagem de erro de alguns casos dos testes
func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

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
	t.Run("Adicionando chave:valor nova", func(t *testing.T) {
		// Criando dicionario vazio e adicionando uma nova chave: valor
		dicionario := Dicionario{}
		chave := "amora"
		valor := "Fruta"
		err := dicionario.Adiciona(chave, valor)

		comparaErro(t, err, nil) // não deve dar erro
		comparaDefinicao(t, dicionario, chave, valor)

	})

	// O dicionario esta sendo criado já com os valores, e a função adiciona tenta adicionar na mesma chave.
	t.Run("Adicionando chave:valor já existente", func(t *testing.T) {

		chave := "amora"
		valor := "Fruta"
		dicionario := Dicionario{chave: valor}

		err := dicionario.Adiciona(chave, "Outro valor")
		comparaErro(t, err, ErrPalavraJaExistente) // deve dar erro
		comparaDefinicao(t, dicionario, chave, valor)
	})

}

func ExampleDicionario_Adiciona() {
	dicionario := Dicionario{}
	err := dicionario.Adiciona("cachorro", "Animal")

	if err != nil {
		fmt.Println(err)
	}

	resultado, _ := dicionario.Busca("cachorro")

	fmt.Println(resultado)
	// Output: Animal
}

func TestAtualiza(t *testing.T) {
	t.Run("Atualizando uma chave que já existe", func(t *testing.T) {
		chave := "bateria"
		valor := "objeto"
		novoValor := "instrumento"

		dicionario := Dicionario{chave: valor} // dicionario criado com chave:valor
		err := dicionario.Atualiza(chave, novoValor)

		comparaErro(t, err, nil) // não deve dar erro
	})

	t.Run("Atualizando uma chave que não existe", func(t *testing.T) {
		chave := "bateria"
		valor := "objeto"

		dicionario := Dicionario{} // dicionario criado vazio
		err := dicionario.Atualiza(chave, valor)

		comparaErro(t, err, ErrPalavraInexistente) // não deve dar erro

	})
}

func ExampleDicionario_Atualiza() {
	dicionario := Dicionario{"bateria": "Objeto"}

	err := dicionario.Atualiza("bateria", "Instrumento")

	if err != nil {
		fmt.Println(err)
	}

	resultado, _ := dicionario.Busca("bateria")
	fmt.Println(resultado)
	// Output: Instrumento
}

func TestDeleta(t *testing.T) {
	chave := "cadeira"
	dicionario := Dicionario{chave: "Objeto"}

	dicionario.Deleta(chave)

	_, err := dicionario.Busca(chave)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", chave)
	}
}

func ExampleDicionario_Deleta() {
	chave := "cadeira"
	dicionario := Dicionario{chave: "Objeto"}
	dicionario.Deleta("jornal")

	_, err := dicionario.Busca(chave)
	if err != nil {
		fmt.Println(err)
	}
	// Output: Palavra não encontrada.
}
