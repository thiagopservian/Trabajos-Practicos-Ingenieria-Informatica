package comandos

import (
	"fmt"
	errores "rerepolez/errores"
	votos "rerepolez/votos"
	TDACola "tdas/cola"
)

// Procesa un voto dado y actualiza el estado de los partidos en función de los tipos de voto especificados.

func procesarVoto(partidos *[]votos.Partido, voto votos.Voto) {
	for i, valor := range voto.VotoPorTipo {
		(*partidos)[valor].VotadoPara(votos.TipoVoto(i))
	}
}

// procesa el comando "deshacer" verificando si la cola de votantes está vacía, y si no lo está,
// intenta deshacer el voto del votante en la cabeza de la cola, manejando errores según sea necesario.

func procesarComandoDeshacer(fila *TDACola.Cola[votos.Votante]) {
	if (*fila).EstaVacia() {
		fmt.Println(errores.FilaVacia{})
		return
	}
	err := (*fila).VerPrimero().Deshacer()
	if err != nil {
		fmt.Println(err)
		if _, ok := err.(errores.ErrorVotanteFraudulento); ok {
			(*fila).Desencolar()
		}
	} else {
		fmt.Println("OK")
	}
}

// procesa el comando "votar" verificando los argumentos, obteniendo el número de lista y el cargo a partir de la validación,
// y realizando el voto correspondiente por parte del votante en la cabeza de la cola.

func procesarComandoVotar(comando []string, partidos *[]votos.Partido, fila *TDACola.Cola[votos.Votante]) {
	if (*fila).EstaVacia() {
		fmt.Println(errores.FilaVacia{})
		return
	}
	numeroLista, cargo := ValidacionVoto(comando[1], comando[2], (*fila).VerPrimero(), len(*partidos), len(comando), (*fila))
	if numeroLista != -1 && cargo != -1 {
		err := (*fila).VerPrimero().Votar(cargo, numeroLista)
		if err != nil {
			fmt.Println(err)
			if _, ok := err.(errores.ErrorVotanteFraudulento); ok {
				(*fila).Desencolar()
			}
		} else {
			fmt.Println("OK")
		}
	}
}

// procesa el comando "ingresar" verificando los argumentos, validando el DNI del votante y encolando al votante en la cola.

func procesarComandoIngresar(comando []string, votantes []votos.Votante, fila *TDACola.Cola[votos.Votante]) {
	err := validarArgumentos(len(comando), 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	votante := validacionDNI(comando[1], votantes)
	if votante != nil {
		(*fila).Encolar(votante)
		fmt.Println("OK")
	}
}

// Procesa el comando "fin-votar" verificando si la cola de votantes está vacía. Si no lo está,
// verifica si el votante en la cabeza de la cola está impugnado, registra impugnación si es el caso,
// y luego procesa el voto o muestra errores según corresponda.

func procesarComandoFinVotar(partidos *[]votos.Partido, fila *TDACola.Cola[votos.Votante], impugnados *int) {
	if (*fila).EstaVacia() {
		fmt.Println(errores.FilaVacia{})
		return
	}
	impugnado := false
	if (*fila).VerPrimero().Impugnado() {
		*impugnados++
		impugnado = true
	}
	voto, err := (*fila).VerPrimero().FinVoto()
	(*fila).Desencolar()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !impugnado {
		procesarVoto(partidos, voto)
	}
	fmt.Println("OK")
}

// decide cómo procesar el comando en función de su primer argumento

func ProcesarComando(comando []string, votantes *[]votos.Votante, partidos *[]votos.Partido, fila *TDACola.Cola[votos.Votante], impugnados *int) {
	switch comando[0] {
	case "ingresar":
		procesarComandoIngresar(comando, *votantes, fila)
	case "votar":
		procesarComandoVotar(comando, partidos, fila)
	case "deshacer":
		procesarComandoDeshacer(fila)
	case "fin-votar":
		procesarComandoFinVotar(partidos, fila, impugnados)
	default:
		fmt.Println(errores.ErrorComando{})
	}
}
