import biblioteca.funciones as func

class Netstats:
    def __init__(self, grafo):
        self.grafo = grafo
        self.cfcs= []          #Lista de listas
        self.index_cfc = {}    #diccionario de vertices con sus cfc

    def camino(self, v, w):
        distancia, padres = func.distancia_bfs(self.grafo, v, w)
        if distancia > 0:
            camino = func.reconstruccion(v, w, padres)
            return distancia, camino
        return distancia, None

    # Busca todas las cfc y devuelve la cfc a la que pertenece v
    def conectados(self, v):
        if v not in self.index_cfc:
            self.buscar_cfc()
        return self.cfcs[self.index_cfc[v]]  

    # Encuentra todas las cfc y guarda en un dic a cual pertenece cada v
    def buscar_cfc(self):
        self.cfcs, self.index_cfc = func.cfc(self.grafo)

    # Crea un mini_grafo con la cfc de v ya que un ciclo solo puede generarse en su propia cfc
    def ciclo(self, p, n):
        cfc = self.conectados(p)
        mini_grafo = func.mini_grafo(self.grafo, cfc)
        return func.ciclo_largo(mini_grafo, int(n), p)

    def lectura(self, pags):
        mini_grafo = func.mini_grafo(self.grafo, pags)
        res = func.bfs_topologico(mini_grafo)
        if len(res)!=len(pags):
            return ["No existe forma de leer las paginas en orden"]
        res.reverse()
        return res

    def diametro(self):
        return func.diametro(self.grafo)

    def en_rango(self, rango, pagina):
        return func.rango_bfs(self.grafo, rango, pagina)

    def comunidades(self, pagina):
        return func.label_propagacion(self.grafo, pagina)

    def navegacion(self, pagina):
        return func.navegacion(self.grafo, pagina)

    def clustering(self, pag):
        if pag:
            if pag not in self.grafo.obtener_vertices():
                return
            return func.calcular_coeficiente(self.grafo, pag)
        return func.calcular_clustering_promedio(self.grafo)
