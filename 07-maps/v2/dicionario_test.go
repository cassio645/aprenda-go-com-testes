package dicionario

import (
	"fmt"
	"testing"
)

func TestBusca(t *testing.T) {

	// Criando o dicionario para os testes
	dicionario := Dicionario{"teste": "isso é um teste"}

	comparaString := func(t *testing.T, resultado, esperado string) {
		t.Helper()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	comparaErro := func(t *testing.T, resultado, esperado error){
		t.Helper()

		if resultado != esperado{
			t.Errorf("resultado %s, erro esperado %s", resultado, esperado)
		}
	}

	t.Run("Busca palavra contida no dicionario", func(t *testing.T){

		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é um teste"
	
		comparaString(t, resultado, esperado)

	})

	t.Run("Busca palavra desconhecida no dicionario", func(t *testing.T){
		_, resultado := dicionario.Busca("cenoura")

		comparaErro(t, resultado, ErrNaoEncontrado)

	})

}

func ExampleDicionario_Busca(){
	dicionario := Dicionario{"amora": "Fruta", "banco": "Objeto", "cachorro": "Animal"}

	resultado, err := dicionario.Busca("cachorro")
	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println(resultado)
	// Output: Animal
}