package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestColaUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerPrimero())
	require.EqualValues(t, 2, cola.Desencolar())
}

func TestVolumen(t *testing.T) {
	tam := 10000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	require.False(t, cola.EstaVacia())
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestRegularesNumeros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerPrimero())
}

func TestRegularesStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Hola", cola.VerPrimero())
	cola.Encolar("Lara")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Hola", cola.VerPrimero())
	cola.Encolar("Scazzola")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Hola", cola.VerPrimero())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Lara", cola.VerPrimero())
}
