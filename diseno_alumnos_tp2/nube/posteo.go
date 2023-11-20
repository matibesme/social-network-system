package nube

import (
	TDADiccionario "tdas/diccionario"
)

type Posteo interface {
	//VerID devuelve el id del posteo
	VerID() int
	//VerContenido devuelve el contenido del posteo
	VerContenido() string
	//VerUsuario devuelve el usuario que realizo el posteo
	VerUsuario() Usuario
	//VerCantidadLikes devuelve la cantidad de likes que tiene el posteo
	VerCantidadLikes() int
	//MostrarLikes devuelve un diccionario de los nombres de usuario que le de dieron like al posteo
	MostrarLikes() TDADiccionario.Diccionario[string, int]
	//AgregarLike agrega un like del nombre de usuario correspondiente
	AgregarLike(nombre string)
}
