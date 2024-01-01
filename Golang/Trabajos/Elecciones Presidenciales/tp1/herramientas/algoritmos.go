package herramientas

import (
	votos "rerepolez/votos"
	"strconv"
)

//ORDENAMIENTOS

// ordena a los votantes usando el algoritmo Counting Sort con división de DNI.

func countingSortTresDigitos(votantes []votos.Votante, divisor int) []votos.Votante {
	frequencias := make([]int, 1000)
	for _, votante := range votantes {
		dni := (votante.LeerDNI() / divisor) % 1000
		frequencias[dni] += 1
	}

	sumas_acumuladas := make([]int, 1000)
	for i := 1; i < len(frequencias); i++ {
		sumas_acumuladas[i] = sumas_acumuladas[i-1] + frequencias[i-1]
	}

	ordenadas := make([]votos.Votante, len(votantes))
	for _, votante := range votantes {
		dni := (votante.LeerDNI() / divisor) % 1000
		ordenadas[sumas_acumuladas[dni]] = votante
		sumas_acumuladas[dni] += 1
	}
	return ordenadas
}

// ordena a los votantes utilizando Radix Sort con múltiples pasos de Counting Sort.

func OrdenarVotantes(votantes []votos.Votante) []votos.Votante {
	votantes = countingSortTresDigitos(votantes, 1)
	votantes = countingSortTresDigitos(votantes, 1000)
	votantes = countingSortTresDigitos(votantes, 1000000)
	return votantes
}

//BUSQUEDAS

// realiza una búsqueda binaria en el slice de votantes.

func BuscarVotante(inicio, final, buscado int, slice []votos.Votante) votos.Votante {
	if inicio > final {
		return nil
	}
	medio := (inicio + final) / 2
	if slice[medio].LeerDNI() == buscado {
		return slice[medio]
	}
	if slice[medio].LeerDNI() > buscado {
		return BuscarVotante(inicio, medio-1, buscado, slice)
	}

	return BuscarVotante(medio+1, final, buscado, slice)
}

//TRANSFORMACIONES

// convierte una cadena de texto en un entero y maneja errores.

func TransfomarString(numero string) (int, error) {
	numeroTransformado, err := strconv.Atoi(numero)
	return numeroTransformado, err
}

// transforma una cadena en un tipo de voto especifico
func TransformarTipoVoto(tipo_dado string) votos.TipoVoto {
	if tipo_dado == "Presidente" {
		return votos.PRESIDENTE
	}
	if tipo_dado == "Gobernador" {
		return votos.GOBERNADOR
	}
	if tipo_dado == "Intendente" {
		return votos.INTENDENTE
	}
	return -1
}
