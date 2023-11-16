package main

import (
	"TP2/acciones"
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	"TP2/diseno_alumnos_tp2/nube"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println(errores.ErrorUserNoExiste{})
		return
	}

	nube := nube.CrearNube()
	acciones.UsuariosEnArchivo(os.Args[1], nube)

	texto_ingresado := bufio.NewScanner(os.Stdin)

	for texto_ingresado.Scan() {
		texto_ingresado := strings.Split(texto_ingresado.Text(), " ")

		switch texto_ingresado[0] {
		case "login":
			acciones.AccionLogIn(texto_ingresado[1])
		case "logout":
			acciones.AccionLogOut()
		case "publicar":
			acciones.AccionPublicar(texto_ingresado[1])
		case "ver_siguiente_feed":
			acciones.AccionVerSiguienteFeed()
		case "likear_post":
			acciones.AccionLikearPost(texto_ingresado[1])
		case "mostrar_likes":
			acciones.MostrarLikes(texto_ingresado[1])
		}
	}
}
