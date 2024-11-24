package registros

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

const (
	_MENSAJE_ERROR                string        = "Error en comando"
	_CANTIDAD_LIMITE_ATAQUE_DOS   int           = 5
	_TIME_TO_LIVE                 time.Duration = 2 * time.Second //tiempo en el que se cuentan las requests de la página para analizar ataques DoS
	_CANTIDAD_CAMPOS_REGISTROS    int           = 4
	_CAPACIDAD_INICIAL_SITIOS     int           = 8
	_AGREGAR_ARCHIVO_COMANDO      string        = "agregar_archivo"
	_VER_VISITANTES_COMANDO       string        = "ver_visitantes"
	_VER_MAS_VISITADOS_COMANDO    string        = "ver_mas_visitados"
	_AGREGAR_ARCHIVO_PARAMETROS   int           = 2
	_VER_VISITANTES_PARAMETROS    int           = 3
	_VER_MAS_VISITADOS_PARAMETROS int           = 2
)

type registros struct {
	diccionarioIPs             TDADiccionario.DiccionarioOrdenado[IPv4, datosDiccionarioIPs]
	diccionarioSitiosVisitados TDADiccionario.Diccionario[string, int]
	registroActual             string
}

type datosDiccionarioIPs struct {
	ultimaVisita       string
	visitasDesdeTiempo int
	ataqueDoSReportado bool
	cola               TDACola.Cola[time.Time]
}

type sitioVisitado struct {
	sitio           string
	cantidadVisitas int
}

// AgregarArchivo recibe la ruta de un log y lo agrega al registro.
// Adicionalmente, imprime por salida estándar las IPs que realizaron ataques DoS en orden creciente.
// Devuelve un booleano indicando si la operación se pudo realizar correctamente.
func (reg *registros) AgregarArchivo(ruta string) bool {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()
	if archivo == nil {
		return false
	}
	reg.registroActual = ruta
	colaAtaquesDoS := reg.lecturaDeArchivo(archivo)
	if colaAtaquesDoS == nil {
		return false
	}
	for !colaAtaquesDoS.EstaVacia() {
		ip := colaAtaquesDoS.Desencolar()
		fmt.Fprintf(os.Stdout, "DoS: %s\n", IPDisplay(ip))
	}
	return true
}

// VerVisitantes imprime por salida estándar las IPs comprendidas entre desde y hasta que visitaron páginas marcadas en el registro.
// Devuelve un booleano indicando si la operación se pudo realizar correctamente.
func (reg *registros) VerVisitantes(desde IPv4, hasta IPv4) bool {
	fmt.Fprintf(os.Stdout, "Visitantes:\n")
	reg.diccionarioIPs.IterarRango(&desde, &hasta, func(ip IPv4, dato datosDiccionarioIPs) bool {
		fmt.Fprintf(os.Stdout, "\t%s\n", IPDisplay(ip))
		return true
	})
	return true
}

// VerMasVisitados imprime por salida estándar las n páginas más visitadas registradas en el registro.
// Devuelve un booleano indicando si la operación se pudo realizar correctamente.
func (reg *registros) VerMasVisitados(n int) bool {
	arrSitiosVisitados := make([]sitioVisitado, reg.diccionarioSitiosVisitados.Cantidad())
	contador := 0
	reg.diccionarioSitiosVisitados.Iterar(func(clave string, dato int) bool { //O(s) para s=cantidad de sitios visitados
		valor := sitioVisitado{
			sitio:           clave,
			cantidadVisitas: dato,
		}
		arrSitiosVisitados[contador] = valor
		contador++
		return true
	})
	colaSitiosVisitados := TDAColaPrioridad.CrearHeapArr(arrSitiosVisitados, compararSitiosVisitados) //es O(s) crear heap a partir de arreglo
	fmt.Fprintf(os.Stdout, "Sitios más visitados:\n")
	for i := 0; i < n && !colaSitiosVisitados.EstaVacia(); i++ { // O(n) * O(log s) = O(n log s)
		valor := colaSitiosVisitados.Desencolar() //O(log s)
		fmt.Fprintf(os.Stdout, "\t%s - %d\n", valor.sitio, valor.cantidadVisitas)
	}
	//Complejidad final 2 * O(s) +  O(n log s) --> O(s) +  O(n log s)
	return true
}

// CrearRegistros crea un registro para que se puedan realizar las operaciones de la interfaz
func CrearRegistros() *registros {
	reg := new(registros)
	reg.diccionarioIPs = TDADiccionario.CrearABB[IPv4, datosDiccionarioIPs](IPCompare)
	reg.diccionarioSitiosVisitados = TDADiccionario.CrearHash[string, int]()
	return reg
}

