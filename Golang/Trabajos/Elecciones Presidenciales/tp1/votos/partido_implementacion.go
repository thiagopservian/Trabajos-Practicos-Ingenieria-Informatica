package votos

import "fmt"

type candidato struct {
	nombre string
	votos  int
}

type PartidoImpl struct {
	nombre     string
	candidatos []candidato
}

type PartidoEnBlanco struct {
	nombre string
	votos  [CANT_VOTACION]int
}

func crearCandidato(nombre string) candidato {
	candidato := new(candidato)
	candidato.nombre = nombre
	return *candidato
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(PartidoImpl)
	partido.nombre = nombre
	for _, candidatoNombre := range candidatos {
		candidato := crearCandidato(candidatoNombre)
		partido.candidatos = append(partido.candidatos, candidato)
	}
	return partido
}

func CrearVotosEnBlanco() Partido {
	blanco := new(PartidoEnBlanco)
	blanco.nombre = "Votos en Blanco"
	return blanco
}

func (partido *PartidoImpl) VotadoPara(tipo TipoVoto) {
	switch tipo {
	case PRESIDENTE:
		partido.candidatos[PRESIDENTE].votos += 1
	case GOBERNADOR:
		partido.candidatos[GOBERNADOR].votos += 1
	case INTENDENTE:
		partido.candidatos[INTENDENTE].votos += 1
	}
}

func (partido PartidoImpl) ObtenerResultado(tipo TipoVoto) string {
	strVoto := ""
	if partido.candidatos[tipo].votos == 1 {
		strVoto = " voto"
	} else {
		strVoto = " votos"
	}
	return partido.nombre + " - " + partido.candidatos[tipo].nombre + ": " + fmt.Sprintf("%d", partido.candidatos[tipo].votos) + strVoto
}

func (blanco *PartidoEnBlanco) VotadoPara(tipo TipoVoto) {
	switch tipo {
	case PRESIDENTE:
		blanco.votos[PRESIDENTE] += 1
	case GOBERNADOR:
		blanco.votos[GOBERNADOR] += 1
	case INTENDENTE:
		blanco.votos[INTENDENTE] += 1
	}
}

func (blanco PartidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	strVoto := ""
	if blanco.votos[tipo] == 1 {
		strVoto = " voto"
	} else {
		strVoto = " votos"
	}
	return blanco.nombre + ": " + fmt.Sprintf("%d", blanco.votos[tipo]) + strVoto
}
