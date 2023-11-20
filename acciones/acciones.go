package acciones

import (
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	"TP2/diseno_alumnos_tp2/nube"
	"fmt"
	"strconv"
	"strings"
)

var COMANDOS = []string{"login", "logout", "publicar", "ver_siguiente_feed", "likear_post", "mostrar_likes"}

func AccionLogIn(nombre_lista []string, nube nube.Nube) {
	nombre_usuario := strings.Join(nombre_lista, " ")
	err := nube.Logear(nombre_usuario)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Hola %s\n", nombre_usuario)
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

func AccionPublicar(contenido []string, nube nube.Nube) {

	publicacion := strings.Join(contenido, " ")

	err := nube.Publicar(publicacion)
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
			fmt.Printf("Post ID %d\n", posteo.VerID())
			fmt.Printf("%s dijo: %s\n", posteo.VerUsuario().VerNombre(), posteo.VerContenido())
			fmt.Printf("Likes: %d\n", posteo.VerCantidadLikes())
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
		} else {
			fmt.Printf("El post tiene %d likes:\n", dicc_likes.Cantidad())
		}
		contador := 0
		for iter := dicc_likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			usuario, _ := iter.VerActual()
			if contador != dicc_likes.Cantidad() {
				contador++
				fmt.Printf("	%s\n", usuario)
			} else {
				fmt.Printf("	%s", usuario)
			}
		}

	} else {
		fmt.Println(errores.ErrorPostInexistenteOSinLikes{})
	}

}
