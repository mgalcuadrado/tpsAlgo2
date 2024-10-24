package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

/* **************** DEFINICIÓN DE VARIABLES **************** */
const (
	_MENSAJE_PANIC_PILA_VACIA         string = "La pila esta vacia"
	_MENSAJE_TESTING_PANIC_PILA_VACIA string = "No hay elementos en la pila"
)

/* **************** ESTAVACIA() **************** */
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

/* **************** APILAR() **************** */

func TestPilaApilarUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(4)
	require.Equal(t, pila.VerTope(), 4)
	require.False(t, pila.EstaVacia())
}

func TestPilaApilarTresElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(4)
	require.Equal(t, pila.VerTope(), 4)
	require.False(t, pila.EstaVacia())
	pila.Apilar(512)
	require.Equal(t, pila.VerTope(), 512)
	require.False(t, pila.EstaVacia())
	pila.Apilar(2)
	require.Equal(t, pila.VerTope(), 2)
	require.False(t, pila.EstaVacia())
}

/* **************** DESAPILAR() **************** */

func TestPilaDesapilarUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	require.Equal(t, pila.Desapilar(), 3)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

func TestPilaDesapilarTresElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	pila.Apilar(2)
	pila.Apilar(13)
	require.Equal(t, pila.Desapilar(), 13)
	require.False(t, pila.EstaVacia())
	require.Equal(t, pila.Desapilar(), 2)
	require.False(t, pila.EstaVacia())
	require.Equal(t, pila.Desapilar(), 3)
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)

}

/* **************** VERTOPE() **************** */
func TestPilaVerTopeUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	require.Equal(t, pila.VerTope(), 3)
	require.Equal(t, pila.Desapilar(), 3)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

func TestPilaVerTopeTresElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	pila.Apilar(3)
	require.Equal(t, pila.VerTope(), 3)
	require.False(t, pila.EstaVacia())
	pila.Apilar(4)
	require.Equal(t, pila.VerTope(), 4)
	require.False(t, pila.EstaVacia())
	pila.Apilar(6)
	require.Equal(t, pila.VerTope(), 6)
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), 4)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), 3)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

/* **************** PRUEBA CON DISTINTO TIPO DE DATO **************** */

func TestPilaDeFloat64(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	pila.Apilar(3.0)
	require.Equal(t, pila.VerTope(), 3.0)
	require.False(t, pila.EstaVacia())
	pila.Apilar(4.0)
	require.Equal(t, pila.VerTope(), 4.0)
	require.False(t, pila.EstaVacia())
	pila.Apilar(6.0)
	require.Equal(t, pila.VerTope(), 6.0)
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), 4.0)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), 3.0)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

func TestPilaDeStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	pila.Apilar("De aquel amor")
	require.Equal(t, pila.VerTope(), "De aquel amor")
	require.False(t, pila.EstaVacia())
	pila.Apilar("De música ligera")
	require.Equal(t, pila.VerTope(), "De música ligera")
	require.False(t, pila.EstaVacia())
	pila.Apilar("Nada nos libra, nada más queda")
	require.Equal(t, pila.VerTope(), "Nada nos libra, nada más queda")
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), "De música ligera")
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.Equal(t, pila.VerTope(), "De aquel amor")
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}

/* **************** PRUEBAS DE VOLUMEN **************** */

func TestPilaCienElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
		require.Equal(t, pila.VerTope(), i)
	}
	require.False(t, pila.EstaVacia())
	for i := 99; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
		if i > 0 {
			require.False(t, pila.EstaVacia())
		} else {
			require.True(t, pila.EstaVacia())
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
		}
	}
}

func TestPilaMilElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pila.EstaVacia())
	for i := 0.0; i < 1000; i++ {
		pila.Apilar(i / 2.0)
		require.Equal(t, pila.VerTope(), i/2.0)
	}
	require.False(t, pila.EstaVacia())
	for i := 999.0; i >= 0; i-- {
		require.Equal(t, i/2.0, pila.VerTope())
		require.Equal(t, i/2.0, pila.Desapilar())
		if i > 0 {
			require.False(t, pila.EstaVacia())
		} else {
			require.True(t, pila.EstaVacia())
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
		}
	}
}

func TestPilaDiezMilElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pila.EstaVacia())
	for i := 0.0; i < 10000; i++ {
		pila.Apilar(i / 2.0)
		require.Equal(t, pila.VerTope(), i/2.0)
	}
	require.False(t, pila.EstaVacia())
	for i := 9999.0; i >= 0; i-- {
		require.Equal(t, i/2.0, pila.VerTope())
		require.Equal(t, i/2.0, pila.Desapilar())
		if i > 0 {
			require.False(t, pila.EstaVacia())
		} else {
			require.True(t, pila.EstaVacia())
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
			require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
		}
	}
}

/* **************** PRUEBAS DE  USO VARIADO **************** */
func TestPilaUsoVariado(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
		require.Equal(t, pila.VerTope(), i)
	}
	require.False(t, pila.EstaVacia())
	for i := 99; i >= 50; i-- { //desapilo solo 50 elementos
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
		require.False(t, pila.EstaVacia())
	}
	for i := 100; i < 150; i++ { //sumo 50 elementos nuevos
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	//desapilo todos los elementos
	for i := 149; i >= 100; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	for i := 49; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.VerTope() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
	require.PanicsWithValue(t, _MENSAJE_PANIC_PILA_VACIA, func() { pila.Desapilar() }, _MENSAJE_TESTING_PANIC_PILA_VACIA)
}
