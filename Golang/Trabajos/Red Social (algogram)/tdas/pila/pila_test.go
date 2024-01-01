package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaConUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(5)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 5, pila.VerTope())
	require.EqualValues(t, 5, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	tam := 10000
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := tam - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	frase := []string{"Hola", "como", "estas", "?"}
	for _, palabra := range frase {
		pila.Apilar(palabra)
		require.EqualValues(t, palabra, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := len(frase) - 1; i >= 0; i-- {
		require.EqualValues(t, frase[i], pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaBool(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	sliceDeBool := []bool{true, true, false, true, false}
	for _, palabra := range sliceDeBool {
		pila.Apilar(palabra)
		require.EqualValues(t, palabra, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := len(sliceDeBool) - 1; i >= 0; i-- {
		require.EqualValues(t, sliceDeBool[i], pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaFloat64(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	numeros := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	for _, num := range numeros {
		pila.Apilar(num)
		require.EqualValues(t, num, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := len(numeros) - 1; i >= 0; i-- {
		require.EqualValues(t, numeros[i], pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

type structEj struct {
	miembro1 int
	miembro2 string
}

func TestPilaEstructura(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[structEj]()
	estructuras := []structEj{
		{miembro1: 1, miembro2: "Uno"},
		{miembro1: 2, miembro2: "Dos"},
		{miembro1: 3, miembro2: "Tres"},
	}
	for _, elem := range estructuras {
		pila.Apilar(elem)
		require.EqualValues(t, elem, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := len(estructuras) - 1; i >= 0; i-- {
		require.EqualValues(t, estructuras[i], pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}
