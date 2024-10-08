package diccionario

import (
	"fmt"
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO string = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_FIN_DE_ITERACION                 string = "El iterador termino de iterar"

	_FACTOR_DE_CARGA_MAXIMO                  float64 = 0.7 //esto lo podemos modificar
	_FACTOR_DE_CARGA_MINIMO                  float64 = 0.2
	_CANTIDAD_INICIAL                        int     = 53
	_MULTIPLICADOR_DE_INCREMENTO_DE_CANTIDAD int     = 2
	_MULTIPLICADOR_DE_DECREMENTO_DE_CANTIDAD float64 = 2 / 3
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
type estados int

const (
	_VACIO estados = iota //a mirar
	_OCUPADO
	_BORRADO
)

type celda[K comparable, V any] struct {
	clave  K
	valor  V
	estado estados
}

type hashCerrado[K comparable, V any] struct {
	celdas       []celda[K, V]
	largo        int
	cantElem     int
	cantBorrados int
}

type IterHashCerrado[K comparable, V any] struct {
	hash          *hashCerrado[K, V]
	indice_actual int
}

/* **************** IMPLEMENTACIÓN DEL HASH **************** */

/* **************** FUNCIONES DE LA INTERFACE **************** */
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return crearHash[K, V](_CANTIDAD_INICIAL)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	factorDeCarga := float64(hash.cantElem+hash.cantBorrados) / float64(hash.largo)
	if (factorDeCarga) >= _FACTOR_DE_CARGA_MAXIMO {
		hash.redimensionarHash(int(hash.largo * _MULTIPLICADOR_DE_INCREMENTO_DE_CANTIDAD))
	}
	pertenece, posicion := hash.buscar(clave)
	if pertenece {
		hash.celdas[posicion].valor = dato
	} else {
		hash.celdas[posicion].estado = _OCUPADO
		hash.celdas[posicion].clave = clave
		hash.celdas[posicion].valor = dato
		hash.cantElem++ //mantengo invariante de representación
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	pertenece, _ := hash.buscar(clave)
	return pertenece
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.verificarPosicion(clave) //si llego al return es porque no saltó el panic
	return hash.celdas[posicion].valor
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	//if float64(hash.cantElem + hash.cantBorrados) /float64(hash.largo) <= _FACTOR_DE_CARGA_MINIMO {
	//hash.redimensionarHash(int(float64(hash.largo) * _MULTIPLICADOR_DE_DECREMENTO_DE_CANTIDAD))
	//}
	posicion := hash.verificarPosicion(clave)
	hash.cantBorrados++
	hash.cantElem--
	hash.celdas[posicion].estado = _BORRADO
	return hash.celdas[posicion].valor
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantElem
}

/* **************** Funciones Auxiliares **************** */

// crearHash es una función auxiliar que recibe el largo del que se quiere que sea la tabla del hash y devuelve un hashCerrado
func crearHash[K comparable, V any](largo int) *hashCerrado[K, V] {
	hash := new(hashCerrado[K, V])
	hash.celdas = make([]celda[K, V], largo)
	hash.largo = largo
	return hash
}

// verificarPosicion recibe la clave y verifica si la posición es válida o no. Se usa al obtener y borrar, pues si no pertenece devuelve panic
func (hash *hashCerrado[K, V]) verificarPosicion(clave K) int {
	if hash.Cantidad() == 0 {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
	}
	pertenece, posicion := hash.buscar(clave)
	if !pertenece {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
	}
	return posicion
}

// redimensionarHash recibe el nuevoLargo de la tabla de hashing y reemplaza el hash por el nuevo
func (hash *hashCerrado[K, V]) redimensionarHash(nuevoLargo int) {
	hashNuevo := crearHash[K, V](nuevoLargo)
	for i := 0; i < hash.largo; i++ {
		if hash.celdas[i].estado == _OCUPADO {
			hashNuevo.Guardar(hash.celdas[i].clave, hash.celdas[i].valor)
		}
	}
	//acá hash.cantElem == hashNuevo.cantElem siempre
	hash.celdas = hashNuevo.celdas
	hash.largo = hashNuevo.largo
	hash.cantBorrados = hashNuevo.cantBorrados
}

// buscar recibe la clave y la busca en el arreglo de celdas. Devuelve un booleano indicando si la clave se halló en el arreglo y la posición en la que esa clave se encuentra en el arreglo.
func (hash *hashCerrado[K, V]) buscar(clave K) (bool, int) {
	posicion := funcionDeHashing(clave, hash.largo)
	indice := posicion                          //una vez que haya redimensión sacamos esto pero mientras lo dejo para que no se rompa nada
	if hash.celdas[posicion].estado == _VACIO { //si no hay elementos en el hash o la celda en la que debería estar esa clave está vacía, entonces no está la clave
		return false, posicion
	}
	for hash.celdas[indice].estado != _VACIO {

		if hash.celdas[indice].estado == _OCUPADO && hash.celdas[indice].clave == clave {
			return true, indice
		}
		if indice == len(hash.celdas)-1 {
			indice = 0
		} else {
			indice++
		}
		if indice == posicion { //una vez que haya redimensión esto nunca debería pasar pero mientras lo dejo para que no se rompa nada
			return false, posicion
		}
	}
	return false, indice
}

/* **************** Función de hashing **************** */

// funcionDeHashing recibe la clave y devuelve un entero
func funcionDeHashing[K comparable](clave K, largo int) int {
	claveEnBytes := convertirABytes[K](clave)
	return int(fnv1aHash(claveEnBytes)) % largo
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fnv1aHash(data []byte) uint32 {
	const FNV_prime uint32 = 0x1000193
	const offset_basis uint32 = 0x811C9DC5
	// Iniciar el hash con el offset basis
	hash := offset_basis
	// Iterar sobre cada byte en el array de bytes
	for _, b := range data {
		// XOR el byte actual con el hash
		hash ^= uint32(b)
		// Multiplicar por el prime (FNV prime)
		hash *= FNV_prime
	}
	return hash
}

/* **************** IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */

/* **************** IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < hash.largo; i++ {
		if hash.celdas[i].estado == _OCUPADO {
			if !visitar(hash.celdas[i].clave, hash.celdas[i].valor) {
				return
			}
		}
	}
}

/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

func (hash *hashCerrado[K, V]) Iterador() *IterHashCerrado[K, V] {
	iterador := new(IterHashCerrado[K, V])
	iterador.hash = hash
	iterador.indice_actual = 0
	return iterador
}

func (iterador *IterHashCerrado[K, V]) HaySiguiente() bool {
	return iterador.hash.celdas[iterador.indice_actual].estado == _OCUPADO
}

func (iterador *IterHashCerrado[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic(_MENSAJE_PANIC_FIN_DE_ITERACION)
	}
	return iterador.hash.celdas[iterador.indice_actual].clave, iterador.hash.celdas[iterador.indice_actual].valor
}

func (iterador *IterHashCerrado[K, V]) Siguiente() {
	iterador.indice_actual++

	for iterador.indice_actual != iterador.hash.largo-1 {
		if iterador.hash.celdas[iterador.indice_actual].estado == _OCUPADO {
			return
		}
	}

	panic(_MENSAJE_PANIC_FIN_DE_ITERACION)
}
