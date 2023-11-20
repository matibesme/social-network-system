package nube

import (
	TDAColaP "tdas/cola_prioridad"
)

// Usuario modela una persona registrada en la red social y sus componentes o acciones.
type Usuario interface {
	//VerNombre Devuelve el nombre del usuario.
	VerNombre() string
	//VerSiguientePosteo Devuelve el siguiente Posteo del usuario segun su prioridad. En el caso que no haya algun posteo para ver devuelve nil.
	VerSiguientePosteo() Posteo
	//Feed Devuelve el heap de posteos del usuario.
	Feed() TDAColaP.ColaPrioridad[Posteo]
	//Posicion Devuelve la posicion del usuario en el archivo de usuarios.
	Posicion() int
}
