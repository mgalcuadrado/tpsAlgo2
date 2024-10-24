package pila

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_CAPACIDAD_INICIAL_PILA                   int    = 16
	_CONSTANTE_DE_REDIMENSION_PILA            int    = 2
	_CONSTANTE_DE_REDUCCION_DE_CAPACIDAD_PILA int    = 4
	_MENSAJE_PANIC_PILA_VACIA                 string = "La pila esta vacia"
)

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(_MENSAJE_PANIC_PILA_VACIA)
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(valor T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * _CONSTANTE_DE_REDIMENSION_PILA)
	}
	pila.datos[pila.cantidad] = valor
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(_MENSAJE_PANIC_PILA_VACIA)
	}
	if pila.cantidad > _CAPACIDAD_INICIAL_PILA && pila.cantidad == cap(pila.datos)/_CONSTANTE_DE_REDUCCION_DE_CAPACIDAD_PILA {
		pila.redimensionar(cap(pila.datos) / _CONSTANTE_DE_REDIMENSION_PILA)
	}
	pila.cantidad--
	return pila.datos[pila.cantidad]
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL_PILA)
	pila.cantidad = 0 // invariante de representación
	//Si bien esta variable se inicializa en cero por las características del lenguaje, se explicita por claridad
	return pila
}

/* ******************** FUNCIONES AUXILIARES ******************** */

// redimensionar es una función auxiliar que recibe un valor de capacidad para el slice a crear y luego lo reemplaza por el slice de datos de la pila
func (pila *pilaDinamica[T]) redimensionar(capacidad int) {
	datosPilaAuxiliar := make([]T, capacidad)
	copy(datosPilaAuxiliar, pila.datos)
	pila.datos = datosPilaAuxiliar
}
