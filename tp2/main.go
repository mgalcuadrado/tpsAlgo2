package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	TDARegistros "tp2/registros"
)

const (
	_AGREGAR_ARCHIVO_PARAMETROS   int    = 2
	_VER_VISITANTES_PARAMETROS    int    = 3
	_VER_MAS_VISITADOS_PARAMETROS int    = 2
	_MENSAJE_ERROR                string = "Error en comando"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)
	funcionesDisponibles := TDADiccionario.CrearHash[string, int]()
	//clave = entradas del usuario, dato = cantidad de parámetros recibidos por linea de comandos (contando la función)
	funcionesDisponibles.Guardar("agregar_archivo", _AGREGAR_ARCHIVO_PARAMETROS)
	funcionesDisponibles.Guardar("ver_visitantes", _VER_VISITANTES_PARAMETROS)
	funcionesDisponibles.Guardar("ver_mas_visitados", _VER_MAS_VISITADOS_PARAMETROS)
	reg := TDARegistros.CrearRegistros()
	for entrada.Scan() { //devuelve false cuando no hay nada más que leer
		linea := entrada.Text()
		input := strings.Split(linea, " ")
		if !funcionesDisponibles.Pertenece(input[0]) || funcionesDisponibles.Obtener(input[0]) != len(input) {
			fmt.Printf("%s %s\n", _MENSAJE_ERROR, input[0])
			break //revisar: no me acuerdo de si hay que seguir
		}
		//revisar: tiene que haber una mejor forma de hacer esto pero... no la estoy viendo
		if strings.Compare(input[0], "agregar_archivo") == 0 {
			reg.AgregarArchivo(input[1])
		}
		if strings.Compare(input[0], "ver_visitantes") == 0 {
			reg.VerVisitantes(TDARegistros.IPParsear(input[1]), TDARegistros.IPParsear(input[2]))
		}
		if strings.Compare(input[0], "ver_mas_visitados") == 0 {
			n, _ := strconv.Atoi(input[1])
			reg.VerMasVisitados(n)
		}
	}
}

//reg := TDARegistros.CrearRegistros()
//ip := TDARegistros.IPParsear("128.23.2.4")
//reg.AgregarArchivo("pruebas_analog/test01.log")
//reg.VerVisitantes(ip, ip)
//reg.VerMasVisitados(4)
