package nube

import (
	TDAColaP "tdas/cola_prioridad"
)

type usuarioImplementacion struct {
	nombre       string
	cola_posteos TDAColaP.ColaPrioridad[Posteo]
	cmp          func(user usuarioImplementacion) func(post1, post2 Posteo) int
	posicion     int
}

func CrearUsuario(nombre string, pos int, cmp func(user usuarioImplementacion) func(post1, post2 Posteo) int) Usuario {
	var usuario usuarioImplementacion
	usuario.nombre = nombre
	usuario.cmp = cmp
	usuario.posicion = pos
	usuario.cola_posteos = TDAColaP.CrearHeap(cmp(usuario))
	return &usuario
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

func Cmp(usuario usuarioImplementacion) func(post1, post2 Posteo) int {
	return func(post1, post2 Posteo) int {
		us1 := post1.VerUsuario()
		us2 := post2.VerUsuario()

		afinidad1 := modulo(usuario.posicion - us1.Posicion())
		afinidad2 := modulo(usuario.posicion - us2.Posicion())
		if afinidad1 < afinidad2 {
			return 1
		}
		if afinidad1 > afinidad2 {
			return -1
		}

		if post1.VerID() < post2.VerID() {
			return 1
		}
		if post1.VerID() > post2.VerID() {
			return -1
		}

		return 0
	}

}

func modulo(numero int) int {
	if numero >= 0 {
		return numero
	} else {
		return -numero
	}
}
