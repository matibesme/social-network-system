package nube

type Nube interface {
	//Loguear ingresa el usuario, pasado por parametro, a su cuenta
	Logear(nombre string) error
	//Logear deslogea al usuario. Si habia un usuario loggeado, este se debe salir de su cuenta.
	//En caso que no haya habido un usuario loggeado, debe imprimir el correspondiente error
	LogOut() error
	//HayLogeado devuelve si existe actualemnete un usuario dentro de su cuenta
	HayLogeado() bool
	//Likear recibe el id del post a likear, y se agrega al usuario loggeado entre los que likean a dicho post
	Likear(id int) error
	//Publicar crea un post con un texto determinado, cuyo publicador es el usuario loggeado actualmente.
	//Dicho post se le publicara al resto de los usuarios
	Publicar(contenido string) error
	//CrearRegistroUsuarios crea un diccionario con todos los usuario recibidos en el archivo correspondiente
	CrearRegistroUsuarios(nombre string)
	//UsuarioActual devuelve en usuario que se encuentra atualmente dentro de su cuenta
	UsuarioActual() Usuario
	//VerPosteo recibe el id de un posteo y devuelve el posteo correspondiente
	VerPosteo(id int) Posteo
}
