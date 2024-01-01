import random

class Grafo: 
    def __init__(self, dirigido = False, vertices = []):
        self.vertices = {}
        if vertices is None:
            vertices = []
        for v in vertices:
            if v not in self.vertices:
                self.vertices[v] = {}
        self.es_dirigido = dirigido

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}

    def borrar_vertices(self, vertice):
        if vertice in self.vertices:
            _ = self.vertices.pop(vertice)
            for w in self.vertices:
                if vertice in self.vertices[w]:
                    _ = self.vertices[w].pop(vertice)

    def agregar_arista(self, v, w, peso = 1):
        if v and w in self.vertices:
            self.vertices[v][w] = peso
            if not self.es_dirigido:
                self.vertices[w][v] = peso

    def borrar_arista(self, v, w):
        if v and w in self.vertices:
            _ = self.vertices[v].pop(w)
            if not self.es_dirigido:
                _ = self.vertices[w].pop(v)

    def estan_unidos(self, v, w):
        if v in self.vertices:
            return w in self.vertices[v]
        return False
    
    def peso_arista(self, v, w):
        if self.estan_unidos(v, w):
            return self.vertices[v][w]
        return None
    
    def obtener_vertices(self):
        res = []
        for v in self.vertices:
            res.append(v)
        return res
    
    def vertice_aleatorio(self):
        if self.vertices:
            return random.choice(list(self.vertices.keys()))
        return None
    
    def adyacentes(self, v):
        if v in self.vertices:
            return self.vertices[v]
        return {}
    
    def imprimir_grafo(self):
        for vertice, adyacentes in self.vertices.items():
            print(f"{vertice}: {adyacentes}")

"""
Grafo(dirigido = False, vertices_iniciales = []) para crear (hacer 'from grafo import Grafo')
agregar_vertice(self, v)
borrar_vertice(self, v)
agregar_arista(self, v, w, peso = 1)
borrar_arista(self, v, w)
estan_unidos(self, v, w)
peso_arista(self, v, w)
obtener_vertices(self)
Devuelve una lista con todos los vértices del grafo
vertice_aleatorio(self)
adyacentes(self, v)
str 
"""