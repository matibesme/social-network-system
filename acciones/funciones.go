package acciones

import (
	"bufio"
	"os"
	Diccionario "tdas/diccionario"
)

const ID = 0

func UsuariosEnArchivo(archivo_lista string) Diccionario.Diccionario[string, int] {
	usuarios := Diccionario.CrearHash[string, int]()
	archivo, err := os.Open(archivo_lista)

	if err != nil {
		return nil
	}
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {
		usuarios.Guardar(lector.Text(), ID+1)
	}

	return usuarios
}
