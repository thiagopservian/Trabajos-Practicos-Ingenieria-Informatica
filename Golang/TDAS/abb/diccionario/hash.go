package diccionario

import (
	"fmt"
)

// CONSTANTES
type estado int

const (
	TAMAÑO_ORIGINAL  = 11
	FNV_OFFSET_BASIS = 2166136261 //decimal
	FNV_PRIME        = 16777619   //decimal
	FACTOR_AUMENTO   = 2
	FACTOR_REDUCTOR  = 2
	MAXIMA_CARGA     = 0.7
	MINIMA_CARGA     = 0.2
)
const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

//STRUCS

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

type iterador[K comparable, V any] struct {
	posActual  int
	recorridos int
	hash       *hashCerrado[K, V]
}

// FUNCIONES CREAR
func crearTabla[K comparable, V any](tam int) []celdaHash[K, V] {
	tabla := make([]celdaHash[K, V], tam)
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = crearTabla[K, V](TAMAÑO_ORIGINAL)
	hash.tam = TAMAÑO_ORIGINAL
	return hash
}

//HERRAMIENTAS

// FNV decimal
func funcionHashing[K comparable](clave K) int {
	var hash uint32 = FNV_OFFSET_BASIS
	bytes := []byte(fmt.Sprintf("%v", clave))
	for _, bit := range bytes {
		hash ^= uint32(bit)
		hash *= FNV_PRIME
	}
	return int(hash)
}

func (hash hashCerrado[K, V]) carga() float32 {
	ocupados := float32(hash.Cantidad())
	borrados := float32(hash.borrados)
	tamaño := float32(hash.tam)
	return float32((ocupados + borrados) / tamaño)
}

func (hash *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	tablaAnterior := hash.tabla
	hash.tabla = crearTabla[K, V](nuevoTam)
	hash.borrados = 0
	hash.tam = nuevoTam
	hash.cantidad = 0
	for _, celda := range tablaAnterior {
		if celda.estado == OCUPADO {
			hash.Guardar(celda.clave, celda.dato)
		}
	}
}

func (hash hashCerrado[K, V]) recorrer(clave K) int {
	pos := funcionHashing(clave) % hash.tam
	for hash.tabla[pos].estado != VACIO && !(hash.tabla[pos].estado == OCUPADO && hash.tabla[pos].clave == clave) {
		pos = (pos + 1) % hash.tam
	}
	return pos
}

func panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

//PRIMITIVAS DICCIONARIO

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.carga() > MAXIMA_CARGA {
		nuevoTam := hash.tam * FACTOR_AUMENTO
		hash.redimensionar(nuevoTam)
	}
	hashPos := hash.recorrer(clave)
	if hash.tabla[hashPos].estado == VACIO {
		hash.cantidad++
	}
	if hash.tabla[hashPos].estado == BORRADO {
		hash.cantidad++
		hash.borrados--
	}
	hash.tabla[hashPos].clave = clave
	hash.tabla[hashPos].dato = dato
	hash.tabla[hashPos].estado = OCUPADO
}

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	if hash.Cantidad() == 0 {
		return false
	}
	hashPos := hash.recorrer(clave)
	return hash.tabla[hashPos].estado == OCUPADO
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	if hash.Cantidad() == 0 {
		panicNoPertenece()
	}
	hashPos := hash.recorrer(clave)
	if hash.tabla[hashPos].estado != OCUPADO {
		panicNoPertenece()
	}
	return hash.tabla[hashPos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if hash.Cantidad() == 0 {
		panicNoPertenece()
	}
	if hash.carga() < MINIMA_CARGA {
		nuevoTam := hash.tam / FACTOR_REDUCTOR
		hash.redimensionar(nuevoTam)
	}
	hashPos := hash.recorrer(clave)
	if hash.tabla[hashPos].estado != OCUPADO {
		panicNoPertenece()
	}
	hash.tabla[hashPos].estado = BORRADO
	hash.borrados++
	hash.cantidad--
	dato := hash.tabla[hashPos].dato
	return dato
}

func (hash hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(iterFunc func(clave K, dato V) bool) {
	for i := 0; i < hash.tam; i++ {
		if hash.tabla[i].estado == OCUPADO {
			clave := hash.tabla[i].clave
			dato := hash.tabla[i].dato
			if !iterFunc(clave, dato) {
				return
			}
		}
	}
}

//ITERADOR EXTERNO

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterador[K, V])
	iter.hash = hash
	if iter.hash.cantidad > 0 {
		for iter.hash.tabla[iter.posActual].estado != OCUPADO && iter.HaySiguiente() {
			iter.posActual++
		}
		iter.recorridos = 1
	}
	return iter
}

// devuelve si hay un elemento en la posActual

func (iter iterador[K, V]) HaySiguiente() bool {
	return iter.recorridos <= iter.hash.cantidad && iter.hash.cantidad != 0
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter iterador[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.hash.tabla[iter.posActual].clave
	dato := iter.hash.tabla[iter.posActual].dato
	return clave, dato
}

// Siguiente si HaySiguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter *iterador[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.posActual++
	for iter.recorridos < iter.hash.cantidad && iter.hash.tabla[iter.posActual].estado != OCUPADO {
		iter.posActual++
	}
	iter.recorridos++
}
