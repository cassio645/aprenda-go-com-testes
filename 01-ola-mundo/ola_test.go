package main

import "testing"

// Função que testa Ola() e compara o resultado obtido com o valor esperado
func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá mundo!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestBomDia(t *testing.T){
	resultado := BomDia("Cássio")
	esperado := "Bom dia Cássio!"

	if resultado != esperado{
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

