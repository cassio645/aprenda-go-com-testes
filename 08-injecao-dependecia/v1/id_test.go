package main

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T){
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Cássio")

	resultado := buffer.String()
	esperado := "Olá, Cássio"

	if resultado != esperado{
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}