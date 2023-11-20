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
		case acciones.COMANDOS[0]:
			acciones.AccionLogIn(texto_ingresado[1:], nube)
		case acciones.COMANDOS[1]:
			acciones.AccionLogOut(nube)
		case acciones.COMANDOS[2]:
			acciones.AccionPublicar(texto_ingresado[1:], nube)
		case acciones.COMANDOS[3]:
			acciones.AccionVerSiguienteFeed(nube)
		case acciones.COMANDOS[4]:
			acciones.AccionLikearPost(texto_ingresado[1], nube)
		case acciones.COMANDOS[5]:
			acciones.AccionMostrarLikes(texto_ingresado[1], nube)
		}
	}
}
