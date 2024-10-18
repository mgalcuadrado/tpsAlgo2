package lista

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_LISTA_VACIA      string = "La lista esta vacia"
	_MENSAJE_PANIC_FIN_DE_ITERACION string = "El iterador termino de iterar"
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
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func crearNodo[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: siguiente}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nodoNuevo := crearNodo[T](valor, lista.primero)

	if lista.EstaVacia() {
		lista.ultimo = nodoNuevo
	}
	lista.primero = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nodoNuevo := crearNodo[T](valor, nil)

	if lista.EstaVacia() {
		lista.primero = nodoNuevo
	} else {
		lista.ultimo.siguiente = nodoNuevo
	}
	lista.ultimo = nodoNuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verificarListaNoVacia()
	prim := lista.primero
	lista.primero = lista.primero.siguiente
	lista.largo--
	return prim.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verificarListaNoVacia()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verificarListaNoVacia()
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) verificarListaNoVacia() {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANIC_LISTA_VACIA)
	}
}

/* ****************  IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil && visitar(actual.dato); actual = actual.siguiente {
	}
}

/* **************** DEFINICIÓN DE LA STRUCT ITERADOR EXTERNO **************** */
type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iteradorLista[T])
	iterador.actual = lista.primero
	iterador.anterior = nil //si bien el lenguaje lo declara de esta manera, se explicitan las invariantes de representación
	iterador.lista = lista
	return iterador
}

func (it *iteradorLista[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iteradorLista[T]) VerActual() T {
	it.verificarHaySiguiente()
	return it.actual.dato
}

func (it *iteradorLista[T]) Siguiente() {
	it.verificarHaySiguiente()
	it.anterior = it.actual
	it.actual = it.actual.siguiente
}

func (it *iteradorLista[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato, nil)

	if it.anterior == nil {
		nuevoNodo.siguiente = it.lista.primero
		it.lista.primero = nuevoNodo
	} else {
		nuevoNodo.siguiente = it.actual
		it.anterior.siguiente = nuevoNodo
	}

	it.actual = nuevoNodo

	if it.actual.siguiente == nil {
		it.lista.ultimo = it.actual
	}

	it.lista.largo++

}

func (it *iteradorLista[T]) Borrar() T {
	it.verificarHaySiguiente()
	borrado := it.actual.dato
	if it.actual == it.lista.primero { //acá anterior vale nil y el siguiente al actual puede o no ser nil
		it.lista.primero = it.actual.siguiente
	} else if it.actual == it.lista.ultimo {
		it.lista.ultimo = it.anterior
	} else {
		it.anterior.siguiente = it.actual.siguiente
	}
	it.actual = it.actual.siguiente //para el primero eso es el que ahora es el primero, para el último eso es nil, para los demás casos va bien
	it.lista.largo--                //invariante de representación
	return borrado
}

func (it *iteradorLista[T]) verificarHaySiguiente() {
	if !it.HaySiguiente() {
		panic(_MENSAJE_PANIC_FIN_DE_ITERACION)
	}
}
