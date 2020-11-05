package main

import (
	"errors"
	"fmt"
)

var listaRecetas []receta

type receta struct {
	id         int
	nombre     string
	duracion   int
	dificultad string
}

func addReceta(id int, nombre string, duracion int, dificultad string) {
	r := receta{id, nombre, duracion, dificultad}
	listaRecetas = append(listaRecetas, r)
}

func getRecetas() []receta {
	return listaRecetas
}

func getReceta(id int) (receta, error) {
	var receta receta

	for _, rec := range listaRecetas {
		if rec.id == id {
			receta = rec
			return receta, nil
		}
	}

	return receta, errors.New("No se encontró receta con este id")
}

func updateReceta(id int, nombre string, duracion int, dificultad string) {
	for i := 0; i < len(listaRecetas); i++ {
		if listaRecetas[i].id == id {
			listaRecetas[i].nombre = nombre
			listaRecetas[i].duracion = duracion
			listaRecetas[i].dificultad = dificultad
		}
	}
}

func deleteReceta(id int) {
	for i, receta := range listaRecetas {
		if receta.id == id {
			listaRecetas = listaRecetas[:i+copy(listaRecetas[i:], listaRecetas[i+1:])]
		}
	}
}

func main() {
	//creo recetas
	addReceta(1, "Tortilla de papa", 30, "Media")
	addReceta(2, "Bifes a la criolla", 60, "Baja")

	//pido las recetas
	recetas := getRecetas()

	//recorro e imprimo las recetas
	for _, v := range recetas {
		fmt.Println(v)
	}

	//pido receta por id válido
	receta1, error := getReceta(1)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(receta1)
	}

	//pido receta por id inválido
	receta2, error := getReceta(3)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(receta2)
	}

	//edito receta
	updateReceta(1, "Editada", 20, "Alta")
	receta3, error := getReceta(1)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(receta3)
	}

	//borro receta
	deleteReceta(1)
	receta4, error := getReceta(1)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(receta4)
	}

}
