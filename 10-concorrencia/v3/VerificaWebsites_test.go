package concorrencia

import (
	"fmt"
	"reflect"
	"testing"
)

// Função de simulação que retorna false apenas para um URL específico.
func mockVerificadorWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	} else {
		return true
	}
}

func TestVerificaWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	esperado := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	resultado := VerificaWebsites(mockVerificadorWebsite, websites)

	if !reflect.DeepEqual(esperado, resultado) {
		t.Fatalf("esperado %v, resultado %v", esperado, resultado)
	}
}

// Função Example para usar VerificaWebsites
func ExampleVerificaWebsites() {
	// Função exemplo que você deve criar para verificar o status de um website
	verificadorSimples := func(url string) bool {
		return true
	}

	// Lista de URLs para verificar
	websites := []string{
		"http://example.com",
		"http://nonexistentwebsite.com",
	}

	// Chamando a função VerificaWebsites com o verificador simples
	resultados := VerificaWebsites(verificadorSimples, websites)

	// Exibindo os resultados
	for url, ativo := range resultados {
		fmt.Printf("Website: %s, Ativo: %t\n", url, ativo)
	}

	// Output:
	// Website: http://nonexistentwebsite.com, Ativo: true
	// Website: http://example.com, Ativo: true

}
