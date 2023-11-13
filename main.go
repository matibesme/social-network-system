package main

import (
	errores "/diseno_alumnos_tp2/errores_tp2"
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdas/cola"
	"tp1/acciones"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	dicc_usuarios := acciones.PartidosEnArchivo(os.Args[1])

	if dicc.Cantidad {

		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}

	crear_partidos := acciones.CrearPartidos(lista_partidos)
	padrones_ordenados := acciones.OrdenarPadron(lista_padrones)
	impugnados_totales := 0
	cola_votantes := cola.CrearColaEnlazada[int]()
	texto_ingresado := bufio.NewScanner(os.Stdin)

	for texto_ingresado.Scan() {
		texto_ingresado := strings.Split(texto_ingresado.Text(), " ")
		var hayError error
		switch texto_ingresado[0] {
		case acciones.ENTRADA[0]:
			hayError = acciones.AccionIngresarVotante(texto_ingresado[1], cola_votantes, padrones_ordenados)

		case acciones.ENTRADA[1]:
			hayError = acciones.AccionVotar(texto_ingresado, cola_votantes, padrones_ordenados, crear_partidos, lista_partidos)

		case acciones.ENTRADA[2]:
			hayError = acciones.AccionDeshacer(cola_votantes, padrones_ordenados)

		case acciones.ENTRADA[3]:
			hayError = acciones.AccionFinVotante(cola_votantes, padrones_ordenados, crear_partidos, &impugnados_totales)
		}
		if hayError == nil {
			fmt.Println("OK")
		} else {
			fmt.Println(hayError)
		}

	}
	acciones.AccionResultadosElectorales(crear_partidos, cola_votantes, padrones_ordenados, &impugnados_totales)

}
