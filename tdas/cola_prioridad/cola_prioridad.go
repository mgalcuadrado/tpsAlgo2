package cola_prioridad

type ColaPrioridad[T any] interface {

	// EstaVacia devuelve true si la la cola se encuentra vacía, false en caso contrario.
	EstaVacia() bool

	// Encolar agrega un elemento a la cola de prioridad.
	Encolar(T)

	// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con el mensaje
	// "La cola esta vacia".
	VerMax() T

	// Desencolar elimina el elemento de la cola con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con el
	// mensaje "La cola esta vacia"
	Desencolar() T

	// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
	Cantidad() int
}
