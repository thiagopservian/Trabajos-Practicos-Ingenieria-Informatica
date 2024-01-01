package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

// Devuelve si hay o no hay solo un elemento en la cola
func (cola *colaEnlazada[T]) hay_un_elemento() bool {
	return cola.primero == cola.ultimo
}

// Desencolar implements Cola.
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	elem := cola.primero.dato
	if cola.hay_un_elemento() {
		cola.ultimo = nil
	}
	cola.primero = cola.primero.prox
	return elem
}

// Encolar implements Cola.
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevo_nodo := crearnodocola[T](elem)
	if cola.EstaVacia() {
		cola.primero = nuevo_nodo
	} else { //si no esta vacia
		cola.ultimo.prox = nuevo_nodo
	}
	cola.ultimo = nuevo_nodo
}

// EstaVacia implements Cola.
func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

// VerPrimero implements Cola.
func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

// Crea una cola enlazada vacia
func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

// Crea un nodo con el elemento recibido
func crearnodocola[T any](elem T) *nodoCola[T] {
	return &nodoCola[T]{dato: elem, prox: nil}
}
