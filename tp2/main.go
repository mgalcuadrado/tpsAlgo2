package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDARegistros "tp2/registros"
)

const (
	_MENSAJE_ERROR string = "Error en comando"
	_MENSAJE_OK    string = "OK"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)
	reg := TDARegistros.CrearRegistros()
	for entrada.Scan() { //devuelve false cuando no hay nada más que leer
		linea := entrada.Text()
		input := strings.Split(linea, " ")
		//Si la operación no es válida u ocurre un error al realizarse la misma imprime "Error en comando <comando>" por la salida de errores.
		//Si la operación se pudo realizar correctamente imprime "OK" por salida estándar.
		if !reg.RealizarOperacion(input) {
			fmt.Fprintf(os.Stderr, "%s %s\n", _MENSAJE_ERROR, input[0])
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", _MENSAJE_OK)
		}
	}
}
