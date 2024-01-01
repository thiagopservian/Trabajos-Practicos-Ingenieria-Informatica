package cola_prioridad

type funcCmp[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      funcCmp[T]
}

const (
	CAPACIDAD_INICIAL        = 10
	FACTOR_AUMENTO           = 2
	FACTOR_REDUCIR           = 2
	LIMITE_REDUCIR_CAPACIDAD = 4
	POS_INICIAL              = 0
)

// HERRAMIENTAS

func panicColaVacia[T any](heap *heap[T]) {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func posicionPrimerHijo(posicion int) int {
	return 2*posicion + 1
}

func posicionPadre(posicion int) int {
	return (posicion - 1) / 2
}

// devuelve true si el primero es mas grande que el segundo

func compararDosElementos[T any](arr []T, posicion1, posicion2 int, cmp funcCmp[T]) bool {
	diferencia := cmp(arr[posicion1], arr[posicion2])
	return diferencia > 0
}

// devuelve la posicion en donde se encuentre el maximo de las n_posiciones dadas

func max[T any](arr []T, posiciones []int, largo int, cmp funcCmp[T]) int {
	max := 0
	for i := 1; i < len(posiciones); i++ {
		if largo > posiciones[i] && compararDosElementos(arr, posiciones[i], posiciones[max], cmp) {
			max = i
		}
	}
	return posiciones[max]
}

func swap[T any](datos []T, posicion1, posicion2 *int) {
	datos[*posicion1], datos[*posicion2] = datos[*posicion2], datos[*posicion1]
}

func upHeap[T any](arr []T, posicion int, cmp funcCmp[T]) {
	posicionPadre := posicionPadre(posicion)
	if posicionPadre < 0 {
		return
	}
	if compararDosElementos(arr, posicion, posicionPadre, cmp) {
		swap(arr, &posicion, &posicionPadre)
		upHeap(arr, posicionPadre, cmp)
	}
}

func downHeap[T any](arr []T, posicion, largo int, cmp funcCmp[T]) {
	posIzq := posicionPrimerHijo(posicion)
	posDer := posIzq + 1
	posiciones := []int{posicion, posIzq, posDer}
	maximo := max(arr, posiciones, largo, cmp)
	dif := cmp(arr[maximo], arr[posicion])
	if dif != 0 {
		swap(arr, &maximo, &posicion)
		downHeap(arr, maximo, largo, cmp)
	}
}

func (heap *heap[T]) redimensionarHeap(nuevaCapacidad int) {
	nuevoHeap := make([]T, nuevaCapacidad)
	copy(nuevoHeap, heap.datos)
	heap.datos = nuevoHeap
}

func heapify[T any](arr []T, largo int, cmp funcCmp[T]) {
	for i := largo - 1; i >= 0; i-- {
		downHeap(arr, i, len(arr), cmp)
	}
}

// FUNCIONES PRINCIPALES

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.datos = make([]T, len(arreglo))
	heap.cantidad = len(arreglo)
	copy(heap.datos, arreglo)
	heapify(heap.datos, heap.cantidad, heap.cmp)
	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	capacidad := len(heap.datos)
	if heap.Cantidad() == capacidad {
		nuevaCapacidad := capacidad * FACTOR_AUMENTO
		if nuevaCapacidad < CAPACIDAD_INICIAL {
			nuevaCapacidad = CAPACIDAD_INICIAL
		}
		heap.redimensionarHeap(nuevaCapacidad)
	}
	heap.datos[heap.Cantidad()] = elem
	upHeap(heap.datos, heap.Cantidad(), heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	panicColaVacia[T](heap)
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	panicColaVacia[T](heap)
	capacidad := len(heap.datos)
	nuevaCapacidad := capacidad / FACTOR_REDUCIR
	if nuevaCapacidad >= CAPACIDAD_INICIAL && heap.Cantidad()*LIMITE_REDUCIR_CAPACIDAD <= nuevaCapacidad {
		heap.redimensionarHeap(nuevaCapacidad)
	}
	posInicial, posFinal := 0, heap.Cantidad()-1
	swap(heap.datos, &posInicial, &posFinal)
	dato := heap.datos[posFinal]
	heap.cantidad--
	downHeap(heap.datos, 0, heap.cantidad, heap.cmp)
	return dato
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	cantidad := len(elementos)
	heapify[T](elementos, cantidad, funcion_cmp)
	for i := cantidad - 1; i > 0; i-- {
		max := 0
		minimo_relativo := i
		swap(elementos, &max, &minimo_relativo)
		cantidad--
		elementos = elementos[:cantidad]
		downHeap[T](elementos, 0, cantidad, funcion_cmp)
	}
}
