package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	CAPACIDAD_INICIAL        = 10
	FACTOR_AUMENTO           = 2
	FACTOR_REDUCIR           = 2
	LIMITE_REDUCIR_CAPACIDAD = 4
)

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		pila.panicPilaVacia()
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	capacidad := len(pila.datos)
	if pila.cantidad == capacidad {
		nuevaCapacidad := capacidad * FACTOR_AUMENTO
		pila.redimensionarPila(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		pila.panicPilaVacia()
	}
	pila.cantidad--
	anteriorTope := pila.datos[pila.cantidad]

	capacidad := len(pila.datos)
	nuevaCapacidad := capacidad / FACTOR_REDUCIR
	if nuevaCapacidad >= CAPACIDAD_INICIAL && pila.cantidad*LIMITE_REDUCIR_CAPACIDAD <= nuevaCapacidad {
		pila.redimensionarPila(nuevaCapacidad)
	}

	return anteriorTope
}

func (pila *pilaDinamica[T]) redimensionarPila(nuevaCapacidad int) {
	nuevaPila := make([]T, nuevaCapacidad)
	copy(nuevaPila, pila.datos)
	pila.datos = nuevaPila
}

func (pila pilaDinamica[T]) panicPilaVacia() {
	panic("La pila esta vacia")
}
