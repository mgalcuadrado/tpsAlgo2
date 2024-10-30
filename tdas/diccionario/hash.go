package diccionario

import (
	"fmt"
)

/* **************** DEFINICIÓN DE CONSTANTES **************** */
const (
	_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO string = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_FIN_DE_ITERACION                 string = "El iterador termino de iterar"

	_FACTOR_DE_CARGA_MAXIMO                  float32 = 0.68
	_FACTOR_DE_CARGA_MINIMO                  float32 = 0.25
	_CANTIDAD_INICIAL                        int     = 157
	_MULTIPLICADOR_DE_INCREMENTO_DE_CANTIDAD int     = 2
	_DIVISOR_DE_DECREMENTO_DE_CANTIDAD       int     = 2
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
type estado int

const (
	_VACIO estado = iota
	_OCUPADO
	_BORRADO
)

type celda[K comparable, V any] struct {
	clave  K
	valor  V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	celdas          []celda[K, V]
	capacidadCeldas int
	cantElem        int
	cantBorrados    int
}

type iterHashCerrado[K comparable, V any] struct {
	hash          *hashCerrado[K, V]
	indice_actual int
}

/* **************** IMPLEMENTACIÓN DEL HASH **************** */

/* **************** FUNCIONES DE LA INTERFACE **************** */

// CrearHash crea un diccionario vacío
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.celdas = crearTabla[K, V](_CANTIDAD_INICIAL)
	hash.capacidadCeldas = _CANTIDAD_INICIAL
	return hash
}

// Guardar guarda una clave y su dato asociado en el diccionario.
func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if (hash.factorDeCarga()) >= _FACTOR_DE_CARGA_MAXIMO {
		hash.redimensionarHash(hash.capacidadCeldas * _MULTIPLICADOR_DE_INCREMENTO_DE_CANTIDAD)
	}
	pertenece, posicion := hash.buscar(clave)
	if !pertenece {
		hash.celdas[posicion].estado = _OCUPADO
		hash.celdas[posicion].clave = clave
		hash.cantElem++ //mantengo invariante de representación
	}
	hash.celdas[posicion].valor = dato
}

// Pertenece indica si una clave se encuentra en el diccionario
func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	pertenece, _ := hash.buscar(clave)
	return pertenece
}

// Obtener devuelve el valor asociado a una clave
func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.obtenerPosicionValida(clave) //si llego al return es porque no saltó el panic
	return hash.celdas[posicion].valor
}

// Borrar borra una clave del diccionario y devuelve su valor asociado
func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if (hash.capacidadCeldas > _CANTIDAD_INICIAL) && hash.factorDeCarga() < _FACTOR_DE_CARGA_MINIMO {
		hash.redimensionarHash(hash.capacidadCeldas / _DIVISOR_DE_DECREMENTO_DE_CANTIDAD)
	}
	posicion := hash.obtenerPosicionValida(clave)
	hash.cantBorrados++
	hash.cantElem--
	hash.celdas[posicion].estado = _BORRADO
	return hash.celdas[posicion].valor
}

// Cantidad indica la cantidad de elementos guardados en el diccionario
func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantElem
}

/* **************** Funciones Auxiliares **************** */

// crearHash es una función auxiliar que recibe el largo del que se quiere que sea la tabla del hash y devuelve un hashCerrado
func crearTabla[K comparable, V any](largo int) []celda[K, V] {
	tabla := make([]celda[K, V], largo)
	return tabla
}

// obtenerPosicionValida recibe la clave y verifica si la posición es válida o no. Si la clave no pertenece al diccionario ocurre un panic y, si está, devuelve la posición en la que se encuentra la clave.
func (hash *hashCerrado[K, V]) obtenerPosicionValida(clave K) int {
	if hash.Cantidad() == 0 {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
	}
	pertenece, posicion := hash.buscar(clave)
	if !pertenece {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
	}
	return posicion
}

func (hash *hashCerrado[K, V]) factorDeCarga() float32 {
	return float32(hash.cantElem+hash.cantBorrados) / float32(hash.capacidadCeldas)
}

