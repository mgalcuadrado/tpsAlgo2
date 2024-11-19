package cola

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_COLA_VACIA string = "La cola esta vacia"
)

/* Definición del struct cola proporcionado por la cátedra. */
type nodoCola[T any] struct {
	dato T
	sig  *nodoCola[T]
}

type colaEnlazada[T any] struct {
	prim *nodoCola[T]
	ult  *nodoCola[T]
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.prim == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_PANIC_COLA_VACIA)
	}
	return cola.prim.dato
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nodoNuevo := crearNodo[T](valor)
	if cola.EstaVacia() {
		cola.prim = nodoNuevo
	} else {
		cola.ult.sig = nodoNuevo
	}
	cola.ult = nodoNuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_PANIC_COLA_VACIA)
	}
	datoCola := cola.prim.dato
	if cola.prim == cola.ult {
		cola.ult = nil
	}
	cola.prim = cola.prim.sig
	return datoCola
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	//los dos nodos se inicializan automáticamente en nil por las características del lenguaje, manteniendo el invariante de representación
	return cola
}

/* ***************** FUNCIONES AUXILIARES  ***************** */
func crearNodo[T any](valor T) *nodoCola[T] {
	nodoNuevo := new(nodoCola[T])
	nodoNuevo.dato = valor
	return nodoNuevo
}
