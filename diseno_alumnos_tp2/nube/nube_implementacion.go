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

func (nube *nubeImplementacion) Likear(id int) error {
	if !nube.HayLogeado() || !nube.dicc_posteos.Pertenece(id) {
		return errores.ErrorUserNoLoggeadoOPostInexistente{}
	}

	posteo := nube.dicc_posteos.Obtener(id)
	nombre := nube.usuario_activo.Desapilar().VerNombre()

	posteo.AgregarLike(nombre)
	return nil
}

func (nube *nubeImplementacion) Publicar(contenido string) error {

	if !nube.HayLogeado() {
		return errores.ErrorSinUserLoggeado{}
	}
	nombre := nube.usuario_activo.Desapilar().VerNombre()
	posteo := CrearPosteo(nombre, contenido, nube.cantidad_posteo)
	nube.dicc_posteos.Guardar(posteo.VerID(), posteo)
	nube.cantidad_posteo++
	return nil
}

func (nube *nubeImplementacion) CrearRegistroUsuarios(nombre string) {
	usuario := CrearUsuario(nombre)
	nube.dicc_usuarios.Guardar(nombre, usuario)
}
