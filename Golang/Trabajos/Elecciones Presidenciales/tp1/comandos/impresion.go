package comandos

import (
	"fmt"
	votos "rerepolez/votos"
)

// Imprime el nombre de un tipo de voto específico

func ImprimirTipoVoto(tipo votos.TipoVoto) {
	switch tipo {
	case votos.PRESIDENTE:
		fmt.Println("Presidente:")
	case votos.GOBERNADOR:
		fmt.Println("Gobernador:")
	case votos.INTENDENTE:
		fmt.Println("Intendente:")
	}
}

// Muestra los resultados de votación para un cargo específico y sus respectivos partidos.

func mostrarResultadosPorCargo(cargo votos.TipoVoto, partidos []votos.Partido) {
	ImprimirTipoVoto(cargo)
	for _, partido := range partidos {
		fmt.Println(partido.ObtenerResultado(cargo))
	}
	fmt.Println()
}

// Muestra los resultados de votación por cargos y la cantidad de votos impugnados.

func MostrarResultados(partidos []votos.Partido, impugnados int) {
	mostrarResultadosPorCargo(votos.PRESIDENTE, partidos)
	mostrarResultadosPorCargo(votos.GOBERNADOR, partidos)
	mostrarResultadosPorCargo(votos.INTENDENTE, partidos)
	var strVoto string
	if impugnados == 1 {
		strVoto = "voto"
	} else {
		strVoto = "votos"
	}
	fmt.Printf("Votos Impugnados: %d %s", impugnados, strVoto)
	fmt.Println()
}
