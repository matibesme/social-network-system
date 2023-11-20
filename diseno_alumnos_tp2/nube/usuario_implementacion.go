package nube

import (
	TDAColaP "tdas/cola_prioridad"
)

type usuarioImplementacion struct {
	nombre       string
	cola_posteos TDAColaP.ColaPrioridad[Posteo]
	posicion     int
}

func CrearUsuario(nombre string, pos int, cmp func(Posteo, Posteo) int) Usuario {
	return &usuarioImplementacion{
		nombre:       nombre,
		cola_posteos: TDAColaP.CrearHeap[Posteo](cmp),
		posicion:     pos,
	}
}

func (usuario *usuarioImplementacion) VerNombre() string {
	return usuario.nombre
}

func (usuario *usuarioImplementacion) VerSiguientePosteo() Posteo {

	if usuario.cola_posteos.EstaVacia() {
		return nil
	}

	return usuario.cola_posteos.Desencolar()

}

func (usuario *usuarioImplementacion) Posicion() int {
	return usuario.posicion
}

func (usuario *usuarioImplementacion) Feed() TDAColaP.ColaPrioridad[Posteo] {

	return usuario.cola_posteos

}
