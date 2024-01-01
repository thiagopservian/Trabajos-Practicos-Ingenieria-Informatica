package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

}

func TestPilaUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(2)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 2, pila.VerTope())
	require.EqualValues(t, 2, pila.Desapilar())
}

func TestVolumen(t *testing.T) {
	tam := 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	for i := 0; i < tam; i++ {
		require.EqualValues(t, tam-i-1, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestRegularesNumeros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 2, pila.VerTope())
}

func TestRegularesStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Hola", pila.VerTope())
	pila.Apilar("Lara")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Lara", pila.VerTope())
	pila.Apilar("Scazzola")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Scazzola", pila.VerTope())
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Lara", pila.VerTope())
}
