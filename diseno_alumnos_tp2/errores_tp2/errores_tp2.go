package errores_tp2

type ErrorUserYaLoggeado struct{}

func (e ErrorUserYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUserNoExiste struct{}

func (e ErrorUserNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorSinUserLoggeado struct{}

func (e ErrorSinUserLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorUserNoLoggeadoOPostInexistente struct{}

func (e ErrorUserNoLoggeadoOPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostInexistenteOSinLikes struct{}

func (e ErrorPostInexistenteOSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErrorColaPosteosVacia struct{}

func (e ErrorColaPosteosVacia) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}
