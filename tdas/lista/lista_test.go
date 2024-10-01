package lista_test

import (
	"fmt"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC_LISTA_VACIA         string = "La lista esta vacia"
	_MENSAJE_TESTING_PANIC_LISTA_VACIA string = "No hay elementos en la lista"
)

/* **************** EstaVacia() **************** */
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerUltimo() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

/* **************** InsertarPrimero() **************** */
func TestInsertarUnElementoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

func TestInsertarDiezElementosPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
}

/* **************** InsertarUltimo() **************** */
func TestInsertarUnElementoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

func TestInsertarDiezElementosUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
}

/* **************** BorrarPrimero() **************** */
func TestBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestInsertarPrimeroYBorrarDiezElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i := 10; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestInsertarUltimoYBorrarDiezElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i := 1; i <= 10; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

/* **************** Mixeando InsertarPrimero() e InsertarUltimo() **************** */

func TestInsertarMixeadoDiezElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 3; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1]
	for i := 4; i <= 7; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1 4 5 6 7]
	for i := 8; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} //[10 9 8 3 2 1 4 5 6 7]
	for i := 10; i >= 8; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	} //[3 2 1 4 5 6 7]
	require.Equal(t, 7, lista.Largo())
	for i := 3; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	} //[4 5 6 7]
	require.Equal(t, 4, lista.Largo())
	for i := 4; i <= 7; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

//yo acá agregaría alguna prueba más

/* **************** VerPrimero() y VerUltimo() **************** */
func TestVerPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerPrimero())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestVerUltimoUnElementoInsertandoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerUltimo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerUltimo() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestVerUltimoUnElementoInsertandoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerUltimo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerUltimo() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestVerDosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.Equal(t, 40, lista.VerPrimero())
	require.Equal(t, 40, lista.VerUltimo())
	lista.InsertarUltimo(6)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 40, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.Equal(t, lista.VerPrimero(), lista.VerUltimo())
	require.Equal(t, 6, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestVerMixeadoDiezElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 3; i++ {
		lista.InsertarPrimero(i * 2)
		require.Equal(t, i*2, lista.VerPrimero())
		require.Equal(t, 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1]
	for i := 4; i <= 7; i++ {
		lista.InsertarUltimo(i * 2)
		require.Equal(t, 3*2, lista.VerPrimero())
		require.Equal(t, i*2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1 4 5 6 7]
	for i := 8; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.Equal(t, i*2, lista.VerPrimero())
		require.Equal(t, 7*2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} //[10 9 8 3 2 1 4 5 6 7]
	for i := 10; i >= 8; i-- {
		require.Equal(t, i*2, lista.VerPrimero())
		require.Equal(t, 7*2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	} //[3 2 1 4 5 6 7]
	require.Equal(t, 7, lista.Largo())
	for i := 3; i >= 1; i-- {
		require.Equal(t, i*2, lista.VerPrimero())
		require.Equal(t, 7*2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	} //[4 5 6 7]
	require.Equal(t, 4, lista.Largo())
	for i := 4; i <= 7; i++ {
		require.Equal(t, i*2, lista.VerPrimero())
		require.Equal(t, 7*2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i*2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestIterarInternoPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var contador int = 0
	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}
	lista.Iterar(func(v int) bool {
		contador++
		return true
	})
	require.Equal(t, contador, 10, "si inserto 10 elementos, el contador debe iterar naturalmente 10 veces (sin interrupciones)")

}

func TestIterarInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var contador int = 0
	lista.Iterar(func(v int) bool {
		fmt.Println(v)
		contador++
		return true
	})
	require.Equal(t, contador, 0, "Cuando iteramos una lista vacia, el contador debe acumular 0 vueltas")
}

func TestIterarInternoCortaIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 2; i++ {
		lista.InsertarPrimero(i)
	}
	var contador int = 0
	lista.Iterar(func(v int) bool {
		contador++
		return v%2 == 0
	})
	require.Equal(t, 2, contador, "Cuando iteramos una lista y devolvemos false, la iteracion debe frenar, por mas que hayan mas elementos.")
}

func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	it := lista.Iterador()

	var contador int = 0
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { it.VerActual() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	for it.HaySiguiente() {
		contador++
		it.Siguiente()
	}

	require.Equal(t, 0, contador, "Cuando iteramos una lista vacia, el contador acumula 0")

}

func TestIteradorExternoRecorridoCompletoPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	letras := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, letra := range letras {
		lista.InsertarPrimero(letra)
	}

	var i int = 0
	it := lista.Iterador()

	for it.HaySiguiente() {
		require.Equal(t, letras[i], it.VerActual(), "a medida que iteramos la lista, el actual se va moviendo")
		it.Siguiente()
	}
}

/*
func TestIteradorExternoRecorridoInsertarYBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}

	it := lista.Iterador()

	for it.HaySiguiente() {

		if it.VerActual()%2 == 0 {
			it.Insertar(100)
		}

		it.Siguiente()
	}
	require.Equal(t, 16, lista.Largo(), "si insertamos intercaladamente en posiciones pares, el largo es 16")
}
*/
