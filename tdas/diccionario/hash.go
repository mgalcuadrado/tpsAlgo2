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
	vacio estados = iota //a mirar
	ocupado
	borrado
)

type celda struct {
	clave comparable
	valor any
	estado estados
}

type hashCerrado[K comparable, V any] struct{
	arreglo *celdas
	largo int
	cantElem int 
	cantBorrados int 
}

type iterHashCerrado[K comparable, V any] struct {
	hash hashCerrado
	indice int 
}

/* **************** IMPLEMENTACIÓN DEL HASH **************** */
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.celdas := make([]celda, _CANTIDAD_INICIAL)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {

}

func (hash *hashCerrado[K, V]) Borrar(clave K) V

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.CantElem
}

func (hash *hashCerrado[K, V]) redimensionarHash() {

}

// funcionDeHashing recibe la clave y devuelve un entero 
func (hash *hashCerrado[K, V]) funcionDeHashing(clave comparable) int {
	return hashSHA256(clave) % hash.largo
}
//función proveída por el paquete SHA256 de golang: https://pkg.go.dev/crypto/sha256
func hashSHA256(key string) string {
    h := sha256.New()
    h.Write([]byte(key))
    return hex.EncodeToString(h.Sum(nil))
}

// buscar recibe la clave y la busca en el arreglo de celdas. Devuelve la posición en la que esa clave se encuentra en el arreglo. Si no está devuelve -1
func func (hash *hashCerrado[K, V]) buscar(clave comparable) int {

}
 
/* **************** IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool)


/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

}
func (*iterHashCerrado[K, V])HaySiguiente() bool {

}

func (*iterHashCerrado[K, V])VerActual() (K, V) {

}
func (*iterHashCerrado[K, V]) Siguiente() {

}