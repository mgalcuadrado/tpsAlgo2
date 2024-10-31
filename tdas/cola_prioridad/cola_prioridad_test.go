package cola_prioridad_test

import (
	"fmt"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC_COLA_VACIA         = "La cola esta vacia"
	_MENSAJE_TESTING_PANIC_COLA_VACIA = "Si la cola esta vacia, debe generarse un panic"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func comparacion(v1 int, v2 int) int {
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return -1
	} else {
		return 0
	}
}

func TestColaPrioridadVacia(t *testing.T) {
	t.Log("Comprueba que la cola de prioridad vacia no tiene datos")
	cp := TDAColaPrioridad.CrearHeap[int](comparacion)
	require.EqualValues(t, 0, cp.Cantidad(), "Una cola sin elementos debe tener 0 de cantidad")
	require.True(t, true, cp.EstaVacia(), "Una cola sin elementos debe estar vacia")
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() {
		cp.VerMax()
	}, "Ver el máximo de una cola vacía debe generar un panic")
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() {
		cp.Desencolar()
	}, "Desencolar una cola vacia debe generar un panic")
}

func TestColaPrioridadPocosElementos(t *testing.T) {
	t.Log("Comprueba que pocos elementos se encolen correctamente en la cola de prioridad")
	cp := TDAColaPrioridad.CrearHeap[int](comparacion)
	for i := 0; i < 10; i++ {
		cp.Encolar(i)
		require.EqualValues(t, cp.VerMax(), i, "Encolando cíclicamente y ascendentemente, el máximo debe ser el elemento recién insertado.")
	}
	require.EqualValues(t, 10, cp.Cantidad(), "La cola debe tener 9 elementos ingresados")
	require.False(t, cp.EstaVacia(), "Una cola sin elementos debe estar vacía")
	require.EqualValues(t, 9, cp.VerMax(), "El máximo debe ser el elemento mayor ingresado (9)")
	for j := 9; j >= 0; j-- {
		require.EqualValues(t, cp.Desencolar(), j, "Desencolando ciclicamente y descendentemente, desencolar devuelve al desencolado")
	}
}

func TestColaPrioridadIniciandoConArregloDeEnteros(t *testing.T) {
	t.Log("Comprueba que se cree bien la cola de prioridad a partir de un arreglo")
	arreglo := []int{12, 3, 45, 2, 4, 50, 22, 16, 7}
	cp := TDAColaPrioridad.CrearHeapArr[int](arreglo, comparacion)
	arr_ordenado := []int{50, 45, 22, 16, 12, 7, 4, 3, 2}
	for i := 0; i < len(arreglo); i++ {
		require.EqualValues(t, cp.Desencolar(), arr_ordenado[i], "Se desencola del máximo al mínimo en una cola de prioridad de máximos")
	}
	require.True(t, cp.EstaVacia())
}

func TestColaPrioridadIniciandoConArregloDeStrings(t *testing.T) {
	t.Log("Comprueba que se cree bien la cola de prioridad a partir de un arreglo")
	animales := []string{"Gato", "Perro", "Vaca", "Pato", "Sapo", "Tortuga", "Conejo", "Avestruz", "Oruga", "Mapache"}
	cp := TDAColaPrioridad.CrearHeapArr[string](animales, strings.Compare)
	arr_ordenado := []string{"Vaca", "Tortuga", "Sapo", "Perro", "Pato", "Oruga", "Mapache", "Gato", "Conejo", "Avestruz"}
	for i := 0; i < len(animales); i++ {
		require.EqualValues(t, cp.VerMax(), arr_ordenado[i])
		require.EqualValues(t, cp.Desencolar(), arr_ordenado[i], "Se desencola del máximo al mínimo en una cola de prioridad de máximos")
	}
	require.True(t, cp.EstaVacia())
}

func ejecutarPruebasVolumenColaPrioridad(b *testing.B, n int) {
	cp := TDAColaPrioridad.CrearHeap[int](comparacion)

	/* Inserta 'n' elementos en el heap */
	for i := 0; i < n; i++ {
		cp.Encolar(i)
	}
	max := cp.VerMax()
	// Prueba de iteración sobre las claves almacenadas.
	for i := n - 1; i >= 0; i-- {
		require.False(b, cp.EstaVacia())
		require.Equal(b, i, cp.VerMax())
		require.True(b, max >= cp.VerMax())
		require.Equal(b, i, cp.Desencolar())
	}
}

func BenchmarkColaPrioridad(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenColaPrioridad(b, n)
			}
		})
	}
}

/* ******* TESTS HEAP SORT ********** */

func TestHeapSortUnElemento(t *testing.T) {
	t.Log("Comprueba que el heap sort funcione correctamente para un solo elemento")
	arr := []string{"Hola"}
	TDAColaPrioridad.HeapSort(&arr, strings.Compare)
	for _, valor := range arr {
		require.Equal(t, "Hola", valor)
	}
	require.Equal(t, 1, len(arr))
}

func TestHeapSortPocosElementos(t *testing.T) {
	t.Log("Comprueba que el heap sort ordene correctamente pocos elementos")
	arr := []int{6, 1, 5, 4, 10, 2}
	arr_ordenado := []int{1, 2, 4, 5, 6, 10}
	TDAColaPrioridad.HeapSort(&arr, comparacion)
	for indice, valor := range arr {
		require.Equal(t, arr_ordenado[indice], valor)
	}
	require.Equal(t, 6, len(arr))
}

func ejecutarPruebasVolumenHeapSort(b *testing.B, n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - 1 - i
	}
	require.Equal(b, n, len(arr))
	TDAColaPrioridad.HeapSort(&arr, comparacion)
	require.Equal(b, n, len(arr))
	anterior := arr[0]
	for i := 0; i < n; i++ {
		require.Equal(b, i, arr[i])
		require.True(b, anterior <= arr[i])
		anterior = arr[i]
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenHeapSort(b, n)
			}
		})
	}
}
