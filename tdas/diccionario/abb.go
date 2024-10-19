package diccionario


//por qué los llamas ab y no abb?
type nodoABB[K comparable, V any] struct {
    clave   K
    valor   V
    izquierda, derecha *nodoABB[K, V]
}

type abb[K comparable, V any] struct {
    raiz       *nodoABB[K, V]
    cmp        func(K, K) int //la funcion de comparacion es un atributo del ab
    cantidad   int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
    return &abb[K, V]{
        raiz:     nil,
        cmp:      funcion_cmp,
        cantidad: 0,
    }
}

func crearNodoABB[K comparable, V any](clave K, dato V, izq *nodoABB[K,V], der *nodoABB[K,V]) *nodoABB[K,V] {
    nodo := new(nodoABB[K, V])
    nodo.clave = clave
    nodo.valor = dato
    nodo.izquierda, nodo.derecha = izq, der
    return nodo
}

/*La función de comparación, recibe dos claves y devuelve:

Un entero menor que 0 si la primera clave es menor que la segunda.
Un entero mayor que 0 si la primera clave es mayor que la segunda.
0 si ambas claves son iguales.*/

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (abb *abb[K, V]) Guardar(clave K, dato V) { 
    ref, pertenece := abb.buscar(clave, &abb.raiz) 
     if pertenece {
        (*ref).valor = dato
    } else {
        nodo := crearNodoABB(clave, dato, nil, nil)
        (*ref) = nodo
        abb.cantidad++
    }
}

// buscar recibe la clave a buscar y un puntero doble al nodo en el que se está buscando. Devuelve un puntero al nodo buscado y un bool indicando si se hallaba o no el elemento. Recursivamente se llama a sí misma descartando sectores del árbol usando la función de comparación abb.cmp
func (abb *abb[K, V]) buscar(clave K, raiz **nodoABB[K, V]) (**nodoABB[K, V], bool) {
   // if abb == nil { //a mirar 
     //   return &abb, false
   // }
    if *raiz == nil {
        return raiz, false
    }
    if (*raiz).clave == clave {
        return raiz, true
    } 
    if (*raiz).izquierda != nil && abb.cmp((*raiz).izquierda.clave, (*raiz).clave) < 0 {
        return abb.buscar(clave, &((*raiz).izquierda))
    } else {
        return abb.buscar(clave, &((*raiz).derecha))
    }
}


// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (abb *abb[K, V]) Pertenece(clave K) bool {
    _, pertenece := abb.buscar(clave, &abb.raiz)
    return pertenece
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
        *reemplazante = nil
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

func (abb *abb[K, V]) obtenerReemplazante(referencia **nodoABB[K,V])  **nodoABB[K,V]{
    if (*referencia).derecha == nil {
        return referencia
    }
    return abb.obtenerReemplazante(&((*referencia).derecha))
}

func (abb *abb[K, V]) obtenerReferenciaValida(clave K) **nodoABB[K,V] {
    if abb.Cantidad() == 0 {
        panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
    }
    ref, pertenece := abb.buscar(clave, &abb.raiz)
    if !pertenece {
        panic(_MENSAJE_PANIC_CLAVE_NO_PERTENECE_A_DICCIONARIO)
    }
    return ref
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (abb *abb[K, V]) Cantidad() int {
    return abb.cantidad
}

func (ab *abb[K, V]) insertarNodo(nodo *nodoABB[K, V], clave K, dato V) *nodoABB[K, V] {
    if nodo == nil {
        ab.cantidad++
        return &nodoABB[K, V]{clave: clave, valor: dato}
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


//func (abb *abb[K, V]) Iterar(func(clave K, dato V) bool)

	// Iterador devuelve un IterDiccionario para este Diccionario
//func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] { //comento esto para poder correr las pruebas
//}
//func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)

// IteradorRango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
//func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
