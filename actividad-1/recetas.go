package main

import (
	"errors"
	"fmt"
	// "errors"
)

type paginaRecetas struct {
	nombre  string
	recetas []receta
}

type receta struct {
	id         int
	nombre     string
	duracion   int
	dificultad string
}

func (p *paginaRecetas) addReceta(id int, nombre string, duracion int, dificultad string) {
	r := receta{id, nombre, duracion, dificultad}
	p.recetas = append(p.recetas, r)
}

func (p *paginaRecetas) getRecetas() []receta {
	return p.recetas
}

func (p *paginaRecetas) getReceta(id int) (receta, error) {
	var receta receta

	for _, rec := range p.recetas {
		if rec.id == id {
			receta = rec
			return receta, nil
		}
	}

	return receta, errors.New("No se encontró receta con este id")
}

func (p *paginaRecetas) updateReceta(id int, nombre string, duracion int, dificultad string) {
	for i := 0; i < len(p.recetas); i++ {
		if p.recetas[i].id == id {
			p.recetas[i].nombre = nombre
			p.recetas[i].duracion = duracion
			p.recetas[i].dificultad = dificultad
		}
	}
}

func (p *paginaRecetas) deleteReceta(id int) {
	for i, receta := range p.recetas {
		if receta.id == id {
			p.recetas = p.recetas[:i+copy(p.recetas[i:], p.recetas[i+1:])]
		}
	}
}

func main() {
	//creo slice de recetas con capacidad inicial de 10
	var sliceRecetas = make([]receta, 10)

	//creo pagina de recetas
	paginaRecetas := paginaRecetas{"Recetas todos los días", sliceRecetas}

	//agrego recetas a la pagina
	paginaRecetas.addReceta(1, "Tortilla de papa", 30, "Media")
	paginaRecetas.addReceta(2, "Bifes a la criolla", 60, "Baja")

	//pido las recetas a la pagina
	recetas := paginaRecetas.getRecetas()

	//recorro e imprimo las recetas
	for _, v := range recetas {
		fmt.Println(v)
	}
}
