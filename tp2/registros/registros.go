package registros

type Registros interface {
	AgregarArchivo(ruta string)
	VerVisitantes(desde IPv4, hasta IPv4)
	VerMasVisitados(n int)
}
