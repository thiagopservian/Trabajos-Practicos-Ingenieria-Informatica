from TDAS.grafo import Grafo
from collections import deque
import heapq
import random
K  = 100

##COMUNIDADES (LABEL)
def max_freq_label(label, entrada):
    freq = {}
    max = label[entrada[0]]
    for w in entrada:
        freq[label[w]] = 0
    for w in entrada:
        freq[label[w]] += 1
        if freq[label[w]] > freq[max]:
            max = label[w]
    return max

def label_propagacion(grafo, vertice):
    label = {}
    for i,v in enumerate(grafo.obtener_vertices()):
        label[v] = i
    entrada = vertices_de_entrada(grafo)
    for _ in range(13):
        for v in grafo.obtener_vertices():
            if len(entrada[v]) > 0:
                label[v] = max_freq_label(label, entrada[v])
    res = []
    for v in grafo.obtener_vertices():
        if label[v] == label[vertice]:
            res.append(v)
    return res

def vertices_de_entrada(grafo):
    res = {}
    for v in grafo.obtener_vertices():
        res[v] = []
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            res[w].append(v)
    return res


def mini_grafo(grafo, vertices):
    mini = Grafo(dirigido=True, vertices=vertices)
    Ouno = set()
    for v in vertices:
        Ouno.add(v)
    for v in vertices:
        for w in grafo.adyacentes(v):
            if w in Ouno:
                mini.agregar_arista(v,w)
    return mini
  
#CICLO DE LARGO N
def ciclo_largo(grafo, n, v):
    origen = v
    camino = []
    visitados = set()
    visitados.add(v)
    camino.append(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados_camino = set()
            visitados_camino.add(w)
            visitados.add(w)
            camino.append(w)
            if _ciclo_largo(grafo, n-1, origen, w, visitados, visitados_camino, camino):
                camino.append(origen)
                return camino
            camino.remove(w)
            visitados.remove(w) 
    return None

def _ciclo_largo(grafo, n, origen, w, visitados, visitados_camino, camino):
    if n == 0:
        return w == origen
    for x in grafo.adyacentes(w):
        if x not in visitados:
            visitados_camino.add(x)
            visitados.add(x)
            camino.append(x)
            if _ciclo_largo(grafo, n-1, origen, x, visitados, visitados_camino, camino):
                return True
            visitados_camino.remove(x)
            camino.remove(x)
    return False

def encontrar_ciclo_de_longitud_n(grafo, inicio, longitud):
    visitados = set()
    camino = [inicio]
    return dfs_ciclo(grafo, inicio, inicio, longitud, visitados, camino)

def dfs_ciclo(grafo, nodo_actual, inicio, longitud, visitados, camino):
    visitados.add(nodo_actual)
    
    if len(camino) == longitud and camino[-1] == inicio:
        return camino  # Ciclo encontrado de longitud N
    
    if len(camino) < longitud:
        for vecino in grafo.adyacentes(nodo_actual):
            if vecino not in visitados:
                camino.append(vecino)
                ciclo = dfs_ciclo(grafo, vecino, inicio, longitud, visitados, camino)
                if ciclo:
                    return ciclo
                camino.pop()
    
    visitados.remove(nodo_actual)
    return None  # No se encontró un ciclo de longitud N



#COEFICIENTE DE CLUSTERING

def calcular_coeficiente(grafo, v): 
    adyacentes = grafo.adyacentes(v)
    if v in adyacentes:
        del adyacentes [v]
    grado = len(adyacentes)
    if grado < 2:
        return 0.000
    aristas_adyacentes_entre_si = 0
    for w in adyacentes:
        for x in adyacentes:
            if x == w:
                continue
            if grafo.estan_unidos(w, x):
                aristas_adyacentes_entre_si += 1
    coeficiente_clustering = aristas_adyacentes_entre_si / (grado * (grado - 1))
    return coeficiente_clustering

def calcular_clustering_promedio(grafo):
    clustering_total = 0
    for v in grafo.obtener_vertices():
        clustering_total += calcular_coeficiente(grafo, v)
    clustering_promedio = clustering_total / len(grafo.obtener_vertices())
    return clustering_promedio


#COMPONENTES FUERTEMENTES CONEXAS
def cfc(grafo):
    resultados = []
    visitados = set()
    index = {}
    for v in grafo.obtener_vertices():
        if v not in visitados:
            _cfc(grafo, v, visitados, {}, {}, deque(), set(), resultados, [0], index)
    return resultados, index

def _cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global, index):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.append(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            _cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global, index)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
    if orden[v] == mas_bajo[v]:
        nueva_cfc = set()
        while True:
            w = pila.pop()
            apilados.remove(w)
            nueva_cfc.add(w)
            index[w] = len(cfcs)
            if w == v:
                break
        cfcs.append(nueva_cfc)

