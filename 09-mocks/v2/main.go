// Conta 3 2 1 Go! no intervalo de 1 segundo.
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Constantes com os valores iniciais e finais do que queremos printar
const (
	valorInicial = 3
	valorFinal   = "Go!"
)

// Função que recebe um buffer e printa 3 2 1 Go! com intervalo de 1s
func Contagem(saida io.Writer) {

	for i := valorInicial; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(saida, valorFinal)
}

func main() {
	Contagem(os.Stdout)
}
