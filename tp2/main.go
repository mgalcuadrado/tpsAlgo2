package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
	"strings"
	//TDADiccionario "tdas/diccionario"
	TDARegistros "tp2/registros"
)

const (
	_AGREGAR_ARCHIVO_COMANDO      string = "agregar_archivo"
	_VER_VISITANTES_COMANDO       string = "ver_visitantes"
	_VER_MAS_VISITADOS_COMANDO    string = "ver_max_visitados"
	_AGREGAR_ARCHIVO_PARAMETROS   int    = 2
	_VER_VISITANTES_PARAMETROS    int    = 3
	_VER_MAS_VISITADOS_PARAMETROS int    = 2
	_MENSAJE_ERROR                string = "Error en comando"
	_MENSAJE_OK                   string = "OK"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)
	reg := TDARegistros.CrearRegistros()
	for entrada.Scan() { //devuelve false cuando no hay nada más que leer
		linea := entrada.Text()
		input := strings.Split(linea, " ")
		if !reg.Operar(input) {
			fmt.Fprintf(os.Stdout, "%s %s\n", _MENSAJE_ERROR, input[0])
		} else {
			fmt.Fprintf(os.Stdout, _MENSAJE_OK)
		}
	}
}
