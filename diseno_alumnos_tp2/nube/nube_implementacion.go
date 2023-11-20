package nube

import (
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	TDADiccionario "tdas/diccionario"
	TDAPila "tdas/pila"
)

type nubeImplementacion struct {
	dicc_usuarios   TDADiccionario.Diccionario[string, Usuario]
	usuario_activo  TDAPila.Pila[Usuario]
	cantidad_posteo int
	dicc_posteos    TDADiccionario.Diccionario[int, Posteo]
}

func CrearNube() Nube {
	return &nubeImplementacion{
		dicc_usuarios:   TDADiccionario.CrearHash[string, Usuario](),
		usuario_activo:  TDAPila.CrearPilaDinamica[Usuario](),
		cantidad_posteo: 0,
		dicc_posteos:    TDADiccionario.CrearHash[int, Posteo](),
	}
}

func (nube *nubeImplementacion) Logear(nombre string) error {

	if nube.HayLogeado() {
		return errores.ErrorUserYaLoggeado{}
	}
	if !nube.dicc_usuarios.Pertenece(nombre) {
		return errores.ErrorUserNoExiste{}
	}

	usuario := nube.dicc_usuarios.Obtener(nombre)
	nube.usuario_activo.Apilar(usuario)

	return nil

}

func (nube *nubeImplementacion) LogOut() error {
	if !nube.HayLogeado() {
		return errores.ErrorSinUserLoggeado{}
	}

	nube.usuario_activo.Desapilar()
	return nil
}

func (nube *nubeImplementacion) HayLogeado() bool {

	return !nube.usuario_activo.EstaVacia()
}

func (nube *nubeImplementacion) UsuarioActual() Usuario {

	return nube.usuario_activo.VerTope()
}

func (nube *nubeImplementacion) Likear(id int) error {
	if !nube.HayLogeado() || !nube.dicc_posteos.Pertenece(id) {
		return errores.ErrorUserNoLoggeadoOPostInexistente{}
	}

	posteo := nube.dicc_posteos.Obtener(id)
	nombre := nube.usuario_activo.VerTope().VerNombre()

	posteo.AgregarLike(nombre)
	return nil
}

func (nube *nubeImplementacion) Publicar(contenido string) error {

	if !nube.HayLogeado() {
		return errores.ErrorSinUserLoggeado{}
	}

	nombre := nube.usuario_activo.VerTope().VerNombre()
	posteo := CrearPosteo(nombre, contenido, nube.cantidad_posteo)
	nube.dicc_posteos.Guardar(posteo.VerID(), posteo)
	nube.cantidad_posteo++
	for iter := nube.dicc_usuarios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre_user, user := iter.VerActual()
		if nombre_user != nombre {
			feed := user.Feed()
			feed.Encolar(posteo)
		}
	}
	return nil
}

func (nube *nubeImplementacion) CrearRegistroUsuarios(nombre string) {
	usuario := CrearUsuario(nombre, nube.dicc_usuarios.Cantidad(), nube.comparacion)
	nube.dicc_usuarios.Guardar(nombre, usuario)
}

func (nube *nubeImplementacion) VerPosteo(id int) Posteo {
	if !nube.dicc_posteos.Pertenece(id) {
		return nil
	}
	return nube.dicc_posteos.Obtener(id)
}

func (nube *nubeImplementacion) comparacion(post1, post2 Posteo) int {

	usuario_actual := nube.usuario_activo.VerTope()

	us1 := nube.dicc_usuarios.Obtener(post1.VerUsuario())
	pos1 := us1.Posicion()
	us2 := nube.dicc_usuarios.Obtener(post2.VerUsuario())
	pos2 := us2.Posicion()

	pos_act := usuario_actual.Posicion()

	if modulo(pos_act-pos1) < modulo(pos_act-pos2) {
		return -1
	} else if modulo(pos_act-pos1) > modulo(pos_act-pos2) {
		return 1
	}

	if post1.VerID() < post2.VerID() {
		return -1
	}
	return 1

}

func modulo(numero int) int {
	if numero >= 0 {
		return numero
	} else {
		return -numero
	}
}
