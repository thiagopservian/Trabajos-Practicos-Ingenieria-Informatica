package comandos

import (
	"fmt"
	errores "rerepolez/errores"
	herramientas "rerepolez/herramientas"
	votos "rerepolez/votos"
	TDACola "tdas/cola"
)

// verifica si la cantidad de argumentos coincide con la esperada.

func validarArgumentos(cantidad_argumentos, argumentos_esperados int) error {
	if cantidad_argumentos != argumentos_esperados {
		return errores.ErrorExcesoParametros{}
	}
	return nil
}

// valida un número de DNI y busca al votante correspondiente en la lista de votantes.

func validacionDNI(dniString string, votantes []votos.Votante) votos.Votante {
	dni, err1 := herramientas.TransfomarString(dniString)
	if dni <= 0 || err1 != nil {
		fmt.Println(errores.DNIError{})
		return nil
	}
	votante := herramientas.BuscarVotante(0, len(votantes)-1, dni, votantes)
	if votante == nil {
		fmt.Println(errores.DNIFueraPadron{})
	}
	return votante
}

// realiza validaciones relacionadas con el voto, como verificar el número de argumentos, el tipo de cargo,
// la existencia de elementos en la cola y la validez de la lista de votos.

func ValidacionVoto(cargo, lista string, votante votos.Votante, cantidad_partidos, cantidad_argumentos int, fila TDACola.Cola[votos.Votante]) (int, votos.TipoVoto) {
	err0 := validarArgumentos(cantidad_argumentos, 3)
	if err0 != nil {
		fmt.Println(err0)
		return -1, -1
	}
	cargoTransformado := herramientas.TransformarTipoVoto(cargo)
	if cargoTransformado == -1 {
		fmt.Println(errores.ErrorTipoVoto{})
		return -1, -1
	}
	listaNumero, err := herramientas.TransfomarString(lista)
	if err != nil || (listaNumero > cantidad_partidos-1) {
		fmt.Println(errores.ErrorAlternativaInvalida{})
		return -1, -1
	}
	return listaNumero, cargoTransformado
}
