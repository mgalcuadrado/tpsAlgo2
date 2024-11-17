package registros

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

const (
	_MENSAJE_ERROR                string        = "Error en comando" //revisar: este no es el texto correcto
	timeToLive                    time.Duration = 2000000000         //tiempo (en ns) en el que se cuentan los pedidos para analizar ataque DoS
	_CAPACIDAD_INICIAL_SITIOS     int           = 8
	_AGREGAR_ARCHIVO_COMANDO      string        = "agregar_archivo"
	_VER_VISITANTES_COMANDO       string        = "ver_visitantes"
	_VER_MAS_VISITADOS_COMANDO    string        = "ver_mas_visitados"
	_AGREGAR_ARCHIVO_PARAMETROS   int           = 2
	_VER_VISITANTES_PARAMETROS    int           = 3
	_VER_MAS_VISITADOS_PARAMETROS int           = 2
	_CANTIDAD_CAMPOS_REGISTROS    int           = 4
	_CANTIDAD_LIMITE_ATAQUE_DOS   int           = 5
)

type datos_diccionario struct {
	ultimaVisita       string
	tiempo             time.Time
	visitasDesdeTiempo int
	ataqueDoSReportado bool
}

type sitioVisitado struct {
	sitio           string
	cantidadVisitas int
}

type registros struct {
	funcionesDisponibles TDADiccionario.Diccionario[string, int]
	abbIPs               TDADiccionario.DiccionarioOrdenado[IPv4, datos_diccionario]
	hashSitiosVisitados  TDADiccionario.Diccionario[string, int]
	registroActual       string
	//guardamos un hash de sitios visitados (por nombre del sitio), lo iteramos para guardarlo en un arreglo y le hacemos HeapSort y los vamos sacando en ese orden
	//inicialmente había pensado directo en una cola de prioridad, pero me cuesta encontrar los elementos para actualizarlos
}

func (reg *registros) AgregarArchivo(ruta string) bool {
	archivo := abrirArchivo(ruta)
	if archivo == nil {
		cerrarArchivo(archivo)
		return false
	}
	reg.registroActual = ruta
	error, heap := reg.lecturaDeArchivo(archivo)
	if error != nil {
		cerrarArchivo(archivo)
		return false
	}
	for !heap.EstaVacia() {
		ip := heap.Desencolar()
		fmt.Fprintf(os.Stdout, "DoS: %d.%d.%d.%d\n", ip.partes[0], ip.partes[1], ip.partes[2], ip.partes[3])
	}
	cerrarArchivo(archivo)
	return true
}

func (reg *registros) VerVisitantes(desde IPv4, hasta IPv4) bool {
	fmt.Fprintf(os.Stdout, "Visitantes:\n")
	reg.abbIPs.IterarRango(&desde, &hasta, func(ip IPv4, dato datos_diccionario) bool {
		//if strings.Compare(dato.ultimaVisita, reg.registroActual) == 0 {
		fmt.Fprintf(os.Stdout, "\t%d.%d.%d.%d\n", ip.partes[0], ip.partes[1], ip.partes[2], ip.partes[3])
		//}
		return true
	})
	return true
}

func (reg *registros) VerMasVisitados(n int) bool {
	heap := TDAColaPrioridad.CrearHeap(compararSitiosVisitados)
	reg.hashSitiosVisitados.Iterar(func(clave string, dato int) bool {
		valor := sitioVisitado{
			sitio:           clave,
			cantidadVisitas: dato,
		}
		heap.Encolar(valor)
		return true
	})
	fmt.Fprintf(os.Stdout, "Sitios más visitados:\n")
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		valor := heap.Desencolar()
		fmt.Fprintf(os.Stdout, "\t%s - %d\n", valor.sitio, valor.cantidadVisitas)
	}
	return true
}

func CrearRegistros() *registros {
	reg := new(registros)
	reg.abbIPs = TDADiccionario.CrearABB[IPv4, datos_diccionario](IPCompare)
	reg.hashSitiosVisitados = TDADiccionario.CrearHash[string, int]()
	reg.funcionesDisponibles = TDADiccionario.CrearHash[string, int]()
	//clave = entradas del usuario, dato = cantidad de parámetros recibidos por linea de comandos (contando la función)
	reg.funcionesDisponibles.Guardar("agregar_archivo", _AGREGAR_ARCHIVO_PARAMETROS)
	reg.funcionesDisponibles.Guardar("ver_visitantes", _VER_VISITANTES_PARAMETROS)
	reg.funcionesDisponibles.Guardar("ver_mas_visitados", _VER_MAS_VISITADOS_PARAMETROS)
	return reg
}

