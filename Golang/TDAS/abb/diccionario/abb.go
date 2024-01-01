package diccionario

import (
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iteradorAbb[K comparable, V any] struct {
	elementos TDAPila.Pila[nodoAbb[K, V]]
	desde     *K
	hasta     *K
	cmp       funcCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}
func crearNodo[K comparable, V any](clave K, dato V) nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return *nodo
}

func buscar[K comparable, V any](nodo **nodoAbb[K, V], clave K, comparar func(K, K) int) **nodoAbb[K, V] {
	if nodo == nil || (*nodo) == nil {
		return nodo
	}
	comparacion := comparar(clave, (**nodo).clave)
	switch {
	case comparacion == 0:
		return nodo
	case comparacion > 0:
		return buscar(&(*nodo).derecho, clave, comparar)
	default:
		return buscar(&(*nodo).izquierdo, clave, comparar)
	}
}

func masGrande[K comparable, V any](puntero **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*puntero).derecho == nil {
		return puntero
	}
	return masGrande(&(*puntero).derecho)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	puntero := buscar(&abb.raiz, clave, abb.cmp)
	if *puntero == nil {
		panicNoPertenece()
	}
	dato := (*puntero).dato
	if (*puntero).derecho != nil && (*puntero).izquierdo != nil {
		reemplazante := masGrande(&(*puntero).izquierdo)
		Rclave := (*reemplazante).clave
		Rdato := abb.Obtener(Rclave)
		(*puntero).clave = Rclave
		(*puntero).dato = Rdato
		if (*reemplazante).izquierdo != nil {
			*reemplazante = (*reemplazante).izquierdo
		} else {
			*reemplazante = nil
		}
	} else if (*puntero).izquierdo != nil {
		*puntero = (*puntero).izquierdo
	} else if (*puntero).derecho != nil {
		*puntero = (*puntero).derecho
	} else {
		*puntero = nil
	}
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	puntero := buscar(&abb.raiz, clave, abb.cmp)
	if *puntero == nil {
		nodoNuevo := crearNodo[K, V](clave, dato) //creamos nodo y despues puntero lo apuntamos
		*puntero = &(nodoNuevo)
		abb.cantidad++
	} else {
		(*puntero).dato = dato
	}
}

func panicNoexiste() {
	panic("La clave no pertenece al diccionario")
}

func (abb *abb[K, V]) Obtener(clave K) V {
	puntero := buscar(&abb.raiz, clave, abb.cmp)
	if *puntero == nil {
		panicNoexiste()
	}
	return (*puntero).dato
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	puntero := buscar(&abb.raiz, clave, abb.cmp)
	return !(*puntero == nil)
}

// Iterador interno

func (abb abb[K, V]) Iterar(iterFunc func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, iterFunc)
}

func (abb abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	sigue := true
	abb.raiz.iterarRango(desde, hasta, visitar, &sigue, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde, hasta *K, visitar func(clave K, dato V) bool, sigue *bool, comparacion func(K, K) int) {
	if nodo == nil || !*sigue {
		return
	}
	noSeDescartaIzq := desde == nil || comparacion(nodo.clave, *desde) > 0 // actual > desde ==> no se descarta izq
	noSeDescartaDer := hasta == nil || comparacion(nodo.clave, *hasta) < 0 // actual < hasta ==> no se descarta der
	noSeDescarta := (desde == nil || comparacion(nodo.clave, *desde) >= 0) && (hasta == nil || comparacion(nodo.clave, *hasta) <= 0)
	// desde <= actual <= desde ==> no se descarta actual
	if noSeDescartaIzq {
		nodo.izquierdo.iterarRango(desde, hasta, visitar, sigue, comparacion)
	}
	if *sigue && noSeDescarta && !visitar(nodo.clave, nodo.dato) {
		*sigue = false
		return
	}
	if *sigue && noSeDescartaDer {
		nodo.derecho.iterarRango(desde, hasta, visitar, sigue, comparacion)
	}
}

// Iterador externo

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (nodo *nodoAbb[K, V]) apilarHijos(pila *TDAPila.Pila[nodoAbb[K, V]], comparacion func(K, K) int, desde, hasta *K) {
	for nodo != nil {
		if (desde == nil || comparacion(*desde, nodo.clave) <= 0) && (hasta == nil || comparacion(*hasta, nodo.clave) >= 0) {
			(*pila).Apilar(*nodo)
			nodo = nodo.izquierdo
		} else if desde == nil || comparacion(*desde, nodo.clave) > 0 { //clave más chica que el desde
			nodo = nodo.derecho
		} else if hasta == nil || comparacion(*hasta, nodo.clave) < 0 { //clave más grande que el hasta
			nodo = nodo.izquierdo
		} else {
			break
		}
	}
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorAbb[K, V])
	iter.cmp = abb.cmp
	iter.elementos = TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	abb.raiz.apilarHijos(&iter.elementos, iter.cmp, desde, hasta)
	iter.desde = desde
	iter.hasta = hasta
	return iter
}

func (iter *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.elementos.EstaVacia()
}

func (iter iteradorAbb[K, V]) panicIter() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	iter.panicIter()
	desapilado := iter.elementos.Desapilar()
	desapilado.derecho.apilarHijos(&iter.elementos, iter.cmp, iter.desde, iter.hasta) //y él mismo
}

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	iter.panicIter()
	nodo := iter.elementos.VerTope()
	return nodo.clave, nodo.dato
}