#NAVEGACION POR PRIMER LINK
def navegacion(grafo, v):
    res = []
    _navegacion(grafo, v, res)
    return res

def _navegacion(grafo, w, res):
    if len(res)>20:
        return
    res.append(w)
    if len(grafo.adyacentes(w))<1:
        return
    for x in grafo.adyacentes(w):
        _navegacion(grafo, x, res)
        break
    


#DIAMETRO

def diametro(grafo):
    diametro = 0
    padres_camino_final = {}
    ini_final = None
    fin_final = None
    for v in grafo.obtener_vertices():
        padres, distancia = caminos_bfs(grafo, v)
        distancia_max, fin = distancia_mas_grande(distancia)
        if distancia_max > diametro:
            diametro = distancia_max
            padres_camino_final = padres
            ini_final = v
            fin_final = fin
    camino = reconstruccion(ini_final, fin_final, padres_camino_final)
    return diametro, camino

def distancia_mas_grande(distancia):
    max = 0
    fin = None
    for v in distancia:
        if distancia[v] > max:
            max, fin = distancia[v], v
    return max, fin


def caminos_bfs(grafo, v):
    visitados = set()
    visitados.add(v)
    padres = {}
    distancia = {}
    q = deque()
    q.append(v)
    padres[v] = None
    distancia[v] = 0
    while not len(q) == 0:
        w = q.popleft()
        for x in grafo.adyacentes(w):
            if x not in visitados:
                visitados.add(x)
                padres[x] = w
                distancia[x] = distancia[w] + 1
                q.append(x)
    return padres, distancia


def reconstruccion(ini, fin, padres):
    v = fin
    res = []
    while v != ini:
        res.append(v)
        v = padres[v]
    res.append(ini)
    res.reverse()
    return res


def distancia_bfs(grafo, origen, destino):
    visitados = set()
    padres = {}
    distancia = {}
    visitados.add(origen)
    padres[origen] = None
    distancia[origen] = 0
    q = deque()
    q.append(origen)

    while not len(q)==0:
        desencolado = q.popleft()
        for w in grafo.adyacentes(desencolado):
            if w not in visitados:
                visitados.add(w)
                padres[w] = desencolado
                distancia[w] = distancia[desencolado] + 1
                q.append(w)
            if w == destino:
                return distancia[w], padres

    return -1, padres


##TODOS EN RANGO
def rango_bfs(grafo, n, v):
    cont = 0
    visitados = set()
    orden = {}
    visitados.add(v)
    orden[v] = 0
    q = deque()
    q.append(v)
    while not len(q) == 0:
        desencolado = q.popleft()
        for w in grafo.adyacentes(desencolado):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[desencolado] + 1
                if orden[w] == n:
                    cont += 1
                    continue
                q.append(w)
    return cont

 

## ORDEN TOPOLOGICO

## bfs con grados
def bfs_topologico(grafo):
    grados = grados_entrada(grafo)
    q = deque()
    for v in grafo.obtener_vertices():
        if grados[v] == 0:
            q.append(v)
    resultado = []
    while not len(q) == 0:
        desencolado = q.popleft()
        resultado.append(desencolado)
        for w in grafo.adyacentes(desencolado):
            grados[w] -= 1
            if grados[w] == 0:
                q.append(w)
    return resultado

def grados_entrada(grafo):
    grados = {}
    for v in grafo.obtener_vertices():
        grados[v] = 0

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados

