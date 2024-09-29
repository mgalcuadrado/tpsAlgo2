package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista"
	"testing"
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

func TestInsertarDiezElementosPrimero(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 10; i++ {
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

func TestInsertarDiezElementosUltimo(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 10; i++ {
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

func TestInsertarPrimeroYBorrarDiezElementos(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i:= 10; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestInsertarUltimoYBorrarDiezElementos(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 10; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i:= 1; i <= 10; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

/* **************** Mixeando InsertarPrimero() e InsertarUltimo() **************** */

func TestInsertarMixeadoDiezElementos(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 3; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1]
	for i:= 4; i <= 7; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1 4 5 6 7]
	for i:=8; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} //[10 9 8 3 2 1 4 5 6 7]
	for i:= 10; i >= 8; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	} //[3 2 1 4 5 6 7]
	require.Equal(t, 7, lista.Largo())
	for i:= 3; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	}//[4 5 6 7]
	require.Equal(t, 4, lista.Largo())
	for i:= 4; i <= 7; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
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

func TestVerMixeadoDiezElementos(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 3; i++ {
		lista.InsertarPrimero(i * 2)
		require.Equal(t, i * 2, lista.VerPrimero())
		require.Equal(t, 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1]
	for i:= 4; i <= 7; i++ {
		lista.InsertarUltimo(i * 2)
		require.Equal(t, 3 * 2, lista.VerPrimero())
		require.Equal(t, i * 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} // [3 2 1 4 5 6 7]
	for i:=8; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.Equal(t, i * 2, lista.VerPrimero())
		require.Equal(t, 7 * 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	} //[10 9 8 3 2 1 4 5 6 7]
	for i:= 10; i >= 8; i-- {
		require.Equal(t, i * 2, lista.VerPrimero())
		require.Equal(t, 7 * 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	} //[3 2 1 4 5 6 7]
	require.Equal(t, 7, lista.Largo())
	for i:= 3; i >= 1; i-- {
		require.Equal(t, i * 2, lista.VerPrimero())
		require.Equal(t, 7 * 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	}//[4 5 6 7]
	require.Equal(t, 4, lista.Largo())
	for i:= 4; i <= 7; i++ {
		require.Equal(t, i * 2, lista.VerPrimero())
		require.Equal(t, 7 * 2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i * 2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestIterarInternoPocosElementos(t * testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i:= 1; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}
	lista.Iterar(func MostrarValor(v int) bool{
		fmt.Println(v)
		return true
	})

}
