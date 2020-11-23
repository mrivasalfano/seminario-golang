package receta

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type httpService struct {
	endpoints []*endpoint
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

//NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/recetas",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/recetas/:id",
		function: getReceta(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/recetas",
		function: addReceta(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/recetas/:id",
		function: updateReceta(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/recetas/:id",
		function: deleteReceta(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"recetas": s.FindAll(),
		})
	}
}

func getReceta(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"receta": s.FindByID(c.Param("id")),
		})
	}
}

func addReceta(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Receta{}
		c.Bind(&requestBody)
		receta := Receta{
			Nombre:     requestBody.Nombre,
			Duracion:   requestBody.Duracion,
			Dificultad: requestBody.Dificultad,
		}

		id, err := s.AddReceta(receta).LastInsertId()

		if err != nil {
			c.Status(500)
		}

		receta.ID = id
		c.JSON(http.StatusCreated, receta)
	}
}

func updateReceta(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Receta{}
		c.Bind(&requestBody)
		receta := Receta{
			Nombre:     requestBody.Nombre,
			Duracion:   requestBody.Duracion,
			Dificultad: requestBody.Dificultad,
		}

		s.UpdateReceta(receta, c.Param("id"))
		c.JSON(http.StatusCreated, receta)
	}
}

func deleteReceta(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Receta{}
		c.Bind(&requestBody)
		s.DeleteReceta(c.Param(("id")))
		c.Status(200)
	}
}

//Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
