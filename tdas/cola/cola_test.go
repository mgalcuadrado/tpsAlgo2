package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_COLA_VACIA         string = "La cola esta vacia"
	_MENSAJE_TESTING_PANIC_COLA_VACIA string = "No hay elementos en la cola"
)

/* **************** PRUEBA COLA VACIA **************** */

func TestColaEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

/* **************** PRUEBA COLA ENCOLAR **************** */

func TestColaEncolarUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(20)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 20, cola.VerPrimero())
}

func TestColaEncolarTresElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(20)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 20, cola.VerPrimero())
	cola.Encolar(40)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 20, cola.VerPrimero())
	cola.Encolar(60)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 20, cola.VerPrimero())
}

/* **************** PRUEBA COLA VER PRIMERO **************** */

func TestColaVerPrimeroUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(0)
	require.Equal(t, 0, cola.VerPrimero())
	require.Equal(t, 0, cola.Desencolar())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

/* **************** PRUEBA COLA DESENCOLAR **************** */

func TestColaDesencolarUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(20)
	require.Equal(t, 20, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

func TestColaDesencolarTresElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(20)
	cola.Encolar(40)
	cola.Encolar(60)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 20, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 40, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 60, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

/* **************** PRUEBA COLA ENCOLAR Y DESENCOLAR INTERCALADAMENTE **************** */

func TestColaEncolarYDesencolarDesordenado(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	for i := 0; i < 10; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i < 5; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
	for i := 10; i < 20; i++ {
		cola.Encolar(i)
		require.Equal(t, 5, cola.VerPrimero())
	}
	for i := 5; i < 20; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

/* **************** PRUEBA CON DISTINTOS TIPOS DE DATO **************** */

func TestColaDeFloat64(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	require.True(t, cola.EstaVacia())
	for i := 0.0; i < 10.0; i++ {
		cola.Encolar(i)
		require.Equal(t, 0.0, cola.VerPrimero())
	}
	for i := 0.0; i < 5.0; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
	for i := 10.0; i < 20.0; i++ {
		cola.Encolar(i)
		require.Equal(t, 5.0, cola.VerPrimero())
	}
	for i := 5.0; i < 20.0; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

func TestColaDeStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola ")
	cola.Encolar("¿Todo ")
	cola.Encolar("Bien?")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "Hola ", cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, "¿Todo ", cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, "Bien?", cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

func TestColaDePunterosAEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[*int]()
	var a, b, c int = 10, 20, 30
	cola.Encolar(&a)
	cola.Encolar(&b)
	cola.Encolar(&c)
	require.False(t, cola.EstaVacia())
	require.Equal(t, &a, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, &b, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, &c, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

/* **************** PRUEBA DE VOLUMEN**************** */

func TestColaMilElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i < 1000; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}

func TestColaCienMilElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for i := 0; i < 100000; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.False(t, cola.EstaVacia())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.Desencolar() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_COLA_VACIA, func() { cola.VerPrimero() }, _MENSAJE_TESTING_PANIC_COLA_VACIA)
}
