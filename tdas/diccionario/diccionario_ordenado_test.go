package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN2 = []int{20, 12500, 25000, 50000, 100000}

func comparacion2(v1 int, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	} else if v1 < v2 {
		return v1 - v2
	} else {
		return 0
	}
}

func buscar2(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

/* ****************** Tests ABB Ordenado ************** */

func TestABBBorrarMultiplesElementos(t *testing.T) {
	t.Log("Esta prueba verifica que se borren múltiples elementos correctamente y se mantengan todos los demás en el ABB")
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 3, 22, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	require.Equal(t, len(arr), dic.Cantidad(), "Inicialmente se guardan todos los elementos en el diccionario")
	dic.Borrar(50) //caso 1: borrar un elemento sin hijos
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(50) }, "El elemento borrado no debería estar en el ABB")

	dic.Borrar(12) //caso 2: borrar un elemento con un hijo
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(12) }, "El elemento borrado no debería estar en el ABB")
	require.True(t, dic.Pertenece(4), "Al borrar un elemento con un hijo su hijo debería seguir perteneciendo al diccionario")

	dic.Borrar(23) //caso 3: borrar un elemento con dos hijos
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(23) }, "El elemento borrado no debería estar en el ABB")
	require.True(t, dic.Pertenece(22), "Al borrar un elemento con un hijo su hijo debería seguir perteneciendo al diccionario")
	require.True(t, dic.Pertenece(56), "Al borrar un elemento con un hijo su hijo debería seguir perteneciendo al diccionario")

	require.Equal(t, len(arr)-3, dic.Cantidad(), "Luego de borrar elementos la cantidad de elementos del diccionario se modifica")
	elementosEnABB := []int{14, 2, 3, 22, 56, 4}
	for i := 0; i < len(elementosEnABB); i++ {
		require.True(t, dic.Pertenece(elementosEnABB[i]), "El elemento no borrado debería estar en el ABB")
	}
}

func TestABBBorrarRaizConDosHijos(t *testing.T) {
	t.Log("Esta prueba verifica que se borre la raíz del abb con dos hijos correctamente y se mantengan todos los demás elementos del ABB")
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 3, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	dic.Borrar(14) //se borra la raíz
	for i := 0; i < len(arr); i++ {
		if arr[i] != 14 {
			require.True(t, dic.Pertenece(arr[i]), "El elemento no borrado debería estar en el ABB")
		} else {
			require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(14) }, "El elemento borrado no debería estar en el ABB")
		}
	}
}

/* ****************** Iterador interno ************** */

func TestIteradorInternoOrdenado(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 3, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	arr_ordenado := []int{2, 3, 4, 12, 14, 23, 50, 56}
	contador := 0
	dic.Iterar(func(clave int, dato int) bool {
		if clave != arr_ordenado[contador] {
			return false
		}
		contador++
		return true
	})
	require.Equal(t, contador, len(arr), "El iterador interno recorre in order")
}

func TestIteradorInternoOrdenadoTrasBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 15, 23, 2, 3, 56, 12, 50, 4, 22}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	dic.Borrar(3)
	dic.Borrar(15)
	dic.Borrar(22)
	arr_ordenado := []int{2, 4, 12, 14, 23, 50, 56}
	contador := 0
	dic.Iterar(func(clave int, dato int) bool {
		if clave != arr_ordenado[contador] {
			return false
		}
		contador++
		return true
	})
	require.Equal(t, contador, len(arr_ordenado), "El iterador interno recorre in order todos los elementos del diccionario")
}

/* ******************* Iterador interno rangos ******************* */

func TestIterarRangoVariantesDe7(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{3, 1, 2, 4, 5, 6, 7}
	for i, _ := range arr {
		dic.Guardar(arr[i], arr[i])
	}
	inicio, fin := 2, 5
	contador := 0
	dic.IterarRango(&inicio, &fin, func(clave int, dato int) bool {
		require.Equal(t, contador+2, clave)
		contador++
		return true
	})
	require.Equal(t, 4, contador)
}

func TestVolumenIteradorRangoCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	/* Inserta 'n' parejas en el abb */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}
	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	desde, hasta := 2500, 7500
	dic.IterarRango(&desde, &hasta, func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

/* ******************* Iterador externo ******************* */

func TestIteradorExternoOrdenado(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 3, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	arr_ordenado := []int{2, 3, 4, 12, 14, 23, 50, 56}
	contador := 0
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		primero, _ := iter.VerActual()
		require.Equal(t, primero, arr_ordenado[contador])
		contador++
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	require.Equal(t, contador, len(arr_ordenado), "El iterador interno recorre in order")
}

func TestIteradorExternoOrdenadoTrasBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 3, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	arr_ordenado := []int{2, 3, 4, 12, 14, 23, 50, 56}
	contador := 0
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		primero, _ := iter.VerActual()
		require.Equal(t, primero, arr_ordenado[contador])
		contador++
		iter.Siguiente()
	}
	require.Equal(t, contador, len(arr), "El iterador interno recorre in order")
}

