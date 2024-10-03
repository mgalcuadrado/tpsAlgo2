package lista

type IteradorLista[T any] interface {

	// HaySiguiente devuelve true cuando hay un elemento luego del actual para ver y false cuando se terminó la lista
	HaySiguiente() bool

	// VerActual devuelve el elemento guardado en el elemento actual en el que se encuentra el iterador dentro de la lista
	VerActual() T

	// Siguiente avanza el iterador al siguiente elemento de la lista
	Siguiente()

	// Insertar permite agregar un elemento a la lista entre el anterior y el actual
	Insertar(T)

	// Borrar permite eliminar el elemento actual de la lista
	Borrar() T
}

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos agregados y false en caso contrario
	EstaVacia() bool

	// InsertarPrimero agrega un elemento al principio de la lista
	InsertarPrimero(T)

	// InsertarUltimo agrega un elemento al final de la lista
	InsertarUltimo(T)

	// BorrarPrimero elimina el elemento al principio de la lista
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al principio de la lista (sin eliminarlo)
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista
	// Si la lista está vacía devuelve 0
	Largo() int

	// Iterar itera la lista internamente visitando cada elemento de la misma siempre que la función auxiliar visitar devuelva true; visitar recibe un dato T de la lista y devuelve true o false
	Iterar(visitar func(T) bool)

	// Iterador crea un IteradorLista para la lista.
	Iterador() IteradorLista[T]
}
