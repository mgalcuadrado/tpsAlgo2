package diccionario

import TDAPila "tdas/pila"

/*La función de comparación, recibe dos claves y devuelve:

Un entero menor que 0 si la primera clave es menor que la segunda.
Un entero mayor que 0 si la primera clave es mayor que la segunda.
0 si ambas claves son iguales.*/

const (
	_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO_ABB string = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_FIN_DE_ITERACION_ABB                 string = "El iterador termino de iterar"
)

/* ************ Definición de estructuras ************ */

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cmp      func(K, K) int //la funcion de comparacion es un atributo del abb
	cantidad int
}

type iteradorExternoRango[K comparable, V any] struct {
	desde *K
	hasta *K
	cmp   func(K, K) int
	pila  TDAPila.Pila[*nodoABB[K, V]]
}

type nodoABB[K comparable, V any] struct {
	clave              K
	valor              V
	izquierda, derecha *nodoABB[K, V]
}

/* ************ Funciones de creación de las estructuras ************ */
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cmp:      funcion_cmp,
		cantidad: 0,
	}
}

func crearNodoABB[K comparable, V any](clave K, dato V, izq *nodoABB[K, V], der *nodoABB[K, V]) *nodoABB[K, V] {
	nodo := new(nodoABB[K, V])
	nodo.clave = clave
	nodo.valor = dato
	nodo.izquierda, nodo.derecha = izq, der
	return nodo
}

func crearIteradorExternoRango[K comparable, V any](desde *K, hasta *K, cmp func(K, K) int, raiz *nodoABB[K, V]) *iteradorExternoRango[K, V] {
	iter := &iteradorExternoRango[K, V]{
		desde: desde,
		hasta: hasta,
		cmp:   cmp,
		pila:  TDAPila.CrearPilaDinamica[*nodoABB[K, V]](),
	}
	if raiz == nil {
		return iter
	}
	nodo := raiz
	iter.apilarNodosDesdeDerecha(raiz)
	for iter.pila.EstaVacia() {
		if nodo == nil {
			break
		}
		iter.apilarNodosDesdeDerecha(nodo)
		nodo = nodo.derecha
	}
	return iter
}

/* ************ FUNCIONES DEL ABB ************ */
// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (abb *abb[K, V]) Guardar(clave K, dato V) {
	ref := abb.buscar(clave, &abb.raiz)
	if *ref != nil {
		(*ref).valor = dato
	} else {
		nodo := crearNodoABB(clave, dato, nil, nil)
		(*ref) = nodo
		abb.cantidad++
	}
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (abb *abb[K, V]) Pertenece(clave K) bool {
	ref := abb.buscar(clave, &abb.raiz)
	return *ref != nil
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (abb *abb[K, V]) Obtener(clave K) V {
	ref := abb.obtenerReferenciaValida(clave)
	return (*ref).valor
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (abb *abb[K, V]) Borrar(clave K) V {
	ref := abb.obtenerReferenciaValida(clave)
	abb.cantidad--
	//Caso 1: nodo sin hijos
	if (*ref).izquierda == nil && (*ref).derecha == nil {
		valor := (*ref).valor
		*ref = nil
		return valor
	}
	reemplazante := ref
	//Caso 2: nodo con dos hijos
	//Nos quedamos con el más grande de los chicos
	if (*ref).izquierda != nil && (*ref).derecha != nil {
		reemplazante = abb.obtenerReemplazante(&((*ref).izquierda))
		claveReemplazante, valorReemplazante := (*reemplazante).clave, (*reemplazante).valor
		*reemplazante = (*reemplazante).izquierda
		valor := (*ref).valor
		(*ref).clave, (*ref).valor = claveReemplazante, valorReemplazante
		return valor
	}
	//Caso 3: nodo con un hijo
	if (*ref).izquierda != nil {
		reemplazante = &((*ref).izquierda)
	}
	if (*ref).derecha != nil {
		reemplazante = &((*ref).derecha)
	}
	valor := (*ref).valor
	*ref = *reemplazante
	return valor
}

/* ************ Funciones auxiliares  del abb ************ */

// buscar recibe la clave a buscar y un puntero doble al nodo en el que se está buscando. Devuelve un puntero al nodo buscado (si la clave no pertenece al diccionario, será un puntero a nil). Recursivamente se llama a sí misma descartando sectores del árbol usando la función de comparación abb.cmp
func (abb *abb[K, V]) buscar(clave K, raiz **nodoABB[K, V]) **nodoABB[K, V] {
	if *raiz == nil {
		return raiz
	}
	if abb.cmp((*raiz).clave, clave) == 0 {
		return raiz
	}
	if abb.cmp((*raiz).clave, clave) > 0 {
		return abb.buscar(clave, &((*raiz).izquierda))
	} else {
		return abb.buscar(clave, &((*raiz).derecha))
	}
}

// obtenerReemplazante recibe una puntero puntero a un nodo y busca al más grande a partir de esa referencia.
func (abb *abb[K, V]) obtenerReemplazante(referencia **nodoABB[K, V]) **nodoABB[K, V] {
	if (*referencia).derecha == nil {
		return referencia
	}
	return abb.obtenerReemplazante(&((*referencia).derecha))
}

// obtenerReferenciaValida recibe una clave, la busca en el árbol y obtiene una referencia a la misma, verificando si se encontraba en el diccionario. De no encontrarse devuelve un panic
func (abb *abb[K, V]) obtenerReferenciaValida(clave K) **nodoABB[K, V] {
	ref := abb.buscar(clave, &abb.raiz)
	if *ref == nil {
		panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO_ABB)
	}
	return ref
}

/* ************ Funciones del iterador interno ************ */

// Iterar recibe una función visitar y visita in-order todos los nodos del árbol hasta que visitar dé false
func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if abb == nil {
		return
	}

	iterador_interno(abb.raiz, nil, nil, abb.cmp, visitar)
}

