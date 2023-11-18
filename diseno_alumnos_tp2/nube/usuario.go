package nube

import (
	TDAColaP "tdas/cola_prioridad"
)

type Usuario interface {
	VerNombre() string
	VerSiguientePosteo() Posteo
	Feed() TDAColaP.ColaPrioridad[Posteo]
	Posicion() int
}
