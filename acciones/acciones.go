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
	} else {
		fmt.Printf("Hola %s \n", nombre)
	}

}

func AccionLogOut(nube nube.Nube) {
	err := nube.LogOut()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adios")
	}
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

	if !nube.HayLogeado() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
	} else {
		user := nube.UsuarioActual()
		posteo := user.VerSiguientePosteo()
		if posteo != nil {
			fmt.Printf("%s dijo: %s\n", posteo.VerUsuario().VerNombre(), posteo.VerContenido())
		} else {
			fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		}
	}
}

func AccionLikearPost(id_str string, nube nube.Nube) {
	id, _ := strconv.Atoi(id_str)

	errores := nube.Likear(id)

	if errores != nil {
		fmt.Println(errores)
	} else {
		fmt.Println("Post likeado")
	}

}

func AccionMostrarLikes(id_str string, nube nube.Nube) {
	id, _ := strconv.Atoi(id_str)
	posteo := nube.VerPosteo(id)
	if posteo != nil {
		dicc_likes := posteo.MostrarLikes()
		if dicc_likes.Cantidad() == 0 {
			fmt.Println(errores.ErrorPostInexistenteOSinLikes{})
		} else if dicc_likes.Cantidad() == 1 {
			fmt.Printf("El post tiene %d likes:\n", dicc_likes.Cantidad())
		} else {
			fmt.Printf("El post tiene %d likes:\n", dicc_likes.Cantidad())
		}
		for iter := dicc_likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			usuario, _ := iter.VerActual()
			fmt.Println(usuario)
		}

	} else {
		fmt.Println(errores.ErrorPostInexistenteOSinLikes{})
	}

}
