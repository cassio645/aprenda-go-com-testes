package main

import (
	"fmt"
	"io"
	"os"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func main() {
    Cumprimenta(os.Stdout, "Elodie")
}