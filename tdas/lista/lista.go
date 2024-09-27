package lista

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

/*	
	//Iterar(visitar func(T) bool)
	
	
	//Iterador() IteradorLista[T]
*/
}