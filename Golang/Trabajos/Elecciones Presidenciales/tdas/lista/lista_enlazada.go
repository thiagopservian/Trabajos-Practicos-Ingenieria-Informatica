package lista

// STRUCS
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *ListaEnlazada[T]
}

// FUNCIONES CREAR
func crearNodolista[T any](elem T) *nodoLista[T] {
	return &nodoLista[T]{dato: elem, siguiente: nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &ListaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

// FUNCIONES LISTA
func (lista ListaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *ListaEnlazada[T]) InsertarPrimero(elem T) {
	nuevo_nodo := crearNodolista[T](elem)
	nuevo_nodo.siguiente = lista.primero
	lista.primero = nuevo_nodo
	if lista.EstaVacia() {
		lista.ultimo = nuevo_nodo
	}
	lista.largo += 1
}

func (lista *ListaEnlazada[T]) InsertarUltimo(elem T) {
	nuevo_nodo := crearNodolista[T](elem)
	if lista.EstaVacia() {
		lista.primero = nuevo_nodo

	} else { //Si no esta vacia
		lista.ultimo.siguiente = nuevo_nodo
	}
	lista.ultimo = nuevo_nodo //se hace siempre
	lista.largo += 1
}
func (lista *ListaEnlazada[T]) BorrarPrimero() T {
	lista.panicListaVacia()
	elem := lista.VerPrimero()
	lista.primero = lista.primero.siguiente
	lista.largo -= 1
	return elem
}
func (lista ListaEnlazada[T]) VerPrimero() T {
	lista.panicListaVacia()
	return lista.primero.dato
}
func (lista ListaEnlazada[T]) VerUltimo() T {
	lista.panicListaVacia()
	return lista.ultimo.dato
}
func (lista ListaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista ListaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.siguiente
	}
}

func (lista *ListaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, anterior: nil, lista: lista}
}

func (lista ListaEnlazada[T]) panicListaVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

// FUNCIONES ITERADOR EXTERNO

func (iter iterListaEnlazada[T]) VerActual() T {
	iter.panicIteradorTerminado()
	return iter.actual.dato
}

func (iter iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}
func (iter *iterListaEnlazada[T]) Siguiente() {
	iter.panicIteradorTerminado()
	iter.anterior, iter.actual = iter.actual, iter.actual.siguiente
}
func (iter *iterListaEnlazada[T]) Insertar(elem T) {
	nodo_nuevo := crearNodolista[T](elem)
	nodo_nuevo.siguiente = iter.actual
	if iter.anterior != nil { // no se inserta al principio
		iter.anterior.siguiente = nodo_nuevo
	} else { //se inserta al principio
		iter.lista.primero = nodo_nuevo
	}
	if iter.actual == nil { //insertamos al final
		iter.lista.ultimo = nodo_nuevo
	}
	iter.actual = nodo_nuevo
	iter.lista.largo += 1
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	iter.panicIteradorTerminado()
	elem := iter.actual.dato
	if iter.anterior != nil { //si NO estamos al principio (NO insertamos en el primer lugar)
		iter.anterior.siguiente = iter.actual.siguiente
	} else {
		iter.lista.primero = iter.actual.siguiente
	}
	iter.actual = iter.actual.siguiente
	if iter.actual == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.lista.largo -= 1
	return elem
}

func (iter iterListaEnlazada[T]) panicIteradorTerminado() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
