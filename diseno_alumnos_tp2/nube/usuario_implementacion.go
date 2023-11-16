package nube

import (
	errores "TP2/diseno_alumnos_tp2/errores_tp2"
	"fmt"
	TDAColaP "tdas/cola_prioridad"
	
)

type usuarioImplementacion struct {
	nombre string
	cola_posteos TDAColaP.ColaPrioridad[Posteo]
}

func CrearUsuario(nombre string) Usuario {
	return &usuarioImplementacion{
		nombre: nombre,
		cola_posteos: TDAColaP.CrearHeap[Posteo]()
	}
}


func(usuario *usuarioImplementacion) VerNombre() string{
	return usuario.nombre
} 


func(usuario *usuarioImplementacion) VerSiguientePosteo() string{

	if usuario.cola_posteos.EstaVacia(){
		return nil
	}

	return usuario.cola_posteos.Desencolar()

} 
