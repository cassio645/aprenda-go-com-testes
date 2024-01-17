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


// Interface Sleeper te permite definir as pausas
type Sleeper interface {
	Pausa()
}

// SleeperPadrao é uma implementação de Sleeper com o tempo de pausa definido
type SleeperPadrao struct{}

// Pausa vai pausar a execução pela Duração definida
func (d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}



// Função que recebe um buffer e printa 3 2 1 Go! com intervalo de 1s
func Contagem(saida io.Writer, sleeper Sleeper) {

	for i := valorInicial; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, valorFinal)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
