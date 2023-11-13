package main

import (
	"TP2/acciones"
	errores "TP2/diseno_alumnos_tp2/errores_tp2"

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

	dicc_usuarios := acciones.PartidosEnArchivo(os.Args[1])

	if dicc.Cantidad {

		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}

	texto_ingresado := bufio.NewScanner(os.Stdin)

	for texto_ingresado.Scan() {
		texto_ingresado := strings.Split(texto_ingresado.Text(), " ")

		switch texto_ingresado[0] {
		case "login":
			acciones.AccionLogIn()

		case "logout":
			acciones.AccionLogOut()
		case "publicar":
			acciones.AccionPublicar()
		case "ver_siguiente_feed":
			acciones.AccionVerSiguienteFeed()
		case "likear_post":
			acciones.AccionLikearPost()
		case "mostrar_likes":
			acciones.MostrarLikes()
		}
	}
}
