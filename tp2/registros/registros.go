package registros

type Registros interface {

	//AgregarArchivo agrega al registro la información en el registro ruta.
	//Devuelve un booleano indicando si la operación se pudo realizar correctamente
	AgregarArchivo(ruta string) bool

	//VerVisitantes recibe un rango de IPs y permite ver los visitantes entre desde y hasta.
	//Devuelve un booleano indicando si la operación se pudo realizar correctamente
	VerVisitantes(desde IPv4, hasta IPv4) bool

	//VerMasVisitados muestra los n sitios más visitados.
	//Devuelve un booleano indicando si la operación se pudo realizar correctamente
	VerMasVisitados(n int) bool
}
