package votos

import (
	errores "rerepolez/errores"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
)

const IMPUGNADO = 0

type ordenVotos struct {
	ordenTipos TDALista.Lista[TipoVoto]
	presidente TDAPila.Pila[int]
	gobernador TDAPila.Pila[int]
	intendente TDAPila.Pila[int]
}

type VotanteImpl struct {
	dni    int
	yaVoto bool
	orden  ordenVotos
	votos  Voto
}

func CrearVotante(dni int) Votante {
	votante := new(VotanteImpl)
	votante.dni = dni
	votante.orden = ordenVotos{
		ordenTipos: TDALista.CrearListaEnlazada[TipoVoto](),
		presidente: TDAPila.CrearPilaDinamica[int](),
		gobernador: TDAPila.CrearPilaDinamica[int](),
		intendente: TDAPila.CrearPilaDinamica[int](),
	}
	return votante
}

func (votante VotanteImpl) LeerDNI() int {
	return votante.dni
}

func (votante *VotanteImpl) Votar(tipo TipoVoto, alternativa int) error {

	if votante.yaVoto {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.orden.ordenTipos.InsertarPrimero(tipo)
	if alternativa == IMPUGNADO {
		votante.votos.VotosImpugnados++
		votante.votos.Impugnado = true
	}
	switch tipo {
	case PRESIDENTE:
		votante.orden.presidente.Apilar(alternativa)
	case GOBERNADOR:
		votante.orden.gobernador.Apilar(alternativa)
	case INTENDENTE:
		votante.orden.intendente.Apilar(alternativa)
	}
	return nil
}

func (votante VotanteImpl) validarDeshacer() error {
	var problema error
	if votante.yaVoto {
		problema = errores.ErrorVotanteFraudulento{Dni: votante.dni}
	} else if votante.orden.ordenTipos.EstaVacia() {
		problema = errores.ErrorNoHayVotosAnteriores{}
	}
	return problema
}

func (votante *VotanteImpl) Deshacer() error {
	problema := votante.validarDeshacer()
	if problema != nil {
		return problema
	}
	ultimoTipo := votante.orden.ordenTipos.BorrarPrimero()
	var votoEliminado int
	switch ultimoTipo {
	case PRESIDENTE:
		votoEliminado = votante.orden.presidente.Desapilar()
	case GOBERNADOR:
		votoEliminado = votante.orden.gobernador.Desapilar()
	case INTENDENTE:
		votoEliminado = votante.orden.intendente.Desapilar()
	}
	if votoEliminado == 0 {
		votante.votos.VotosImpugnados--
		if votante.votos.VotosImpugnados == 0 {
			votante.votos.Impugnado = false
		}
	}
	return problema

}

func (votante *VotanteImpl) FinVoto() (Voto, error) {
	if votante.yaVoto {
		return Voto{}, errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if votante.votos.Impugnado {
		return Voto{}, nil
	}
	if !votante.orden.presidente.EstaVacia() {
		votante.votos.VotoPorTipo[PRESIDENTE] = votante.orden.presidente.Desapilar()
	}
	if !votante.orden.gobernador.EstaVacia() {
		votante.votos.VotoPorTipo[GOBERNADOR] = votante.orden.gobernador.Desapilar()
	}
	if !votante.orden.intendente.EstaVacia() {
		votante.votos.VotoPorTipo[INTENDENTE] = votante.orden.intendente.Desapilar()
	}
	votante.yaVoto = true
	return votante.votos, nil
}

func (votante VotanteImpl) Impugnado() bool {
	return votante.votos.Impugnado
}