// RealizarOperacion realiza la operación solicitada.
// Devuelve un booleano indicando si la operación se pudo realizar correctamente.
func (reg *registros) RealizarOperacion(input []string) bool {
	switch input[0] {
	case _AGREGAR_ARCHIVO_COMANDO:
		if len(input) != _AGREGAR_ARCHIVO_PARAMETROS {
			return false
		}
		return reg.AgregarArchivo(input[1])
	case _VER_VISITANTES_COMANDO:
		if len(input) != _VER_VISITANTES_PARAMETROS {
			return false
		}
		return reg.VerVisitantes(IPParsear(input[1]), IPParsear(input[2]))
	case _VER_MAS_VISITADOS_COMANDO:
		if len(input) != _VER_MAS_VISITADOS_PARAMETROS {
			return false
		}
		n, _ := strconv.Atoi(input[1])
		return reg.VerMasVisitados(n)
	default:
		return false
	}
}

/* ********** FUNCIONES AUXILIARES ********** */

// abrirArchivo es una función interna que recibe una ruta y devuelve el archivo.
// Si ocurre un error al abrirlo devuelve nil.
func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}
	return archivo
}

// lecturaDeArchivos recibe un archivo y lo agrega a los registros.
// Adicionalmente, crea la cola de prioridad de IPs de las que se detectaron ataques DoS.
func (reg *registros) lecturaDeArchivo(archivo *os.File) TDAColaPrioridad.ColaPrioridad[IPv4] {
	entrada := bufio.NewScanner(archivo)
	colaAtaquesDoS := TDAColaPrioridad.CrearHeap[IPv4](IPCompareInverso) //se usa la función IPCompareInverso para que el heap de máximos funcione como un heap de mínimos.
	for entrada.Scan() {
		campos := strings.Split(entrada.Text(), "\t")
		if len(campos) != _CANTIDAD_CAMPOS_REGISTROS {
			return nil
		}
		reg.actualizarDiccionarioIPs(campos, colaAtaquesDoS)
		reg.actualizarSitiosVisitados(campos[3])
	}
	return colaAtaquesDoS
}

// actualizarABBIPs recibe los campos y la cola de prioridad.
// Guarda en los registros las IPs, detecta si se realizaron ataques DoS y, de ser el caso, actualiza la cola de prioridad.
func (reg *registros) actualizarDiccionarioIPs(campos []string, colaAtaquesDoS TDAColaPrioridad.ColaPrioridad[IPv4]) {
	ip := IPParsear(campos[0])
	tiempo, _ := time.Parse(time.RFC3339, campos[1])
	datos := crearDatosDiccionarioIP()
	if !reg.diccionarioIPs.Pertenece(ip) {
		resetearDatos(&datos, reg.registroActual, tiempo)
		reg.diccionarioIPs.Guardar(ip, *datos)
		return
	}
	*datos = reg.diccionarioIPs.Obtener(ip)
	if strings.Compare((*datos).ultimaVisita, reg.registroActual) != 0 {
		resetearDatos(&datos, reg.registroActual, tiempo)
	} else if (*datos).ataqueDoSReportado {
	} else {
		(*datos).cola.Encolar(tiempo)
		(*datos).visitasDesdeTiempo++
		for tiempo.Sub((*datos).cola.VerPrimero()) >= _TIME_TO_LIVE {
			(*datos).cola.Desencolar()
			(*datos).visitasDesdeTiempo--
		}
		if (*datos).visitasDesdeTiempo == _CANTIDAD_LIMITE_ATAQUE_DOS {
			colaAtaquesDoS.Encolar(ip)
			(*datos).ataqueDoSReportado = true
		}
	}
	reg.diccionarioIPs.Guardar(ip, *datos)
}

// actualizarSitiosVisitados recibe un sitio y le suma una visita al mismo en los registros.
func (reg *registros) actualizarSitiosVisitados(sitio string) {
	if reg.diccionarioSitiosVisitados.Pertenece(sitio) {
		cantidad := reg.diccionarioSitiosVisitados.Obtener(sitio)
		reg.diccionarioSitiosVisitados.Guardar(sitio, cantidad+1)
	} else {
		reg.diccionarioSitiosVisitados.Guardar(sitio, 1)
	}
}

// resetearDatos recibe un puntero a un puntero de datosDiccionarioIPs, el registro actual y el tiempo actual y lo resetea.
func resetearDatos(datos **datosDiccionarioIPs, log string, t time.Time) {
	(*datos).ultimaVisita = log
	(*datos).visitasDesdeTiempo = 1
	(*datos).ataqueDoSReportado = false
	(*datos).cola = TDACola.CrearColaEnlazada[time.Time]() //resetteo la cola
	(*datos).cola.Encolar(t)
}

// compararSitiosVisitados devuelve un número menor a cero si s1 < s2, 0 si s1=s2 y un número mayor a cero si s1>s2
func compararSitiosVisitados(s1 sitioVisitado, s2 sitioVisitado) int {
	return s1.cantidadVisitas - s2.cantidadVisitas
}

// crearDatosDiccionarioIP crea la estructura de datosDiccionarioIPs
func crearDatosDiccionarioIP() *datosDiccionarioIPs {
	var datos datosDiccionarioIPs
	datos.cola = TDACola.CrearColaEnlazada[time.Time]()
	return &datos
}
