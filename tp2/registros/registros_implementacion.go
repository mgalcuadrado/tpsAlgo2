package registros

import (
	"fmt"
	//TDAColaPrioridad "tdas/cola_prioridad"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
	//"strconv"
)

const (
	_MENSAJE_ERROR            string = "error" //revisar: este no es el texto correcto
	timeToLive                int    = 2       //tiempo en el que se cuentan los pedidos para analizar ataque DoS
	_CAPACIDAD_INICIAL_SITIOS int    = 8
)

type log string

type datos_diccionario struct {
	ultimaVisita       log
	tiempo             time.Time
	visitasDesdeTiempo int
}

type sitiosVisitados struct {
	sitio            string
	cantidad_visitas int
}

type registros struct {
	dic                    TDADiccionario.DiccionarioOrdenado[IPv4, datos_diccionario]
	arregloSitiosVisitados []sitiosVisitados
	//guardamos un arreglo ordenado de sitios visitados (por nombre del sitio) y le hacemos heapify (por cantidad de visitas) para ver_mas_visitados (O(n)) y después los vamos sacando en orden!
	//inicialmente había pensado directo en una cola de prioridad, pero me cuesta encontrar los elementos para actualizarlos
}

func (reg *registros) AgregarArchivo(ruta string) {
	cadena := strings.Split(ruta, ".")
	fmt.Printf(cadena[0])
}

func (reg *registros) VerVisitantes(desde IPv4, hasta IPv4) {
	fmt.Printf("%d", desde)
}

func (reg *registros) VerMasVisitados(n int) {
	fmt.Printf("%d", n)
}

func CrearRegistros() *registros {
	reg := new(registros)
	reg.dic = TDADiccionario.CrearABB[IPv4, datos_diccionario](IPCompare)
	reg.arregloSitiosVisitados = make([]sitiosVisitados, _CAPACIDAD_INICIAL_SITIOS)
	return reg
}