func (reg *registros) Operar(input []string) bool {
	if !reg.funcionesDisponibles.Pertenece(input[0]) || reg.funcionesDisponibles.Obtener(input[0]) != len(input) {
		return false //revisar: no me acuerdo de si hay que seguir
	}
	//revisar: tiene que haber una mejor forma de hacer esto pero... no la estoy viendo
	if strings.Compare(input[0], "agregar_archivo") == 0 {
		return reg.AgregarArchivo(input[1])
	}
	if strings.Compare(input[0], "ver_visitantes") == 0 {
		return reg.VerVisitantes(IPParsear(input[1]), IPParsear(input[2]))
	}
	if strings.Compare(input[0], "ver_mas_visitados") == 0 {
		n, _ := strconv.Atoi(input[1])
		return reg.VerMasVisitados(n)
	}
	return true
}

/* ********** FUNCIONES AUXILIARES ********** */
func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}
	return archivo
}

func cerrarArchivo(archivo *os.File) error {
	return archivo.Close()
}

func (reg *registros) lecturaDeArchivo(archivo *os.File) (error, TDAColaPrioridad.ColaPrioridad[IPv4]) {
	entrada := bufio.NewScanner(archivo)
	heap := TDAColaPrioridad.CrearHeap[IPv4](IPCompareInverso)
	for entrada.Scan() {
		campos := strings.Split(entrada.Text(), "\t")
		if len(campos) != _CANTIDAD_CAMPOS_REGISTROS {
			return errors.New("Error"), nil
		}
		reg.actualizarABBIPs(campos, heap)
		reg.actualizarSitiosVisitados(campos[3])
	}
	return nil, heap
}

func (reg *registros) actualizarABBIPs(campos []string, heap TDAColaPrioridad.ColaPrioridad[IPv4]) {
	ip := IPParsear(campos[0])
	tiempo, _ := time.Parse(time.DateTime, campos[1])
	datos := new(datos_diccionario)
	if !reg.abbIPs.Pertenece(ip) {
		resetearDatos(&datos, reg.registroActual, tiempo)
	} else {
		*datos = reg.abbIPs.Obtener(ip)
		if strings.Compare((*datos).ultimaVisita, reg.registroActual) != 0 {
			resetearDatos(&datos, reg.registroActual, tiempo)
		} else if !(*datos).ataqueDoSReportado && tiempo.Sub((*datos).tiempo) < timeToLive {
			(*datos).visitasDesdeTiempo++
			if (*datos).visitasDesdeTiempo >= _CANTIDAD_LIMITE_ATAQUE_DOS {
				heap.Encolar(ip)
				(*datos).ataqueDoSReportado = true
			}
		}
	}
	reg.abbIPs.Guardar(ip, *datos)
}

func (reg *registros) actualizarSitiosVisitados(sitio string) {
	if reg.hashSitiosVisitados.Pertenece(sitio) {
		cantidad := reg.hashSitiosVisitados.Obtener(sitio)
		reg.hashSitiosVisitados.Guardar(sitio, cantidad+1)
	} else {
		//sitioVisitado := memcopy(sitio) //revisar: me gustaría duplicar la cadena pero no me acuerdo de cómo JAJAJAAJ
		reg.hashSitiosVisitados.Guardar(sitio, 1)
	}
}

func resetearDatos(datos **datos_diccionario, log string, t time.Time) {
	(*datos).ultimaVisita = log
	(*datos).tiempo = t
	(*datos).visitasDesdeTiempo = 1
	(*datos).ataqueDoSReportado = false
}

// compararSitiosVisitados devuelve un número menor a cero si s1 < s2, 0 si s1=s2 y un número mayor a cero si s1>s2
func compararSitiosVisitados(s1 sitioVisitado, s2 sitioVisitado) int {
	return s1.cantidadVisitas - s2.cantidadVisitas
}
