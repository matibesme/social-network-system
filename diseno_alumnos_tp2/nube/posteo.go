package nube

type Posteo interface {
	VerID() int
	VerContenido() string
	VerUsuario() //Usuario
	VerCantidadLikes() int
	MostrarLikes() //lista o dicc usuarios

}
