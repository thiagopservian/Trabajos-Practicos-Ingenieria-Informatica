package cola_prioridad_test

import (
	"strings"
	TDAheap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestColaUnElemento(t *testing.T) {
	cola := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
	require.EqualValues(t, 2, cola.Desencolar())
}

func TestDesencolarUnelemento(t *testing.T) {
	cola := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	cola.Encolar(2)
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 0, cola.Cantidad())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestVolumen(t *testing.T) {
	tam := 10000
	cola := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		require.EqualValues(t, i, cola.VerMax())
	}
	require.False(t, cola.EstaVacia())
	for i := tam - 1; i >= 0; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestRegularesNumeros(t *testing.T) {
	cola := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerMax())
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
	cola.Encolar(3)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 3, cola.VerMax())
	require.EqualValues(t, 3, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
}

func TestRegularesStrings(t *testing.T) {
	cola := TDAheap.CrearHeap[string](func(i1, i2 string) int { return strings.Compare(i1, i2) })
	cola.Encolar("a")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "a", cola.VerMax())
	cola.Encolar("b")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "b", cola.VerMax())
	cola.Encolar("c")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "c", cola.VerMax())
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "b", cola.VerMax())
}

func TestColaVaciaconArr(t *testing.T) {
	arr := []int{}
	heap := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestColaUnElementoarr(t *testing.T) {
	arr := []int{2}
	cola := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
	require.EqualValues(t, 2, cola.Desencolar())
}

func TestColaconArr(t *testing.T) {
	arr := []int{}
	cola := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
	require.EqualValues(t, 2, cola.Desencolar())
}

func TestDesencolarUnelementoArr(t *testing.T) {
	arr := []int{2}
	cola := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 0, cola.Cantidad())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestVolumenArr(t *testing.T) {
	tam := 1000
	arr := make([]int, tam)
	for i := 0; i < tam; i++ {
		arr[i] = i
	}
	cola := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, tam-1, cola.VerMax())
	require.EqualValues(t, tam, cola.Cantidad())
	for i := tam - 1; i >= 0; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.EqualValues(t, 0, cola.Cantidad())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestRegularesNumerosArr(t *testing.T) {
	arr := []int{1, 2, 3}
	cola := TDAheap.CrearHeapArr[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 3, cola.VerMax())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 3, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2, cola.VerMax())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 1, cola.VerMax())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestHeapsortVacio(t *testing.T) {
	arr := []int{}
	TDAheap.HeapSort[int](arr, func(i1, i2 int) int { return i1 - i2 })
	sliceVacio := []int{}
	require.EqualValues(t, arr, sliceVacio)
}

func TestHeapsortUnElemento(t *testing.T) {
	arr := []int{1}
	TDAheap.HeapSort[int](arr, func(i1, i2 int) int { return i1 - i2 })
	slice := []int{1}
	require.EqualValues(t, arr, slice)
}

func TestHeapsortRegular(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	TDAheap.HeapSort[int](arr, func(i1, i2 int) int { return i1 - i2 })
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.EqualValues(t, arr, slice)
}

func TestHeapsortVolumen(t *testing.T) {
	tam := 1000
	arr := make([]int, tam)
	sliceOrd := make([]int, tam)
	for i := 0; i < tam; i++ {
		sliceOrd[i] = i
	}
	for i := tam - 1; i >= 0; i-- {
		arr[i] = i
	}
	TDAheap.HeapSort[int](arr, func(i1, i2 int) int { return i1 - i2 })
	require.EqualValues(t, arr, sliceOrd)
}
