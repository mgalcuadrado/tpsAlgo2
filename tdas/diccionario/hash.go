package diccionario

import (
    "crypto/sha256"
    "encoding/hex"
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO      string = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_FIN_DE_ITERACION string = "El iterador termino de iterar"
	
	_FACTOR_DE_CARGA float64 = 0.7 //esto lo podemos modificar
	_CANTIDAD_INICIAL int = 53
	_MULTIPLICADOR_DE_INCREMENTO_DE_CANTIDAD int = 2
	_MULTIPLICADOR_DE_DECREMENTO_DE_CANTIDAD int = 3
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
type estados int

const (
	_VACIO estados = iota //a mirar
	_OCUPADO
	_BORRADO
)

type celda[K comparable, V any] struct {
	clave K
	valor V
	estado estados
}

type hashCerrado[K comparable, V any] struct{
	celdas []celda[K, V]
	largo int
	cantElem int 
	cantBorrados int 
}

type iterHashCerrado[K comparable, V any] struct {
	hash hashCerrado[K, V]
	indice int 
}

/* **************** IMPLEMENTACIÓN DEL HASH **************** */
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.celdas = make([]celda[K, V], _CANTIDAD_INICIAL)
	return hash
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	//este if es para zafar hasta que haya redimensión
	if hash.cantElem + hash.cantBorrados == hash.largo {
		return
	}
	//acá vamos a tener el if para redimensionar!
	posicion := hash.buscar(clave)
	if posicion != -1 {
		hash.celdas[posicion].valor = dato
	} else {
		hash.celdas[posicion].estado = _OCUPADO
		hash.celdas[posicion].clave = clave
		hash.celdas[posicion].valor = dato
		hash.cantElem++
	}
}

func (hash * hashCerrado[K,V]) hallarPosicionDisponible(clave K) int {
	posicion := hash.buscar(clave)
	for ; hash.celdas[posicion].estado != _VACIO; {
		if posicion == len(hash.celdas) {
			posicion = 0
		} else {
			posicion++
		}
	}
	return posicion
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return hash.buscar(clave) != -1
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.verificarPosicion(clave) //si llego al return es porque no saltó el panic
	return hash.celdas[posicion].valor
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	//acá vamos a tener el if para redimensionar
	posicion := hash.verificarPosicion(clave)
	hash.cantBorrados++
	hash.cantElem--
	hash.celdas[posicion].estado = _BORRADO
	return hash.celdas[posicion].valor
}

func (hash *hashCerrado[K, V]) verificarPosicion(clave K) int {
	posicion := hash.buscar(clave)
	if posicion == -1 {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
	}
	return posicion
}


func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantElem
}

func (hash *hashCerrado[K, V]) redimensionarHash() {

}

// funcionDeHashing recibe la clave y devuelve un entero 
func (hash *hashCerrado[K, V]) funcionDeHashing(clave comparable) int {
	return hashSHA256(clave) % hash.largo
}

//función proveída por el paquete SHA256 de golang: https://pkg.go.dev/crypto/sha256
func hashSHA256(key string) string { //esta función recibe y devuelve string, así que me complica un poco las operaciones
    h := sha256.New()
    h.Write([]byte(key))
    return hex.EncodeToString(h.Sum(nil))
}

// buscar recibe la clave y la busca en el arreglo de celdas. Devuelve la posición en la que esa clave se encuentra en el arreglo. Si no está devuelve -1
func (hash *hashCerrado[K, V]) buscar(clave comparable) int {
	posicion := hash.funcionDeHashing (clave)
	indice := posicion //una vez que haya redimensión sacamos esto pero mientras lo dejo para que no se rompa nada
	if hash.cantElem == 0 || hash.celdas[posicion].estado == _VACIO { //si no hay elementos en el hash o la celda en la que debería estar esa clave está vacía, entonces no está la clave
		return -1
	}
	for hash.celdas[indice].estado != _VACIO {
		if indice == len(hash.celdas) {
			indice = 0
		} else {
			indice++
		}
		if indice == posicion{ //una vez que haya redimensión esto nunca debería pasar pero mientras lo dejo para que no se rompa nada
			return -1
		}
		if hash.celdas[indice].estado == _OCUPADO && hash.celdas[indice].clave == clave {
			return indice
		}
	}
	return -1
	 

}
 
/* **************** IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {

}


/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */
/*
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

}

func (*iterHashCerrado[K, V])HaySiguiente() bool {
	return true
}

func (*iterHashCerrado[K, V])VerActual() (K, V) {

}
func (*iterHashCerrado[K, V]) Siguiente() {

}
*/