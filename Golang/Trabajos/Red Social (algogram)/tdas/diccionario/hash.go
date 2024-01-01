package diccionario

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type estado int

const (
	TAMAÑO_ORIGINAL = 11
)
const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado // usar constante
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

func funcionHashing[K comparable](clave K, tam int) int {
	data := []byte(fmt.Sprintf("%v", clave))
	byt := sha256.Sum256(data)
	intValue := int(binary.BigEndian.Uint32(byt[:]))
	return intValue % tam
}

func (hash hashCerrado[K, V]) carga() float32 {
	ocupados := float32(hash.Cantidad())
	borrados := float32(hash.borrados)
	tamaño := float32(hash.tam)
	return float32((ocupados + borrados) / tamaño)
}

func (hash *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	nuevaTabla := crearTabla[K, V](nuevoTam)
	var hashValue int
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			hashValue = funcionHashing(celda.clave, nuevoTam)
			for i := hashValue; i < nuevoTam; i++ {
				if nuevaTabla[i].estado == VACIO {
					nuevaTabla[i] = celda
					break
				}
				if i == nuevoTam-1 {
					i = -1
				}
			}
		}
	}
	hash.borrados = 0
	hash.tam = nuevoTam
	hash.tabla = nuevaTabla
}

func panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.carga() > 0.7 {
		nuevoTam := hash.tam * 2
		hash.redimensionar(nuevoTam)
	}
	hashValue := funcionHashing(clave, hash.tam)
	for i := hashValue; i < hash.tam; i++ {
		if hash.tabla[i].estado == VACIO {
			hash.tabla[i].clave = clave
			hash.tabla[i].dato = dato
			hash.tabla[i].estado = OCUPADO
			hash.cantidad++
			return
		} else if hash.tabla[i].estado == OCUPADO && hash.tabla[i].clave == clave {
			hash.tabla[i].dato = dato
			return
		} else if hash.tabla[i].estado == BORRADO && hash.tabla[i].clave == clave {
			hash.tabla[i].dato = dato
			hash.tabla[i].estado = OCUPADO
		}
		if i == hash.tam-1 {
			i = -1
		}
	}
}

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	if hash.Cantidad() == 0 {
		return false
	}
	hashValue := funcionHashing[K](clave, hash.tam)
	for i := hashValue; i < hash.tam; i++ {
		if hash.tabla[i].estado == VACIO {
			return false
		}
		if hash.tabla[i].clave == clave {
			if hash.tabla[i].estado == BORRADO {
				return false
			}
			return true
		}
		if i == hash.tam-1 {
			i = -1
		}
	}
	return false
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	if hash.Cantidad() == 0 {
		panicNoPertenece()
	}
	hashValue := funcionHashing[K](clave, hash.tam)
	var dato V
	for i := hashValue; i < hash.tam; i++ {
		if hash.tabla[i].estado == VACIO {
			panicNoPertenece()
		}
		if hash.tabla[i].clave == clave && hash.tabla[i].estado == OCUPADO {
			dato = hash.tabla[i].dato
			break
		}
		if i == hash.tam-1 {
			i = -1
		}
	}
	return dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if hash.Cantidad() == 0 {
		panicNoPertenece()
	}
	if hash.carga() < 0.2 {
		nuevoTam := hash.tam / 2
		hash.redimensionar(nuevoTam)
	}
	hashValue := funcionHashing[K](clave, hash.tam)
	var dato V
	for i := hashValue; i < hash.tam; i++ {
		if hash.tabla[i].estado == VACIO {
			panicNoPertenece()
		}
		if hash.tabla[i].clave == clave && hash.tabla[i].estado == OCUPADO {
			hash.tabla[i].estado = BORRADO
			hash.borrados++
			hash.cantidad--
			dato = hash.tabla[i].dato
			break
		}
		if i == hash.tam-1 {
			i = -1
		}
	}
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
