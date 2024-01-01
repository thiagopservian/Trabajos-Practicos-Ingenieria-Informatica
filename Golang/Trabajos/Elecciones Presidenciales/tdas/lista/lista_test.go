package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// TESTS LISTA
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

// test usando insertaPrimero
func TestListaUnElementoPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

// test usando InsertarUltimo
func TestListaUnElementoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestVolumenPrimero(t *testing.T) {
	tam := 10000
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, 0, lista.VerUltimo())
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i+1, lista.Largo())
		require.False(t, lista.EstaVacia())
	}
	for i := 0; i < tam-1; i++ {

		require.EqualValues(t, tam-1-i, lista.BorrarPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		require.EqualValues(t, tam-2-i, lista.VerPrimero())
		require.EqualValues(t, tam-1-i, lista.Largo())
		require.False(t, lista.EstaVacia())
	}
}

func TestVolumenUltimo(t *testing.T) {
	tam := 10000
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, 0, lista.VerPrimero())
		require.EqualValues(t, i+1, lista.Largo())
		require.False(t, lista.EstaVacia())
	}
	for i := 0; i < tam-1; i++ {

		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, tam-1, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.VerPrimero())
		require.EqualValues(t, tam-1-i, lista.Largo())
		require.False(t, lista.EstaVacia())
	}

}

func TestRegularNumeros(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.Largo())
}

func TestRegularString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Hola")
	lista.InsertarUltimo("Emiliano")
	lista.InsertarUltimo("Gomez")
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, "Gomez", lista.VerUltimo())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, "Hola", lista.BorrarPrimero())
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, "Gomez", lista.VerUltimo())
	require.EqualValues(t, "Emiliano", lista.VerPrimero())
	require.EqualValues(t, 2, lista.Largo())
}

//TEST ITERADOR EXTERNO
// 1) Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
// 2) Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
// 3) Insertar un elemento en el medio se hace en la posición correcta.
// 4) Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
// 5) Remover el último elemento con el iterador cambia el último de la lista.
// 6) Verificar que al remover un elemento del medio, este no está.

func TestIterVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestIterUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, 1, iterador.Borrar())
	require.True(t, lista.EstaVacia())
}

func TestIterVariosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	iterador := lista.Iterador()
	valores := []int{1, 2, 3, 4, 5}
	for _, valor := range valores {
		require.True(t, iterador.HaySiguiente())
		require.EqualValues(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
	require.False(t, iterador.HaySiguiente())
}

func TestIterVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.True(t, lista.EstaVacia())
	tam := 10000
	for i := 0; i < tam; i++ {
		iterador.Insertar(i)
		require.EqualValues(t, i, iterador.VerActual())
	}
	require.False(t, lista.EstaVacia())
	for i := tam - 1; i <= 0; i-- {
		require.EqualValues(t, i, iterador.VerActual())
		require.True(t, iterador.HaySiguiente())
		iterador.Siguiente()
	}
}

func TestIterVolumenRecorrer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 10000
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	i := 0
	for iter.HaySiguiente() {
		dato := iter.VerActual()
		require.EqualValues(t, i, dato)
		iter.Siguiente()
		i++
	}
}

func TestIterBorrarVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestIterBorrarVarios(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	for i := 0; i < 5; i++ {
		iterador.Insertar(i)
		iterador.Siguiente()
	}
	iterador = lista.Iterador()
	iterador.Siguiente()
	iterador.Borrar()
	iterador.Siguiente()
	iterador.Borrar()
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 4, lista.VerUltimo())
}

func TestIterBorrarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	for i := 0; i < 5; i++ {
		iterador.Insertar(i)
		iterador.Siguiente()
	}
	iterador = lista.Iterador()
	for i := 0; i < 5; i++ {
		iterador.Borrar()
	}
	require.True(t, lista.EstaVacia())
}

func TestIterRecorridoCompleto(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	tam := 5
	for i := 0; i < tam; i++ {
		iterador.Insertar(i)
		iterador.Siguiente()
	}
	iterador = lista.Iterador()
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, iterador.VerActual())
		require.True(t, iterador.HaySiguiente())
		iterador.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.False(t, iterador.HaySiguiente())
}

//TEST ITERADOR INTERNO

/* Hay que probar que funcione
Se puede usar el iterador, por ejemplo, para calcular una suma de todos los elementos en la lista
La prueba NO debe depender de imprimir dentro de la función visitar
Probar el caso de iteración sin condición de corte (iterar toda la lista)
Probar iteración con condición de corte (que en un momento determinado la función visitar dé false)
*/

func TestIterarSumaListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.EqualValues(t, 0, suma)
}

func TestIterarSumaTodosLosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.EqualValues(t, 6, suma)
}

func TestIterarSumaParesHastaPosicion7(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	sumaPares := 0
	lista.Iterar(func(v int) bool {
		if v == 7 {
			return false
		}
		if v%2 == 0 {
			sumaPares += v
		}
		return true
	})
	require.EqualValues(t, 12, sumaPares)
}

func TestIterarBuscarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	elemento := 2
	encontrado := false
	lista.Iterar(func(v int) bool {
		if v == elemento {
			encontrado = true
			return false
		}
		return true
	})
	require.True(t, encontrado)
}
