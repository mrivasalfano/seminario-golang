package main

import (
	"fmt"
	// "errors"
)

type paginaRecetas struct {
	nombre  string
	recetas []receta
}

type receta struct {
	nombre     string
	duracion   int
	dificultad string
}

func (p *paginaRecetas) addReceta(nombre string, duracion int, dificultad string) {
	r := receta{nombre, duracion, dificultad}
	p.recetas = append(p.recetas, r)
}

func (p *paginaRecetas) getRecetas() []receta {
	return p.recetas
}

func main() {
	//creo slice de recetas con capacidad inicial de 10
	var sliceRecetas = make([]receta, 10)

	//creo pagina de recetas
	paginaRecetas := paginaRecetas{"Recetas todos los días", sliceRecetas}

	//agrego recetas a la pagina
	paginaRecetas.addReceta("Tortilla de papa", 30, "Media")
	paginaRecetas.addReceta("Bifes a la criolla", 60, "Baja")

	//pido las recetas a la pagina
	recetas := paginaRecetas.getRecetas()

	//recorro e imprimo las recetas
	for _, v := range recetas {
		fmt.Println(v)
	}
}
