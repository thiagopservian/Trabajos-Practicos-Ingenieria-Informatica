#!/usr/bin/python3

import sys, os
from TDAS.grafo import Grafo
from TDAS.netsats import Netstats

sys.setrecursionlimit(10000)

comandos_globales = ["camino", "ciclo","conectados", "diametro", "lectura", "rango", "comunidad", "navegacion", "clustering", "listar_operaciones"]

def main():
    parametros = sys.argv # $ ./netstats wiki-reducido-75000.tsv
    if len(parametros) != 2:
        return
    archivo_tsv = parametros[1]
    grafo = modelar_grafo(archivo_tsv)
    netstats = Netstats(grafo)
    comandos(netstats)
    
def modelar_grafo(archivo_tsv):
    grafo = Grafo(dirigido=True)
    if not os.path.exists(archivo_tsv):
        print(f"El archivo {archivo_tsv} no existe")
        return
    with open(archivo_tsv) as archivo:
        for linea in archivo:
            linea = linea.split("\t")
            pagina = linea[0].strip()
            links = linea[1:]
            grafo.agregar_vertice(pagina)
            for link in links:
                link = link.strip()
                grafo.agregar_vertice(link)
                grafo.agregar_arista(pagina, link)
    return grafo

def inicializar_opciones():
    diccionario_comandos = {
        "camino": camino,
        "conectados": conectados,
        "ciclo": ciclo,
        "lectura": lectura,
        "diametro": diametro,
        "rango": rango,
        "comunidad": comunidad,
        "navegacion": navegacion,
        "clustering": clustering,
        "listar_operaciones": listar_operaciones
    }
    return diccionario_comandos

def comandos(netstats):
    opciones = inicializar_opciones()
    for linea in sys.stdin:
        linea = linea.rstrip('\n')
        linea = linea.strip().split(" ", maxsplit=1)
        comando = linea[0].strip()
        parametros = []
        if len(linea) > 1:
            parametros = linea[1].split(',')
            parametros = [param.strip() for param in parametros]
        if comando in opciones:
            opciones[comando](netstats, parametros)

def imprimir_paginas_flecha(paginas):
    paginas_formato = " -> ".join(paginas)
    print(paginas_formato)

def imprimir_paginas_con_comas(paginas):
    paginas_formato = ", ".join(paginas)
    print(paginas_formato)

def camino(netstats, parametros):
    v = parametros[0]
    w = parametros[1]
    distancia, camino = netstats.camino(v, w)
    if distancia == -1 or camino == None:
        print("No se encontro recorrido")
        return
    imprimir_paginas_flecha(camino)
    print(f"Costo: {distancia}")

def conectados(netstats, parametros):
    v = parametros[0]
    lista =  netstats.conectados(v)
    imprimir_paginas_con_comas(lista)

def ciclo(netstats, parametros):
    p = parametros[0]
    n = parametros[1]
    ciclo =  netstats.ciclo(p, n)
    if not ciclo:
        print("No se encontro recorrido.")
        return
    imprimir_paginas_flecha(ciclo)

def lectura(netstats, parametros):
    lista = netstats.lectura(parametros)
    imprimir_paginas_con_comas(lista)

def diametro(netstats, _):
    diametro, camino = netstats.diametro()
    imprimir_paginas_flecha(camino)
    print(f"Costo: {diametro}")

def rango(netstats, parametros):
    rango = parametros[1]
    pagina = parametros[0]
    en_rango = netstats.en_rango(int(rango), pagina)
    print(en_rango)

def comunidad(netstats, parametros):
    pagina = parametros[0]
    imprimir_paginas_con_comas(netstats.comunidades(pagina))

def navegacion(netstats, parametros):
    pagina = parametros[0]
    lista =  netstats.navegacion(pagina)
    imprimir_paginas_flecha(lista)

def clustering(netstats, parametros):
    pagina = None
    if len(parametros) == 1:
        pagina = parametros[0]
    print(netstats.clustering(pagina))

def listar_operaciones(_, __):
    for comando in comandos_globales[:-1]:
        print(comando)

main()