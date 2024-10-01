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
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero

	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

/* **************** DEFINICIÓN DE LA LINTERFAZ ITERADOR LISTA (EXTERNO)**************** */
type IteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (lista *listaEnlazada[T]) Iterador() *IteradorLista[T] {
	return &IteradorLista[T]{actual: lista.primero}
}

func (it *IteradorLista[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *IteradorLista[T]) VerActual() T {
	return it.actual.dato
}

func (it *IteradorLista[T]) Siguiente() {
	it.actual = it.actual.siguiente
}

func (it *IteradorLista[T]) Insertar(dato T) {
	nuevoNodo := &nodoLista[T]{dato: dato}

	if it.anterior == nil {
		nuevoNodo.siguiente = it.lista.primero
		it.lista.primero = nuevoNodo
	} else {
		nuevoNodo.siguiente = it.actual
		it.anterior.siguiente = nuevoNodo
	}

	if it.anterior != nil {
		it.anterior = nuevoNodo
	}

}

func (it *IteradorLista[T]) Borrar(dato T) {
	if it.actual == nil {
		return
	}

	if it.actual == it.lista.primero {
		it.lista.primero = it.actual.siguiente
	} else {
		it.anterior.siguiente = it.actual.siguiente
	}

	it.actual = it.actual.siguiente

}
