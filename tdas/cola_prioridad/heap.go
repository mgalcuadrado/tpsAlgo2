package cola_prioridad

/* ******* DECLARACIÓN DE VARIABLES Y ESTRUCTURAS ********** */
const (
	_CAPACIDAD_INICIAL        = 8
	_FACTOR_REDIM             = 2
	_INDICADOR_CAPACIDAD      = 4
	_MENSAJE_PANIC_COLA_VACIA = "La cola esta vacia"
)

type colaConPrioridad[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

/* ******* FUNCIONES DEL TDA ********** */

// CrearHeapArr recibe una función de comparación y crea una cola de prioridad
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
		cmp:      funcion_cmp,
		cantidad: 0,
	}
}

// CrearHeapArr recibe un arreglo y una función de comparación y crea una cola de prioridad con los datos de ese arreglo
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, cap(arreglo))
	copy(arr, arreglo) //copio el arreglo en memoria nueva para no modificarle el arreglo original a Bárbara!
	heapify[T](&arr, funcion_cmp)
	return &colaConPrioridad[T]{
		datos:    arr,
		cmp:      funcion_cmp,
		cantidad: len(arr),
	}
}

// EstaVacia indica si la cola de prioridad se encuentra vacía
func (ccp *colaConPrioridad[T]) EstaVacia() bool {
	return ccp.cantidad == 0
}

// Encolar recibe un dato y lo encola en función de su prioridad
func (ccp *colaConPrioridad[T]) Encolar(dato T) {
	if ccp.cantidad == cap(ccp.datos) {
		ccp.redimensionar(cap(ccp.datos) * _FACTOR_REDIM)
	}
	ccp.datos[ccp.cantidad] = dato
	ccp.upheap(ccp.cantidad)
	ccp.cantidad++
}

// VerMax devuelve el máximo de la cola sin modificarla
func (ccp *colaConPrioridad[T]) VerMax() T {
	if ccp.EstaVacia() {
		panic(_MENSAJE_PANIC_COLA_VACIA)
	}
	return ccp.datos[0]
}

// Desencolar saca el elemento de mayor prioridad de la cola y devuelve su valor
func (ccp *colaConPrioridad[T]) Desencolar() T {
	valor := ccp.VerMax()
	swap(&(ccp.datos), 0, ccp.cantidad-1)
	ccp.cantidad--
	if ccp.cantidad > _CAPACIDAD_INICIAL && ccp.cantidad == cap(ccp.datos)/_INDICADOR_CAPACIDAD {
		ccp.redimensionar(cap(ccp.datos) / _FACTOR_REDIM)
	}
	downheap(0, &(ccp.datos), ccp.cmp, ccp.cantidad)
	return valor
}

// Cantidad devuelve la cantidad de elementos en la cola
func (ccp *colaConPrioridad[T]) Cantidad() int {
	return ccp.cantidad
}

// HeapSort recibe un puntero a un arreglo y una funcion de comparacion y ordena el arreglo de menor a mayor (in-place)
func HeapSort[T any](elementos *[]T, funcion_cmp func(T, T) int) {
	heapify[T](elementos, funcion_cmp)
	for cantidad := len(*elementos) - 1; cantidad > 0; cantidad-- { //arranco en len - 1 porque trabajo con posiciones hasta -1
		swap(elementos, 0, cantidad)
		downheap(0, elementos, funcion_cmp, cantidad)
	}
}

/* ******* FUNCIONES AUXILIARES ********** */

// redimensionar recibe la nueva capacidad que se le quiere asignar al arreglo de datos del heap,
// y realiza la resimension del mismo
func (ccp *colaConPrioridad[T]) redimensionar(capacidad int) {
	datosNuevo := make([]T, capacidad)
	copy(datosNuevo, ccp.datos)
	ccp.datos = datosNuevo
}

// heapify recibe un puntero a un arreglo y una función de comparación y organiza el arreglo para que cumpla con las propiedades del heap
func heapify[T any](arreglo *[]T, funcion_cmp func(T, T) int) {
	for i := cap(*arreglo) - 1; i >= 0; i-- {
		downheap(i, arreglo, funcion_cmp, len(*arreglo))
	}
}

// upheap recibe una posicion y hace cumplir la propiedad de heap hasta la raíz.
// upheap asume que hacia arriba hay un heap
func (ccp *colaConPrioridad[T]) upheap(pos int) {
	if pos == 0 {
		return
	}
	posPadre := hallarPosicionPadre(pos)
	if ccp.cmp(ccp.datos[pos], ccp.datos[posPadre]) > 0 {
		swap(&(ccp.datos), pos, posPadre)
	}
	ccp.upheap(posPadre)
}

// downheap recibe una posición, un puntero a un arreglo, una función de comparación cmp y la cantidad de elementos válidos del arreglo y acomoda el heap con downheap desde la posición pedida
func downheap[T any](pos int, ptr_arr *[]T, cmp func(T, T) int, cantidad int) {
	posHijoIzq := hallarPosicionHijoIzquierdo(pos)
	posHijoDer := hallarPosicionHijoDerecho(pos)
	if posHijoIzq >= cantidad { // Caso base: hoja
		return
	}
	arr := *ptr_arr
	// Se determina el hijo con mayor prioridad. En empate se prioriza el izquierdo
	posMayor := posHijoIzq
	if posHijoDer < cantidad && cmp(arr[posHijoDer], arr[posHijoIzq]) > 0 {
		// se elige el hijo de mayor prioridad, o el izquierdo en caso de empate
		posMayor = posHijoDer
	}
	// Si el hijo mayor es mas grande que el padre, se intercambian y se realiza el sucesivo downheap
	if cmp(arr[pos], arr[posMayor]) < 0 {
		swap(ptr_arr, pos, posMayor)
		downheap(posMayor, ptr_arr, cmp, cantidad)
	}
}

// swap recibe un puntero a un arreglo y dos posiciones e intercambia los elementos en esas posiciones de lugar
func swap[T any](ptr_arr *[]T, pos1 int, pos2 int) {
	(*ptr_arr)[pos1], (*ptr_arr)[pos2] = (*ptr_arr)[pos2], (*ptr_arr)[pos1]
}

// hallarPosicionHijoIzquierdo devuelve, en base a una posición de un arreglo, en qué posición estaría su hijo izquierdo
func hallarPosicionHijoIzquierdo(pos int) int {
	return 2*pos + 1
}

// hallarPosicionHijoDerecho devuelve, en base a una posición de un arreglo, en qué posición estaría su hijo derecho
func hallarPosicionHijoDerecho(pos int) int {
	return 2*pos + 2
}

// hallarPosicionPadre devuelve, en base a una posición de un arreglo, en qué posición estaría su padre
func hallarPosicionPadre(pos int) int {
	return (pos - 1) / 2
}
