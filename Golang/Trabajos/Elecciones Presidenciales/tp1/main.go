package main

import (
	"bufio"
	"fmt"
	"os"
	comandos "rerepolez/comandos"
	errores "rerepolez/errores"
	herramientas "rerepolez/herramientas"
	votos "rerepolez/votos"
	"strings"
	TDACola "tdas/cola"
)

const (
	PRIMER_ARGUMENTO = 1
	POSICION_LISTA   = 0
	POSICION_PADRON  = 1
	CANT_ARCHIVOS    = 2
)

func main() {
	parametros := os.Args[PRIMER_ARGUMENTO:] // ./rerepolez ARCHIVO_LISTA ARCHIVO_PADRON
	if len(parametros) != CANT_ARCHIVOS {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	lista, padron := parametros[POSICION_LISTA], parametros[POSICION_PADRON]
	votantes, partidos := inicializar(padron, lista)
	if len(votantes) == 0 || len(partidos) == 0 {
		return
	}
	fila := TDACola.CrearColaEnlazada[votos.Votante]()
	var impugnados int
	ejecutarComandosDesdeEntrada(&votantes, &partidos, &fila, &impugnados)
	comandos.MostrarResultados(partidos, impugnados)
}

// lee el archivo de padrón y crea una lista de votantes. Los ordena y los devuelve.
func inicializarVotantes(padron string) []votos.Votante {
	archivo, err := os.Open(padron)
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return nil
	}
	defer archivo.Close()
	votantes := make([]votos.Votante, 0, 1)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		dni, err := herramientas.TransfomarString(linea)
		if err != nil {
			fmt.Println(errores.ErrorLeerArchivo{})
			return nil
		}
		votantes = append(votantes, votos.CrearVotante(dni))
	}
	err = s.Err()
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
	}
	votantes = herramientas.OrdenarVotantes(votantes)
	return votantes
}

// lee el archivo de lista y crea una lista de partidos.
func inicalizarPartidos(lista string) []votos.Partido {
	archivo, err := os.Open(lista)
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return nil
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	partidos := make([]votos.Partido, 0, 1)
	partidos = append(partidos, votos.CrearVotosEnBlanco())
	for s.Scan() {
		linea := strings.Split((s.Text()), ",")
		var candidato [votos.CANT_VOTACION]string
		copy(candidato[:votos.CANT_VOTACION], linea[1:])
		partido := votos.CrearPartido(linea[0], candidato)
		partidos = append(partidos, partido)
	}
	err = s.Err()
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
	}
	return partidos
}

// se encarga de cargar los votantes y partidos desde los archivos de padrón y lista.
func inicializar(padron, lista string) ([]votos.Votante, []votos.Partido) {
	votantes := inicializarVotantes(padron)
	partidos := inicalizarPartidos(lista)
	return votantes, partidos
}

// lee los comandos desde la entranda estandar, los procesa y muestra errores de ser necesario
func ejecutarComandosDesdeEntrada(votantes *[]votos.Votante, partidos *[]votos.Partido, fila *TDACola.Cola[votos.Votante], impugnados *int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		comando := strings.Fields(linea)
		if len(comando) >= 1 {
			comandos.ProcesarComando(comando, votantes, partidos, fila, impugnados)
		} else {
			fmt.Println(errores.ErrorParametros{})
		}
	}
	if !(*fila).EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{})
	}
}
