package acciones

import (
	"TP2/diseno_alumnos_tp2/nube"
	"bufio"
	"os"
)

func LeerArchivosYCrearUsuario(archivo_lista string, nube nube.Nube) {

	archivo, err := os.Open(archivo_lista)

	if err != nil {
		return
	}
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {
		nube.CrearRegistroUsuarios(lector.Text())
	}

}
