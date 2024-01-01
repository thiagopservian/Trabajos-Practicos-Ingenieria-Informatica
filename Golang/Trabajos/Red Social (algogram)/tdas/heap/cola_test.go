package cola_prioridad_test

import (
	TDAheap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacia(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(i1, i2 int) int { return i1 - i2 })
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola de prioridad esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola de prioridad esta vacia", func() { heap.Desencolar() })
}
