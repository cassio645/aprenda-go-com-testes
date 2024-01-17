// Conta 3 2 1 Vai! no intervalo de 1 segundo.
package main

import (
	"fmt"
	"io"
	"os"
)

// Função que recebe um buffer e printa 3
func Contagem(saida io.Writer){
	fmt.Fprintf(saida, "3")
}

func main(){
	Contagem(os.Stdout)
}