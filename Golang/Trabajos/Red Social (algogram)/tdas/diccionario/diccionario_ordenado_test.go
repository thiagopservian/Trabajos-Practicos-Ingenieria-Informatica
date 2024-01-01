package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(1) })
}

// guardar

func TestOrdenadoUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(1, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(2))
	require.EqualValues(t, 10, dic.Obtener(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })
}

func TestOrdenadoUnElement2(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestOrdenadoReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	t.Log(dic.Cantidad())
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	t.Log(dic.Cantidad())
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestOrdenadoConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestOrdenadoClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestOrdenadoValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestOrdenadoCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

// borrar

func TestOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestOrdenadoGuardarYBorrarRepetidasVeces(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestOrdenadoAgregarunBorrado(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(8, 8)
	dic.Guardar(12, 12)
	dic.Guardar(11, 11)
	dic.Guardar(9, 9)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.False(t, dic.Pertenece(10))
	dic.Guardar(6, 6)
	dic.Guardar(10, 10)
	require.EqualValues(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.EqualValues(t, 10, dic.Obtener(10))
	require.EqualValues(t, 10, dic.Borrar(10))
}

func TestOrdenadoBorrarRaizunElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestOrdenadoBorrarRaizUnHijo(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(11, 11)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 1, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestOrdenadoBorrarRaizUnHijoyNietoDer(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(11, 11)
	dic.Guardar(12, 12)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 2, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestOrdenadoBorrarRaizUnHijoyNietoIzq(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(12, 12)
	dic.Guardar(11, 11)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 2, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestOrdenadoBorrarRaizDosHijos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(12, 12)
	dic.Guardar(9, 9)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 2, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
}

func TestOrdenadoBorrarCaso2HijosRaiz(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(12, 12)
	dic.Guardar(20, 20)
	require.EqualValues(t, 10, dic.Borrar(10))
	require.EqualValues(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(8))
	require.True(t, dic.Pertenece(12))
	require.True(t, dic.Pertenece(20))
}

func TestOrdenadoBorrarCaso2HijosReemplazantesProfundo(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(12, 12)
	dic.Guardar(20, 20)
	require.EqualValues(t, 5, dic.Borrar(5))
	require.EqualValues(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(3))
	require.True(t, dic.Pertenece(8))
	require.True(t, dic.Pertenece(12))
	require.True(t, dic.Pertenece(20))
}

func TestOrdenadoBorrarCaso2HijosReemplazanteProfundoCon1Hijo(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	dic.Guardar(10, 10)
	dic.Guardar(5, 5)
	dic.Guardar(15, 15)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(12, 12)
	dic.Guardar(20, 20)
	require.EqualValues(t, 5, dic.Borrar(5))
	require.EqualValues(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(3))
	require.True(t, dic.Pertenece(8))
	require.True(t, dic.Pertenece(12))
	require.True(t, dic.Pertenece(20))
}

func TestOrdenadoBorradosNoGeneraStackOverflow(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })
	for i := 1; i <= 1000; i++ {
		dic.Guardar(i, i)
	}
	for i := 1; i <= 1000; i++ {
		require.EqualValues(t, i, dic.Borrar(i))
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func buscarIterador2(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

// ITERADORES

// interno

func TestOrdenadoIterarABBVacio(t *testing.T) {
	t.Log("Iterar sobre un ABB vacío debería finalizar inmediatamente")
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Iterar(func(clave string, dato int) bool {
		return true
	})
}

func TestOrdenadoIterarABBVacioConRango(t *testing.T) {
	t.Log("Iterar sobre un ABB vacío con rango debería finalizar inmediatamente")
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	desde := "A"
	hasta := "Z"
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		return true
	})
}

func TestOrdenadoIterarABBUnElement(t *testing.T) {
	t.Log("Iterar sobre un ABB con un solo elemento debería funcionar")
	clave := "Gato"
	valor := 42
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave, valor)
	dic.Iterar(func(clave string, dato int) bool {
		require.EqualValues(t, clave, "Gato")
		require.EqualValues(t, dato, 42)
		return true
	})
}

func TestOrdenadoIterarClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarIterador2(cs[0], claves))
	require.NotEqualValues(t, -1, buscarIterador2(cs[1], claves))
	require.NotEqualValues(t, -1, buscarIterador2(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestOrdenadoIterarValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestOrdenadoIterarFueraDeRango(t *testing.T) {
	t.Log("Iterar fuera del rango debería finalizar inmediatamente")
	clave1 := "A"
	clave2 := "B"
	clave3 := "C"
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave1, 1)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)

	desde := "D"
	hasta := "Z"
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		t.Errorf("No debería haber iteración fuera del rango")
		return true
	})
}

func TestOrdenadoIterarConRangoCruzado(t *testing.T) {
	t.Log("Iterar con rango cruzado (desde > hasta) debería finalizar inmediatamente")
	clave1 := "A"
	clave2 := "B"
	clave3 := "C"
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave1, 1)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)

	desde := "Z"
	hasta := "A"
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		t.Errorf("No debería haber iteración con rango cruzado")
		return true
	})
}