// redimensionarHash recibe el nuevoLargo de la tabla de hashing y reemplaza el hash por el nuevo
func (hash *hashCerrado[K, V]) redimensionarHash(nuevoLargo int) {
	tablaAnterior, capacidadCeldasAnterior := hash.celdas, hash.capacidadCeldas
	hash.celdas = crearTabla[K, V](nuevoLargo)
	hash.capacidadCeldas = nuevoLargo
	hash.cantElem = 0
	hash.cantBorrados = 0
	for i := 0; i < capacidadCeldasAnterior; i++ {
		if tablaAnterior[i].estado == _OCUPADO {
			hash.Guardar(tablaAnterior[i].clave, tablaAnterior[i].valor)
		}
	}
	//hash.cantElem == hashNuevo.cantElem siempre
}

// buscar recibe la clave y la busca en el arreglo de celdas. Devuelve un booleano indicando si la clave se halló en el arreglo y la posición en la que esa clave se encuentra en el arreglo.
func (hash *hashCerrado[K, V]) buscar(clave K) (bool, int) {
	indice := funcionDeHashing[K](clave, hash.capacidadCeldas)
	for hash.celdas[indice].estado != _VACIO {
		if hash.celdas[indice].estado == _OCUPADO && hash.celdas[indice].clave == clave {
			return true, indice
		}
		indice = (indice + 1) % hash.capacidadCeldas
	}
	return false, indice
}

/* **************** Función de hashing **************** */

func funcionDeHashing[K comparable](clave K, modulo int) int {
	claveEnBytes := convertirABytes[K](clave)
	return int(fnv1aHash(claveEnBytes) % uint32(modulo))
}

// convertirABytes convierte una clave de tipo de dato comparable a un arreglo de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// obtuve esta función de ChatGPT, es una implementación en Go de una función de hashing existente llamada FNV1A. Dejo el enlace a la conversación: https://chatgpt.com/share/6708b012-bc00-8011-bf09-7e32b7f7e70e
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

// Iterar recibe una función visitar e itera hasta que la función dé false
func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < hash.capacidadCeldas; i++ {
		if hash.celdas[i].estado == _OCUPADO && !visitar(hash.celdas[i].clave, hash.celdas[i].valor) {
			return
		}
	}
}

/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

// Iterador crea un iterador externo del hash
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iterHashCerrado[K, V])
	iterador.hash = hash
	iterador.avanzarAlSiguienteOcupado()
	iterador.avanzarAlSiguienteOcupado()
	return iterador
}

// HaySiguiente informa si hay un elemento a continuación que se pueda visitar
func (iterador *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iterador.indice_actual != iterador.hash.capacidadCeldas
}

// VerActual permite obtener la clave y el valor en la que se encuentra el iterador. Si se terminó la iteración da panic
func (iterador *iterHashCerrado[K, V]) VerActual() (K, V) {
	iterador.verificarSiguiente()
	return iterador.hash.celdas[iterador.indice_actual].clave, iterador.hash.celdas[iterador.indice_actual].valor
}

// Siguiente avanza el iterador a la siguiente posición ocupada del hash. Si se terminó la iteración da panic
func (iterador *iterHashCerrado[K, V]) Siguiente() {
	iterador.verificarSiguiente()
	iterador.indice_actual++
	iterador.avanzarAlSiguienteOcupado()
	iterador.indice_actual++
	iterador.avanzarAlSiguienteOcupado()
}

// verificarSiguiente() corrobora si hay un siguiente; si no lo hay, da un panic.
func (iterador *iterHashCerrado[K, V]) verificarSiguiente() {
	if !iterador.HaySiguiente() {
		panic(_MENSAJE_PANIC_FIN_DE_ITERACION)
	}
}

// avanzarAlSiguienteOcupado() es una función interna que devuelve el índice en el que está el siguiente elemento válido.
func (iterador *iterHashCerrado[K, V]) avanzarAlSiguienteOcupado() {
	for (iterador.indice_actual < iterador.hash.capacidadCeldas) && iterador.hash.celdas[iterador.indice_actual].estado != _OCUPADO {
		iterador.indice_actual++
	}
}
