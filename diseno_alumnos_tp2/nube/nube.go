package nube

type Nube interface {
	Logear(nombre string) error
	LogOut() error
	HayLogeado() bool

	Likear(id int) error
	Publicar(contenido string) error
	CrearRegistroUsuarios(nombre string)
	UsuarioActual() Usuario
	VerPosteo(id int) Posteo
}
