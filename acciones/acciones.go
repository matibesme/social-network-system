package acciones

import (
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	"TP2/diseno_alumnos_tp2/nube"
	"fmt"
	"strconv"
	"strings"
)

var COMANDOS = []string{"login", "logout", "publicar", "ver_siguiente_feed", "likear_post", "mostrar_likes"}

func EjecutarLogInEImprimirRespuesta(nombre_lista []string, nube nube.Nube) {
	nombre_usuario := strings.Join(nombre_lista, " ")
	err := nube.Logear(nombre_usuario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hola %s\n", nombre_usuario)

}

func EjecutarLogOutEImprimirRespuesta(nube nube.Nube) {
	err := nube.LogOut()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Adios")

}

func EjecutarPublicarEImprimirRespuesta(contenido []string, nube nube.Nube) {

	publicacion := strings.Join(contenido, " ")

	err := nube.Publicar(publicacion)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Post publicado")
}

func EjecutarVerSiguienteFeedEImprimirRespuesta(nube nube.Nube) {

	if !nube.HayLogeado() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}
	user := nube.UsuarioActual()
	posteo := user.VerSiguientePosteo()
	if posteo != nil {
		posteo.ImprimirInformacion()
		return
	}
	fmt.Println("Usuario no loggeado o no hay mas posts para ver")

}

func EjecutarLikearPostEImprimirRespuesta(id_str string, nube nube.Nube) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return
	}

	errores := nube.Likear(id)

	if errores != nil {
		fmt.Println(errores)
		return
	}
	fmt.Println("Post likeado")

}

func EjecutarMostrarLikesEImprimirRespuesta(id_str string, nube nube.Nube) {
	id, err := strconv.Atoi(id_str)
	posteo := nube.VerPosteo(id)

	if err != nil {
		return
	}

	if posteo == nil {
		fmt.Println(errores.ErrorPostInexistenteOSinLikes{})
		return
	}

	arreglo_likes := posteo.MostrarLikes()

	if len(arreglo_likes) == 0 {
		fmt.Println(errores.ErrorPostInexistenteOSinLikes{})
	} else {
		fmt.Printf("El post tiene %d likes:\n", len(arreglo_likes))
	}

	for i, usuario := range arreglo_likes {
		if i != len(arreglo_likes) {
			i++
			fmt.Printf("	%s\n", usuario)
		} else {
			fmt.Printf("	%s", usuario)
		}
	}

}
