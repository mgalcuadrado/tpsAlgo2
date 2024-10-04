package diccionario


/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO      string = "La clave no pertenece al diccionario"
	_FACTOR_DE_CARGA float64 = 1.0 //esto lo vamos a modificar en función de qué hash hagamos
	_MENSAJE_PANIC_FIN_DE_ITERACION string = "El iterador termino de iterar"
)



/* **************** IMPLEMENTACIÓN DEL HASH **************** */
func CrearHash[K comparable, V any]() Diccionario[K, V] {

}

func (/*completar*/) Guardar(clave K, dato V) {

}

func (/*completar*/) Pertenece(clave K) bool {

}

func (/*completar*/) Borrar(clave K) V

func (/*completar*/) Cantidad() int 

/* **************** IMPLEMENTACIÓN DEL ITERADOR INTERNO **************** */

func (/*completar*/) Iterar(func(clave K, dato V) bool)


/* **************** IMPLEMENTACIÓN DEL ITERADOR EXTERNO **************** */

func (/*completar*/)  Iterador() IterDiccionario[K, V] {

}

func (/*completar*/) HaySiguiente() bool {

}

func (/*completar*/) VerActual() (K, V) {

}

func (/*completar*/) Siguiente() {

}