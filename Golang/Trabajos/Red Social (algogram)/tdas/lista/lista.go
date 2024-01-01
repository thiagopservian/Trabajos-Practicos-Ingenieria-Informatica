package lista

type Lista[T any] interface {

	//Esta vacia devuelve si la lista no tiene elementos, false en caso contrario
	EstaVacia() bool

	//InsertarPrimero agrega un nuevo elemento a la lista, al final de la misma.
	InsertarPrimero(T)

	//InsertarUltimo agrega un nuevo elemento a la lista, al principio de la misma
	InsertarUltimo(T)

	//BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	//Largo devuelve el numero de elementos que hay en la lista
	Largo() int

	//Itera hasta donde la función dada lo defina, iterando de manera interna
	Iterar(visitar func(T) bool)

	//Crea un Iterador Externo en base a la Lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//VerActual devuelve el elemento en la posicion de la iteracion actual
	VerActual() T

	//HaySiguiente devuelve si hay un elemento en la posicion actual
	HaySiguiente() bool

	//Siguiente avanza a la siguiente iteración
	Siguiente()

	//Insertar agrega un elemento a la lista
	//en la posicion de la iteracion actual
	Insertar(T)

	//Borrar saca el elemento de la lista que esta en la posicion de la
	// iteración actual, y devuelve ese valor
	Borrar() T
}