// Iterar recibe una función visitar y visita in-order todos los nodos del árbol cuyas clavesse encuentren entre lo referenciado por desde y hasta mientras que visitar dé true
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if abb == nil {
		return
	}

	iteradorInterno(abb.raiz, desde, hasta, abb.cmp, visitar)
}

// iterador_interno recibe un nodo y una función de verificación para saber si la clave analizada se encuentra o no en el rango
// no se incluye en la estructura para evitar una dependencia circular, pues puede depender de los límites del rango guardados en el iterador (desde y hasta)
func iterador_interno[K comparable, V any](nodo *nodoABB[K, V], desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde != nil && cmp(nodo.clave, *desde) < 0 && !iterador_interno(nodo.derecha, desde, hasta, cmp, visitar) {
		return false
	}
	if hasta != nil && cmp(nodo.clave, *hasta) > 0 && !iterador_interno(nodo.izquierda, desde, hasta, cmp, visitar) {
		return false
	}
	if (nodo.izquierda != nil && !iterador_interno(nodo.izquierda, desde, hasta, cmp, visitar)) || !visitar(nodo.clave, nodo.valor) || (nodo.derecha != nil && !iterador_interno(nodo.derecha, desde, hasta, cmp, visitar)) {
		return false
	}
	return true
}

func iteradorInterno[K comparable, V any](nodo *nodoABB[K, V], desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde != nil && cmp(nodo.clave, *desde) < 0 { //todo lo que esté a izquierda de un valor menor a desde será menor a desde
		return iteradorInterno(nodo.derecha, desde, hasta, cmp, visitar) //si se llega a este punto iteradorInterno(nodo.derecha, ...) devolvió false
	} else if hasta != nil && cmp(nodo.clave, *hasta) > 0 { //todo lo que esté a derecha de un valor mayor a hasta será mayor a hasta
		return iteradorInterno(nodo.izquierda, desde, hasta, cmp, visitar)
	} else { //si se llega a este punto entonces nodo.clave está en el rango
		return iteradorInterno(nodo.izquierda, desde, hasta, cmp, visitar) && visitar(nodo.clave, nodo.valor) && iteradorInterno(nodo.derecha, desde, hasta, cmp, visitar)
	}
}

/* ************ FUNCIONES DE ITERADORES EXTERNOS ************ */

// Iterador devuelve un IterDiccionario para este Diccionario
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] { //comento esto para poder correr las pruebas
	return crearIteradorExternoRango[K, V](nil, nil, abb.cmp, abb.raiz)
}

// IteradorRango crea un IterDiccionario que sólo itera por las claves que se encuentren en el rango entre desde y hasta (inclusive)
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorExternoRango[K, V](desde, hasta, abb.cmp, abb.raiz)
}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
func (iter *iteradorExternoRango[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter *iteradorExternoRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_FIN_DE_ITERACION_ABB)
	}
	actual := iter.pila.VerTope()
	return actual.clave, actual.valor
}

// Siguiente si HaySiguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter *iteradorExternoRango[K, V]) Siguiente() {
	iter.VerActual()
	nodo := iter.pila.Desapilar()
	if nodo.derecha == nil {
		return
	}
	iter.apilarNodosDesdeDerecha(nodo.derecha)
}

/***************Funciones auxiliares iterador externo *************/

func (iter *iteradorExternoRango[K, V]) apilarNodosDesdeDerecha(nodo *nodoABB[K, V]) {
	nodoActual := nodo
	for nodoActual != nil {
		if iter.desde != nil && iter.cmp(nodoActual.clave, *(iter.desde)) < 0 && nodoActual.derecha != nil {
			iter.apilarNodosDesdeDerecha(nodoActual.derecha)
			break
		} else if iter.hasta == nil || (iter.hasta != nil && iter.cmp(nodoActual.clave, *(iter.hasta)) <= 0) {
			iter.pila.Apilar(nodoActual)
		}
		nodoActual = nodoActual.izquierda
	}
}
