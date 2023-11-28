package nube

// Posteo modela una publicacion realizada por un usuario de la red social y sus respectivos componentes
type Posteo interface {
	//VerID devuelve el id del posteo
	VerID() int
	//VerContenido devuelve el contenido del posteo
	VerContenido() string
	//VerUsuario devuelve el usuario que realizo el posteo
	VerUsuario() Usuario
	//VerCantidadLikes devuelve la cantidad de likes que tiene el posteo
	VerCantidadLikes() int
	//MostrarLikes devuelve un arreglo de los nombres de usuario que le de dieron like al posteo
	MostrarLikes() []string
	//AgregarLike agrega un like del nombre de usuario correspondiente
	AgregarLike(nombre string)
	//ImprimirInformacion imprime el ID del posteo, el nombre del usuario que lo creo,
	//el contenido y la cantidad de likes
	ImprimirInformacion()
}
