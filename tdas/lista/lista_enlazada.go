package lista

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_LISTA_VACIA string = "La lista esta vacia"
)

/* **************** DEFINICIÓN DEL STRUCT NODO PROPORCIONADO POR LA CÁTEDRA **************** */
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

/* **************** DEFINICIÓN DEL STRUCT LISTA PROPORCIONADO POR LA CÁTEDRA **************** */
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

/* ****************  IMPLEMENTACIÓN DE LA LISTA **************** */

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])

	//si bien el lenguaje por sí solo declara esto de esta manera, explicito mi invariante de representación
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0

	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nodoNuevo := new(nodoLista[T])
	nodoNuevo.dato = valor
	nodoNuevo.siguiente = lista.primero
	if lista.primero == nil {
		lista.ultimo = nodoNuevo
	}
	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nodoNuevo := new(nodoLista[T])
	nodoNuevo.dato = valor
	if lista.largo == 0 {
		lista.primero = nodoNuevo
	} else if lista.largo == 1 {
		lista.primero.siguiente = nodoNuevo
	} else {
		lista.ultimo.siguiente = nodoNuevo
	}
	lista.ultimo = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	prim := lista.primero
	lista.primero = lista.primero.siguiente
	lista.largo--
	return prim.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

/* ****************  IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */
//@anto acá no entiendo qué estás haciendo
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero

	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

/* **************** DEFINICIÓN DE LA STRUCT ITERADOR EXTERNO **************** */
type iteradorLista[T any] struct { //justamente el iterador lista no tiene una interfaz que el usuario tenga que conocer
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iteradorLista[T])
	iterador.actual = lista.primero
	iterador.anterior = nil //igual esto lo define el lenguaje directamente
	iterador.lista = lista
	return iterador
	//return &IteradorLista[T]{actual: lista.primero, anterior: nil, lista: lista}
}

func (it *iteradorLista[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iteradorLista[T]) VerActual() T {
	if it.actual == nil {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	return it.actual.dato
}

func (it *iteradorLista[T]) Siguiente() {
	if it.lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	it.anterior = it.actual //hay que mantener la invariante de representación (y por ende el anterior actualizado)
	it.actual = it.actual.siguiente
}

func (it *iteradorLista[T]) Insertar(dato T) {
	if it.anterior == nil { //si el anterior es nil estoy insertando al principio, reuso primitiva de la pila
		it.lista.InsertarPrimero(dato) //acá ya se suma a largo directamente
		it.actual = it.lista.primero
	} else {
		nuevoNodo :=new(nodoLista[T])
		nuevoNodo.dato = dato
		nuevoNodo.siguiente = it.actual
		it.anterior.siguiente = nuevoNodo
		it.actual = nuevoNodo
		it.lista.largo++ //invariante de representación
	}
	if it.actual.siguiente == nil {
		it.lista.ultimo = it.actual
	}
	
}

func (it *iteradorLista[T]) Borrar() T {
	if it.lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
	borrado := it.actual.dato
	if it.actual == it.lista.primero { //acá anterior vale nil y el siguiente al actual puede o no ser nil
		it.lista.primero = it.actual.siguiente
	} else if it.actual == it.lista.ultimo {
		it.lista.ultimo = it.anterior
	} else {
		it.anterior.siguiente = it.actual.siguiente
	}
	it.actual = it.actual.siguiente //para el primero eso es el que ahora es el primero, para el último eso es nil, para los demás casos va bien
	it.lista.largo-- //invariante de representación
	return borrado
}

