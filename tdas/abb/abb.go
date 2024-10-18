package diccionario

import (
	"fmt"
)

type nodoAB[K comparable, V any] struct {
    clave   K
    valor   V
    izquierda, derecha *nodoab[K, V]
}

type ab[K comparable, V any] struct {
    raiz       *nodoab[K, V]
    cmp        func(K, K) int //la funcion de comparacion es un atributo del ab
    cantidad   int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
    return &ab[K, V]{
        raiz:     nil,
        cmp:      funcion_cmp,
        cantidad: 0,
    }
}

/*La función de comparación, recibe dos claves y devuelve:

Un entero menor que 0 si la primera clave es menor que la segunda.
Un entero mayor que 0 si la primera clave es mayor que la segunda.
0 si ambas claves son iguales.*/

func (ab *ab[K, V]) Guardar(clave K, dato V) {
    ab.raiz = ab.insertarNodo(ab.raiz, clave, dato)
}

func (ab *ab[K, V]) insertarNodo(nodo *nodoAB[K, V], clave K, dato V) *nodoAB[K, V] {
    if nodo == nil {
        ab.cantidad++
        return &nodoAB[K, V]{clave: clave, valor: dato}
    }

    comparacion := ab.cmp(clave, nodo.clave)
    if comparacion < 0 {
        nodo.izquierda = ab.insertarNodo(nodo.izquierda, clave, dato) //si la clave a ingresar es mayor a la raiz, va a la izquierda
    } else if comparacion > 0 {
        nodo.derecha = ab.insertarNodo(nodo.derecha, clave, dato) //si la clave a ingresar es menor a la raiz, va a la derecha
    } else {
        nodo.valor = dato  // Si la clave ya existe, solo se actualiza el valor
    }
    return nodo
}
