package receta

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrivasalfano/seminario-golang/internal/config"
)

// Receta ...
type Receta struct {
	ID         int
	nombre     string
	duracion   int
	dificultad string
}

// RecetaService ...
type RecetaService interface {
	AddReceta(Receta) error
	FindByID(int) *Receta
	FindAll() []*Receta
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (RecetaService, error) {
	return service{db, c}, nil
}

func (s service) AddReceta(r Receta) error {
	return nil
}

func (s service) FindByID(ID int) *Receta {
	return nil
}

func (s service) FindAll() []*Receta {
	var list []*Receta
	s.db.Select(&list, "SELECT * FROM receta")
	return list
}
