package nube

import (
	"strings"
	TDADiccionario "tdas/diccionario"
)

type posteoImplementacion struct {
	id             int
	usuario        Usuario
	contenido      string
	likes          TDADiccionario.DiccionarioOrdenado[string, int]
	cantidad_likes int
}

func CrearPosteo(usuario Usuario, contenido string, cant_posteos int) Posteo {
	return &posteoImplementacion{
		id:             cant_posteos,
		usuario:        usuario,
		contenido:      contenido,
		likes:          TDADiccionario.CrearABB[string, int](strings.Compare),
		cantidad_likes: 0,
	}
}

func (posteo *posteoImplementacion) VerID() int {
	return posteo.id
}

func (posteo *posteoImplementacion) VerContenido() string {
	return posteo.contenido
}

func (posteo *posteoImplementacion) VerCantidadLikes() int {
	return posteo.cantidad_likes
}

func (posteo *posteoImplementacion) VerUsuario() Usuario {
	return posteo.usuario
}

func (posteo *posteoImplementacion) MostrarLikes() TDADiccionario.Diccionario[string, int] {

	return posteo.likes
}

func (posteo *posteoImplementacion) AgregarLike(nombre string) {
	if !posteo.likes.Pertenece(nombre) {
		posteo.likes.Guardar(nombre, 1)
		posteo.cantidad_likes++
	}

}
