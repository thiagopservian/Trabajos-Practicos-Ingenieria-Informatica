package cola_test

import (
	TDAcola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaConUnElemento(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
	require.EqualValues(t, 5, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestVolumen(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	tam := 10000
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaStrings(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[string]()
	frase := []string{"Hola", "como", "estas", "?"}
	for _, palabra := range frase {
		cola.Encolar(palabra)
	}
	require.False(t, cola.EstaVacia())
	for _, palabra := range frase {
		require.EqualValues(t, palabra, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestColaBool(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[bool]()
	sliceDeBool := []bool{true, true, false, true, false}
	for _, booleano := range sliceDeBool {
		cola.Encolar(booleano)
	}
	require.False(t, cola.EstaVacia())
	for _, booleano := range sliceDeBool {
		require.EqualValues(t, booleano, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestColaFloat64(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[float64]()
	numeros := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	for _, num := range numeros {
		cola.Encolar(num)
	}
	require.False(t, cola.EstaVacia())
	for _, num := range numeros {
		require.EqualValues(t, num, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

type structEj struct {
	miembro1 int
	miembro2 string
}

func TestColaEstructura(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[structEj]()
	estructuras := []structEj{
		{miembro1: 1, miembro2: "Uno"},
		{miembro1: 2, miembro2: "Dos"},
		{miembro1: 3, miembro2: "Tres"},
	}
	for _, elem := range estructuras {
		cola.Encolar(elem)
	}
	require.False(t, cola.EstaVacia())
	for _, elem := range estructuras {
		require.EqualValues(t, elem, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}
