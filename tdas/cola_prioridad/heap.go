package cola_prioridad

import(
	"fmt"
)

const(
	_CAPACIDAD_INICIAL=8
	_FACTOR_REDIM=2
	_INDICADOR_CAPACIDAD=4
	_MENSAJE_PANIC_COLA_VACIA="La cola esta vacia"
)

type colaConPrioridad [T any] struct {
	datos []T	
	cantidad int
	cmp func(T,T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T]{
	return &colaConPrioridad[T]{
		datos:     make([]T, _CAPACIDAD_INICIAL),
		cmp:      funcion_cmp,
		cantidad: 0,
	}
}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]{
	heapify(&arreglo)
	return &colaConPrioridad[T]{
		datos:     arreglo,
		cmp:      funcion_cmp,
		cantidad: len(arreglo),
	}
}

func (ccp *colaConPrioridad[T]) EstaVacia() bool{
	return ccp.cantidad==0
}

func (ccp *colaConPrioridad[T]) Encolar(dato T) {
	if ccp.cantidad==cap(ccp.datos){
		ccp.redimensionar(cap(ccp.datos)*_FACTOR_REDIM)
	}
	ccp.datos[ccp.cantidad] = dato

	ccp.upheap(ccp.cantidad)
	ccp.cantidad++
}

func (ccp *colaConPrioridad[T]) VerMax() T {
	if ccp.EstaVacia(){
		panic(_MENSAJE_PANIC_COLA_VACIA)
	}
	return ccp.datos[0]
}

func (ccp *colaConPrioridad[T]) Desencolar() T {

	if ccp.EstaVacia(){
		panic(_MENSAJE_PANIC_COLA_VACIA)
	}

	valor:=ccp.datos[0]
	ccp.datos[0],ccp.datos[ccp.cantidad-1]=ccp.datos[ccp.cantidad-1], ccp.datos[0]
	ccp.cantidad--

	if ccp.cantidad>_CAPACIDAD_INICIAL && ccp.cantidad==cap(ccp.datos)/_INDICADOR_CAPACIDAD {
		ccp.redimensionar(cap(ccp.datos)/_FACTOR_REDIM)
	}

	ccp.downheap(0)
	ccp.mostrar()

	return valor
}

func (ccp *colaConPrioridad[T]) Cantidad() int {
	return ccp.cantidad
}

//heapsort recibe un arreglo y una funcion de comparacion, y modifica el arreglo de tal manera que cumpla
//la propiedad de heap
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int){
	return
}

/********FUNCIONES AUXILIARES***********/

//upheap recibe una posicion y hace cumplir la propiedad de heap hasta la raíz
//upheap asume que hacia arriba hay un heap
func (ccp *colaConPrioridad[T]) upheap(pos int) {
	posPadre:=(pos-1)/2

	if pos==0{
		return
	}
	
	if ccp.cmp(ccp.datos[pos],ccp.datos[posPadre]) > 0 {
		ccp.datos[pos],ccp.datos[posPadre]=ccp.datos[posPadre],ccp.datos[pos]
	}

	ccp.upheap(posPadre)
}

//downheap recibe una posicion y hace cumplir la propiedad hacia abajo
//downheap asume que hacia arriba hay un heap
func (ccp *colaConPrioridad[T]) downheap(pos int){
	posHijoIzq:=2*pos+1
	posHijoDer:=2*pos+2
	if posHijoDer > ccp.cantidad && posHijoIzq > ccp.cantidad {
		return
	}

		if ccp.cmp(ccp.datos[posHijoDer],ccp.datos[posHijoIzq]) >= 0 && ccp.cmp(ccp.datos[pos],ccp.datos[posHijoDer]) < 0 {
			ccp.datos[pos],ccp.datos[posHijoDer]=ccp.datos[posHijoDer],ccp.datos[pos]
			ccp.downheap(posHijoDer)
		} else if ccp.cmp(ccp.datos[posHijoDer],ccp.datos[posHijoIzq]) < 0 && ccp.cmp(ccp.datos[pos],ccp.datos[posHijoIzq]) < 0{
			ccp.datos[pos],ccp.datos[posHijoIzq]=ccp.datos[posHijoIzq],ccp.datos[pos]
			ccp.downheap(posHijoIzq)
		}
}

//redimensionar recibe la nueva capacidad que se le quiere asignar al arreglo de datos del heap,
//y realiza la resimension del mismo
func (ccp *colaConPrioridad[T]) redimensionar(capacidad int) {
	datosNuevo := make([]T, capacidad)
	copy(datosNuevo, ccp.datos)
	ccp.datos = datosNuevo
}

/******MUY AUXILIAR***/

// mostrar imprime el arreglo datos (solo para probar, despues la sacamos)
func (ccp *colaConPrioridad[T]) mostrar() {
	for _, val:=range(ccp.datos){
		fmt.Printf("%v", val)
	}
	fmt.Printf(" \n")
}