/* ******************* Iterador externo rangos ******************* */
func TestIteradorRangoDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	strings := []string{"Hola", "Chau"}
	iter := dic.IteradorRango(&strings[0], &strings[1])
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoDiccionarioClavesYValores(t *testing.T) {
	t.Log("Se guardan claves de tipo int y valores de tipo string y se itera con el iterador externo en un rango de valores. El test verifica que se itere in order")
	dic := TDADiccionario.CrearABB[int, string](comparacion2)
	arr := []int{3, 15, 2, 12, 54, 20, 37, 33, 16, 7}
	arr_ordenado := []int{2, 3, 7, 12, 15, 16, 20, 33, 37, 54}
	str := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	str_ordenado := []string{"C", "A", "J", "D", "B", "I", "F", "H", "G", "E"}
	for indice, valor := range arr {
		dic.Guardar(valor, str[indice])
	}
	iter := dic.IteradorRango(&arr[3], &arr[5])
	contador := 0
	for ; iter.HaySiguiente(); contador++ {
		clave, valor := iter.VerActual()
		require.Equal(t, arr_ordenado[3+contador], clave)
		require.Equal(t, str_ordenado[3+contador], valor)
		iter.Siguiente()
	}
	require.Equal(t, 4, contador) //verifica que se corte luego de arr[5]
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.IteradorRango(&claves[0], &claves[1])
	iter2 := dic.IteradorRango(&claves[0], &claves[2])
	cero, _ := iter2.VerActual()
	iter2.Siguiente()
	iter3 := dic.IteradorRango(&claves[1], &claves[2])
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, -1, buscar2(primero, claves))
	require.NotEqualValues(t, -1, buscar2(segundo, claves))
	require.NotEqualValues(t, -1, buscar2(cero, claves))
}

func TestPruebaIteradorRangoTrasBorrados(t *testing.T) {
	t.Log("Esta prueba intenta verificar el comportamiento del ABB cuando " +
		"se eliminan elementos. El iterador no debería iterar sobre elementos eliminados")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.IteradorRango(&clave1, &clave2)

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.IteradorRango(&clave1, &clave2)

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIteradorRango(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[int, *int](comparacion2)

	claves := make([]int, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el abb */
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			claves[i] = -i
			valores[i] = -i
		} else {
			claves[i] = i
			valores[i] = i
		}
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	clave1 := 0
	clave2 := n
	iter := dic.IteradorRango(&clave1, &clave2)
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	contador := 0

	for i = 0; i < n/2; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		if c1 > n {
			ok = false
			break
		}
		if *v1 < -n || *v1 > n {
			ok = false
			break
		}
		if *v1 >= 0 {
			contador++
		}
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, (n / 2), i, "No se recorrió todo el rango")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")
	require.EqualValues(b, (n / 2), contador)
}

func BenchmarkIteradorRango(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) n elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN2 {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorRango(b, n)
			}
		})
	}
}

func TestIteradorRangoVariantesDe7(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{3, 1, 2, 4, 5, 6, 7}
	for i, _ := range arr {
		dic.Guardar(arr[i], arr[i])
	}
	inicio, fin := 2, 5
	contador := 0
	iter := dic.IteradorRango(&inicio, &fin)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.Equal(t, contador+2, clave)
		contador++
		iter.Siguiente()
	}
	require.Equal(t, 4, contador)
}

func TestIteradorRangoCasoBorde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{10, 5, 7, 9}
	for i, _ := range arr {
		dic.Guardar(arr[i], arr[i])
	}
	inicio := 8
	contador := 0
	iter := dic.IteradorRango(&inicio, nil)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.Equal(t, contador+9, clave)
		contador++
		iter.Siguiente()
	}
	require.Equal(t, 2, contador)
}

func TestIterarInternoConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparacion2)
	arr := []int{14, 23, 2, 56, 12, 50, 4}
	for indice, valor := range arr {
		dic.Guardar(arr[indice], valor)
	}
	arr_ordenado := []int{2, 4, 12, 14, 23, 50, 56}
	contador := 0
	dic.Iterar(func(clave int, dato int) bool {
		require.Equal(t, arr_ordenado[contador], clave)
		contador++
		return clave%2 == 0
	})
	require.Equal(t, contador, 5)
}
