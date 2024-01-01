package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	cola.panicColaVacia()
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := nodoCrear(dato)
	if !cola.EstaVacia() {
		cola.ultimo.prox = nuevoNodo
	} else {
		cola.primero = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.panicColaVacia()
	dato := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.EstaVacia() {
		cola.ultimo = nil
	}
	return dato
}

func (cola colaEnlazada[T]) panicColaVacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}
