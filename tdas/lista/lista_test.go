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

/* ******************************** */

/* **************** PRUEBAS DE LA LISTA **************** */

/* ******************************** */

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
func TestListaInsertarUnElementoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

func TestListaInsertarDiezElementosPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
}

/* **************** InsertarUltimo() **************** */
func TestListaInsertarUnElementoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

func TestListaInsertarDiezElementosUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i * 2)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
}

/* **************** BorrarPrimero() **************** */
func TestListaBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestListaInsertarPrimeroYBorrarDiezElementos(t *testing.T) {
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

func TestListaInsertarUltimoYBorrarDiezElementos(t *testing.T) {
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

func TestListaInsertarMixeadoDiezElementos(t *testing.T) {
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

/* **************** VerPrimero() y VerUltimo() **************** */
func TestListaVerPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerPrimero())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestListaVerUltimoUnElementoInsertandoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerUltimo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerUltimo() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestListaVerUltimoUnElementoInsertandoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(40)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 40, lista.VerUltimo())
	require.Equal(t, 40, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.VerUltimo() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestListaVerDosElementos(t *testing.T) {
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

func TestListaVerMixeadoDiezElementos(t *testing.T) {
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

/* **************** Pruebas de volumen **************** */

func TestListaCienElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 100; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 1, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i := 100; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
		require.Equal(t, 1, lista.VerUltimo())
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestListaDiezMilElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10000; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 1, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
	}
	for i := 10000; i >= 1; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.Largo())
		require.Equal(t, 1, lista.VerUltimo())
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

/* **************** Pruebas con distintos tipos de dato **************** */

func TestListaTipoDeDatoRuna(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[rune]()
	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	for ind, elem := range arr {
		require.Equal(t, ind, lista.Largo())
		if ind%2 == 0 {
			lista.InsertarPrimero(elem)
		} else {
			lista.InsertarUltimo(elem)
		}
		require.False(t, lista.EstaVacia())

	}
	// i g e c a b d f h j
	for i := 8; i >= 0; i -= 2 {
		require.Equal(t, arr[i], lista.BorrarPrimero())
		require.False(t, lista.EstaVacia())
	}
	require.Equal(t, 5, lista.Largo())
	for i := 1; i <= 9; i += 2 {
		require.False(t, lista.EstaVacia())
		require.Equal(t, arr[i], lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestListaTipoDeDatoFloat64(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	for i := float64(1); i <= float64(3); i++ {
		lista.InsertarPrimero(i / 2)
		require.Equal(t, i/2, lista.VerPrimero())
		require.Equal(t, float64(1.0/2), lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, int(i), lista.Largo())
	} // [3 2 1]
	for i := float64(4); i <= float64(7); i++ {
		lista.InsertarUltimo(i / 2)
		require.Equal(t, float64(3.0/2), lista.VerPrimero())
		require.Equal(t, i/2, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, int(i), lista.Largo())
	} // [3 2 1 4 5 6 7]
	for i := float64(8); i <= float64(10); i++ {
		lista.InsertarPrimero(i / 2)
		require.Equal(t, i/2, lista.VerPrimero())
		require.Equal(t, float64(7.0/2), lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, int(i), lista.Largo())
	} //[10 9 8 3 2 1 4 5 6 7]
	for i := float64(10); i >= float64(8); i-- {
		require.Equal(t, i/2, lista.VerPrimero())
		require.Equal(t, float64(7.0/2), lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i/2, lista.BorrarPrimero())
	} //[3 2 1 4 5 6 7]
	require.Equal(t, 7, lista.Largo())
	for i := float64(3); i >= float64(1); i-- {
		require.Equal(t, i/2, lista.VerPrimero())
		require.Equal(t, float64(7.0/2), lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i/2, lista.BorrarPrimero())
	} //[4 5 6 7]
	require.Equal(t, 4, lista.Largo())
	for i := float64(4); i <= float64(7); i++ {
		require.Equal(t, i/2, lista.VerPrimero())
		require.Equal(t, float64(7.0/2), lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i/2, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

/* ******************************** */

/* **************** PRUEBAS DEL ITERADOR INTERNO **************** */

/* ******************************** */

func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var contador int = 0
	lista.Iterar(func(v int) bool {
		fmt.Println(v)
		contador++
		return true
	})
	require.Equal(t, contador, 0, "Cuando iteramos una lista vacia, el contador debería recorrer cero elementos")
}

func TestIteradorInternoIterarDiezElementos(t *testing.T) {
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

func TestIterarInternoIterarConFreno(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaString := TDALista.CrearListaEnlazada[string]()
	letras := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < 10; i++ {
		listaInt.InsertarPrimero(i)
		listaString.InsertarPrimero(letras[i])
	}
	var contador int = 0
	listaInt.Iterar(func(v int) bool {
		contador++
		return v%2 == 0
	})
	require.Equal(t, 1, contador, "Cuando iteramos una lista y la función visitar devuelve false la iteración debe frenar (aunque queden elementos en la lista)")
	contador1 := 0
	listaString.Iterar(func(v string) bool {
		contador1++
		return letras[contador1] == "b"
	})
	require.Equal(t, 2, contador1, "Cuando iteramos una lista y devolvemos false, la iteracion debe frenar, por mas que hayan mas elementos.")
}

func TestIteradorInternoIterarDiezMilElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var contador int = 0
	for i := 1; i <= 10000; i++ {
		lista.InsertarPrimero(i)
	}
	lista.Iterar(func(v int) bool {
		contador++
		return true
	})
	require.Equal(t, contador, 10000, "si inserto 10000 elementos, el contador debe iterar naturalmente 10 veces (sin interrupciones)")
}

//faltan pruebas de volumen y con otro tipo de dato.

/* ******************************** */
/* **************** PRUEBAS DEL ITERADOR EXTERNO **************** */
/* ******************************** */

func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	it := lista.Iterador()
	var contador int = 0
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { it.VerActual() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	for it.HaySiguiente() {
		contador++
		it.Siguiente()
	}
	require.Equal(t, 0, contador, "Cuando iteramos una lista vacia, no hay elementos para iterar")
}

func TestIteradorExternoListaConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[rune]()
	lista.InsertarPrimero('a')
	contador := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); contador++ {
		require.Equal(t, 'a', iter.VerActual())
		iter.Siguiente()
	}
	require.Equal(t, 1, contador)
}

func TestIteradorExternoIteracionCompletaConDiezElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	letras := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for _, letra := range letras {
		lista.InsertarUltimo(letra)
	} //insertandoUltimo queda a b c d e f g h i j
	var i int = 0

	for it := lista.Iterador(); it.HaySiguiente(); i++ {
		require.Equal(t, letras[i], it.VerActual(), "A medida que iteramos la lista, el actual se va moviendo")
		it.Siguiente()
	}
}

func TestIteradorExternoIteracionCompletaConMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numeros := make([]int, 10001)

	for i := 0; i <= 10000; i++ {
		numeros[i] = i
	}
	for i := 0; i <= 10000; i++ {
		lista.InsertarUltimo(i)
	}
	var i int = 0

	for it := lista.Iterador(); it.HaySiguiente(); i++ {
		require.Equal(t, numeros[i], it.VerActual(), "A medida que iteramos la lista, el actual se va moviendo")
		it.Siguiente()
	}
}

/* **************** Insertar() **************** */

func TestIteradorExternoInsertarUnElementoAlMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 4; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	for it := lista.Iterador(); it.HaySiguiente(); it.Siguiente() {
		if contador == 1 {
			it.Insertar(4)
		}
		contador++
	}
	orden := []int{1, 4, 2, 3}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoInsertarUnElementoAlPrincipio(t *testing.T) {
	list := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 4; i++ {
		list.InsertarUltimo(i)
	} // 1 2 3
	contador := 0
	for it := list.Iterador(); it.HaySiguiente(); it.Siguiente() {
		if contador == 0 {
			it.Insertar(4)
			contador++
		}
	}
	orden := []int{4, 1, 2, 3}
	for i := 0; i < len(orden); i++ {
		require.False(t, list.EstaVacia())
		require.Equal(t, list.Largo(), 4-i)
		require.Equal(t, list.BorrarPrimero(), orden[i])
	}
	require.True(t, list.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { list.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
}

func TestIteradorExternoInsertarUnElementoAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 4; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	for it := lista.Iterador(); it.HaySiguiente(); it.Siguiente() {
		if contador == 2 {
			it.Insertar(4)
		}
		contador++
	}
	orden := []int{1, 2, 4, 3}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

/* **************** Borrar() **************** */
func TestIteradorExternoBorrarUnElementoDelPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 6; i++ {
		lista.InsertarUltimo(i)
	} //1 2 3 4 5
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador == 1 {
			require.Equal(t, it.Borrar(), 1)
		}
		if it.HaySiguiente() {
			it.Siguiente()
		}
	}
	orden := []int{2, 3, 4, 5}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoBorrarUnElementoDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 6; i++ {
		lista.InsertarUltimo(i)
	} //1 2 3 4 5
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador == 3 {
			require.Equal(t, it.Borrar(), 3)
		}

		if it.HaySiguiente() {
			it.Siguiente()
		}
	}

	require.Equal(t, lista.Largo(), 4)
	orden := []int{1, 2, 4, 5}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoBorrarUnElementoDelFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 6; i++ {
		lista.InsertarUltimo(i)
	} //1 2 3 4 5
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador == 5 {
			require.Equal(t, it.Borrar(), 5)
		}
		if it.HaySiguiente() {
			it.Siguiente()
		}
	}
	orden := []int{1, 2, 3, 4}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoBorrarDosElementosSeguidos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i < 6; i++ {
		lista.InsertarUltimo(i)
	} //1 2 3 4 5
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador == 5 || contador == 4 {
			require.Equal(t, it.Borrar(), contador)
		} else {
			it.Siguiente()
		}
	}
	require.Equal(t, lista.Largo(), 3)
	orden := []int{1, 2, 3}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoBorrarIntercalado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	} //1 2 3 4 5 6 7 8 9 10
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador%2 == 0 {
			require.Equal(t, it.Borrar(), contador)
		} else {
			it.Siguiente()
		}
	}
	require.Equal(t, lista.Largo(), 5)
	orden := []int{1, 3, 5, 7, 9}
	for i := 0; i < len(orden); i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), orden[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

/* **************** Pruebas de volumen **************** */

func TestIteradorExternoCienElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 100; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador%2 == 0 {
			require.Equal(t, it.Borrar(), contador)
		} else {
			it.Siguiente()
		}
	}
	require.Equal(t, lista.Largo(), 100/2)
	for i := 1; i < 100; i += 2 {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), i)
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoDiezMilElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10000; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 1
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador%2 == 0 {
			require.Equal(t, it.Borrar(), contador)
		} else {
			it.Siguiente()
		}
	}
	require.Equal(t, lista.Largo(), 10000/2)
	for i := 1; i < 10000; i += 2 {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), i)
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}

/* **************** Pruebas con distintos tipos de datos **************** */

func TestIteradorExternoConTipoDeDatoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	arr := []string{"Wall-e", "quiere", "a", "Eva"}
	for i := 0; i < len(arr); i++ {
		lista.InsertarUltimo(arr[i])
	}
	contador := 0
	for it := lista.Iterador(); it.HaySiguiente(); contador++ {
		if contador%2 == 0 {
			require.Equal(t, it.Borrar(), arr[contador])
		} else {
			it.Siguiente()
		}
	}
	require.Equal(t, lista.Largo(), len(arr)/2)
	for i := 1; i < len(arr); i += 2 {
		require.False(t, lista.EstaVacia())
		require.Equal(t, lista.BorrarPrimero(), arr[i])
	}
	require.PanicsWithValue(t, _MENSAJE_PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() }, _MENSAJE_TESTING_PANIC_LISTA_VACIA)
	require.True(t, lista.EstaVacia())
}
