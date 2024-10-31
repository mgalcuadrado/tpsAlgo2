package cola_prioridad_test

import (
	"testing"
	TDAColaPrioridad "tdas/cola_prioridad"
	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC_COLA_VACIA="La cola esta vacia"
	_MENSAJE_TESTING_PANIC_COLA_VACIA="Si la cola esta vacia, debe generarse un panic"
)

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
	for i:=0;i<10;i++{
		cp.Encolar(i)
		require.EqualValues(t,cp.VerMax(),i,"Encolando ciclicamente y ascendentemente, el maximo debe ser el elemento recien insertado.")
	}
	require.EqualValues(t, 10, cp.Cantidad(), "La cola debe tener 9 elementos ingresados")
	require.False(t, cp.EstaVacia(), "Una cola sin elementos debe estar vacia")
	require.EqualValues(t, 9, cp.VerMax(), "El maximo debe ser el elemento mayor ingresado (9)")
	for j:=9;j>=0;j--{
		require.EqualValues(t,cp.Desencolar(),j,"Desencolando ciclicamente y descendentemente, desencolar devuelve al desencolado")
	}
}
