package cola_prioridad

type funcCmp[K comparable] func(K, K) int

type heap[K comparable] struct {
	datos    []K
	cantidad int
	cmp      funcCmp[K]
}

const (
	CAPACIDAD_INICIAL        = 10
	FACTOR_AUMENTO           = 2
	FACTOR_REDUCIR           = 2
	LIMITE_REDUCIR_CAPACIDAD = 4
)

// HERRAMIENTAS

func panicColaVacia[K comparable](heap *heap[K]) {
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

func compararDosElementos[K comparable](arr []K, posicion1, posicion2 int, cmp funcCmp[K]) bool {
	diferencia := cmp(arr[posicion1], arr[posicion2])
	return diferencia > 0
}

// devuelve la posicion en donde se encuentre el maximo de las n_posiciones dadas

func max[K comparable](arr []K, posiciones []int, largo int, cmp funcCmp[K]) int {
	max := 0
	for i := 1; i < len(posiciones); i++ {
		if largo > posiciones[i] && compararDosElementos(arr, posiciones[i], posiciones[max], cmp) {
			max = i
		}
	}
	return posiciones[max]
}

func swap[K comparable](datos []K, posicion1, posicion2 int) {
	datos[posicion1], datos[posicion2] = datos[posicion2], datos[posicion1]
}

func upHeap[K comparable](arr []K, posicion int, cmp funcCmp[K]) {
	posicionPadre := posicionPadre(posicion)
	if posicionPadre < 0 {
		return
	}
	if compararDosElementos(arr, posicion, posicionPadre, cmp) {
		swap(arr, posicion, posicionPadre)
		upHeap(arr, posicionPadre, cmp)
	}
}

func downHeap[K comparable](arr []K, posicion, largo int, cmp funcCmp[K]) {
	posIzq := posicionPrimerHijo(posicion)
	posDer := posIzq + 1
	posiciones := []int{posicion, posIzq, posDer}
	maximo := max(arr, posiciones, largo, cmp)
	if arr[maximo] != arr[posicion] {
		swap(arr, maximo, posicion)
		downHeap(arr, maximo, largo, cmp)
	}
}

func (heap *heap[K]) redimensionarHeap(nuevaCapacidad int) {
	nuevoHeap := make([]K, nuevaCapacidad)
	copy(nuevoHeap, heap.datos)
	heap.datos = nuevoHeap
}

func heapify[K comparable](arr []K, largo int, cmp funcCmp[K]) {
	for i := largo - 1; i >= 0; i-- {
		downHeap(arr, i, len(arr), cmp)
	}
}

// FUNCIONES PRINCIPALES

func CrearHeap[K comparable](funcion_cmp func(K, K) int) ColaPrioridad[K] {
	heap := new(heap[K])
	heap.datos = make([]K, CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[K comparable](arreglo []K, funcion_cmp func(K, K) int) ColaPrioridad[K] {
	heap := new(heap[K])
	heap.cmp = funcion_cmp
	heap.datos = make([]K, len(arreglo))
	heap.cantidad = len(arreglo)
	copy(heap.datos, arreglo)
	heapify(heap.datos, heap.cantidad, heap.cmp)
	return heap
}

func (heap *heap[K]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[K]) Encolar(elem K) {
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

func (heap *heap[K]) VerMax() K {
	panicColaVacia[K](heap)
	return heap.datos[0]
}

func (heap *heap[K]) Desencolar() K {
	panicColaVacia[K](heap)
	capacidad := len(heap.datos)
	nuevaCapacidad := capacidad / FACTOR_REDUCIR
	if nuevaCapacidad >= CAPACIDAD_INICIAL && heap.Cantidad()*LIMITE_REDUCIR_CAPACIDAD <= nuevaCapacidad {
		heap.redimensionarHeap(nuevaCapacidad)
	}
	swap(heap.datos, 0, heap.Cantidad()-1)
	dato := heap.datos[heap.Cantidad()-1]
	heap.cantidad--
	downHeap(heap.datos, 0, heap.cantidad, heap.cmp)
	return dato
}

func (heap *heap[K]) Cantidad() int {
	return heap.cantidad
}

func recursivoHeapSort[K comparable](elementos []K, cantidad int, cmp func(K, K) int) {
	if cantidad <= 1 {
		return
	}
	max := 0
	minimo_relativo := cantidad - 1
	swap(elementos, max, minimo_relativo)
	cantidad--
	elementos = elementos[:cantidad]
	downHeap[K](elementos, 0, cantidad, cmp)
	recursivoHeapSort[K](elementos, cantidad, cmp)
}

func HeapSort[K comparable](elementos []K, funcion_cmp func(K, K) int) {
	cantidad := len(elementos)
	heapify[K](elementos, cantidad, funcion_cmp)
	recursivoHeapSort[K](elementos, cantidad, funcion_cmp)
}
