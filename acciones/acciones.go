package acciones

import (
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	"TP2/diseno_alumnos_tp2/nube"
	"fmt"
	"strconv"
)

func AccionLogIn(nombre string, nube nube.Nube) {
	err := nube.Logear(nombre)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf("Hola %s", nombre)
}

func AccionLogOut(nube nube.Nube) {
	err := nube.LogOut()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Adios")
}

func AccionPublicar(contenido string, nube nube.Nube) {
	err := nube.Publicar(contenido)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Post publicado")
}

func AccionVerSiguienteFeed(nube nube.Nube) {

	user := nube.UsuarioActual()

	if !nube.HayLogeado() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
	}
	posteo := user.VerSiguientePosteo()
	if posteo != nil {
		fmt.Sprintf("%s dijo: %d", posteo.VerUsuario(), posteo.VerContenido())
	} else {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
	}
}

func AccionLikearPost(id_str string, nube nube.Nube) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		//hacer todo aca adentro
	}

	errores := nube.Likear(id)

	if errores != nil {
		fmt.Println(errores)
	} else {
		fmt.Println("Post likeado")
	}

}

func AccionMostrarLikes(id_str string, nube nube.Nube) {
	id, err := strconv.Atoi(id_str)
	posteo := nube.VerPosteo(id)
	if posteo != nil {
		dicc_likes := posteo.MostrarLikes()
		fmt.Sprintf("El post tiene %s likes:", dicc_likes.Cantidad())
		for iter := dicc_likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			usuario, id := iter.VerActual()
			fmt.Println(usuario)
		}

	} else {
		fmt.Println(errores.ErrorPostInexistenteOSinLIkes{})
	}

}