func TestOrdenadoIterarEnRango(t *testing.T) {
	t.Log("Iterar con un rango de elementos aceptable debería funcionar")
	clave1 := "A"
	clave2 := "B"
	clave3 := "C"
	clave4 := "D"
	clave5 := "E"
	valor1 := 1
	valor2 := 2
	valor3 := 3
	valor4 := 4
	valor5 := 5
	dic := TDADiccionario.CrearABB[string, int](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Guardar(clave4, valor4)
	dic.Guardar(clave5, valor5)

	desde := "B"
	hasta := "D"
	elementosDentroDelRango := make(map[string]int)
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		elementosDentroDelRango[clave] = dato
		return true
	})

	require.EqualValues(t, 3, len(elementosDentroDelRango))
	require.EqualValues(t, 2, elementosDentroDelRango["B"])
	require.EqualValues(t, 3, elementosDentroDelRango["C"])
	require.EqualValues(t, 4, elementosDentroDelRango["D"])
}

func TestOrdenadoIterarCondicionDeCorte(t *testing.T) {
	t.Log("Iterar se detiene cuando la función visitar devuelve false")

	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })

	dic.Guardar("2", "dos")
	dic.Guardar("3", "tres")
	dic.Guardar("4", "cuatro")
	dic.Guardar("1", "uno")
	pararEnClave := "3"
	visitar := func(clave string, valor string) bool {
		t.Logf("Visitando elemento: clave=%s, valor=%s", clave, valor)
		if clave == pararEnClave {
			t.Log("Deteniendo la iteración en clave:", clave)
			return false
		}
		return true

	}

	resultados := make(map[string]string)
	dic.IterarRango(nil, nil, func(clave string, valor string) bool {
		resultados[clave] = valor
		return visitar(clave, valor)
	})

	expected := map[string]string{"1": "uno", "2": "dos", "3": "tres"}
	require.EqualValues(t, expected, resultados, "El resultado de la iteración no coincide con lo esperado")
}

func TestOrdenadoIterarVolumenCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](func(i1, i2 int) int { return i1 - i2 })

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
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

// externo

func TestOrdenadoIteradorVacio(t *testing.T) {
	t.Log("Iterar sobre un ABB vacío debe devolver false en HaySiguiente")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	iterador := abb.Iterador()
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestOrdenadoIteradorVacioConRango(t *testing.T) {
	t.Log("Iterar sobre ABB vacío con rango utilizando el iterador debe devolver false en HaySiguiente")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	rangoDesde := 0
	rangoHasta := 10
	iterador := abb.IteradorRango(&rangoDesde, &rangoHasta)
	require.False(t, iterador.HaySiguiente())
}

func TestOrdenadoIteradorUnElemento(t *testing.T) {
	t.Log("Iterar sobre un ABB con un solo elemento debe devolver true en HaySiguiente y el elemento correcto en VerActual")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	abb.Guardar(42, "respuesta")
	iterador := abb.Iterador()
	require.True(t, iterador.HaySiguiente())
	clave, valor := iterador.VerActual()
	require.Equal(t, 42, clave)
	require.Equal(t, "respuesta", valor)
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
}

func TestOrdenadoIteradorFueraDeRango(t *testing.T) {
	t.Log("Iterar fuera del rango utilizando el iterador debe hacer que el iterador esté al final")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")
	rangoDesde := 4
	rangoHasta := 5
	iterador := abb.IteradorRango(&rangoDesde, &rangoHasta)
	require.False(t, iterador.HaySiguiente())
}

func TestOrdenadoIteradorConRangoCruzado(t *testing.T) {
	t.Log("Iterar con rango cruzado (desde > hasta) utilizando el iterador debe hacer que el iterador esté al final")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")
	rangoDesde := 3
	rangoHasta := 2
	iterador := abb.IteradorRango(&rangoDesde, &rangoHasta)
	require.False(t, iterador.HaySiguiente())
}

func TestOrdenadoIteradorEnRango(t *testing.T) {
	t.Log("Iterar en un rango de elementos aceptable utilizando el iterador")
	abb := TDADiccionario.CrearABB[int, string](func(i1, i2 int) int { return i1 - i2 })
	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")
	rangoDesde := 1
	rangoHasta := 3
	iterador := abb.IteradorRango(&rangoDesde, &rangoHasta)

	require.True(t, iterador.HaySiguiente())
	clave, valor := iterador.VerActual()
	require.Equal(t, 1, clave)
	require.Equal(t, "uno", valor)

	iterador.Siguiente()

	require.True(t, iterador.HaySiguiente())
	clave, valor = iterador.VerActual()
	require.Equal(t, 2, clave)
	require.Equal(t, "dos", valor)

	iterador.Siguiente()

	require.True(t, iterador.HaySiguiente())
	clave, valor = iterador.VerActual()
	require.Equal(t, 3, clave)
	require.Equal(t, "tres", valor)

	iterador.Siguiente()

	require.False(t, iterador.HaySiguiente())
}

func TestOrdenadoIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarIterador2(primero, claves))
	require.NotEqualValues(t, -1, buscarIterador2(segundo, claves))
	require.NotEqualValues(t, -1, buscarIterador2(tercero, claves))
}
