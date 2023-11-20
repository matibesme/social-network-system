package nube

import (
	TDADiccionario "tdas/diccionario"
)

type Posteo interface {
	VerID() int
	VerContenido() string
	VerUsuario() Usuario //Usuario
	VerCantidadLikes() int
	MostrarLikes() TDADiccionario.Diccionario[string, int]
	AgregarLike(nombre string)
}
