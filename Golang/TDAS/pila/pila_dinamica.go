package pila

// CONSTANTES
const (
	CAPACIDAD_BASE = 4
	FACTOR_AUMENTO = 2
	FACTOR_REDUCIR = 2
	LIMITE_REDUCIR = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) redimensionar(nueva_capacidad int) {
	nuevo_array := make([]T, nueva_capacidad)
	copy(nuevo_array, p.datos)
	p.datos = nuevo_array
}

// Apilar implements Pila.
func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == len(p.datos) {
		nueva_capacidad := FACTOR_AUMENTO * len(p.datos)
		p.redimensionar(nueva_capacidad)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

// Desapilar implements Pila.
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	elem := p.datos[p.cantidad-1]
	p.cantidad--
	if p.cantidad*LIMITE_REDUCIR == len(p.datos) {
		nueva_capacidad := len(p.datos) / FACTOR_REDUCIR

		if nueva_capacidad < CAPACIDAD_BASE {
			return elem
		}

		p.redimensionar(nueva_capacidad)
	}

	return elem
}

// EstaVacia implements Pila.
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// VerTope implements Pila.
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CAPACIDAD_BASE), cantidad: 0}
}
