package receta

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/config"
)

// Receta ...
type Receta struct {
	ID         int
	Nombre     string
	Duracion   int
	Dificultad string
}

// Service ...
type Service interface {
	AddReceta(Receta) error
	FindByID(string) *Receta
	FindAll() []*Receta
	UpdateReceta(Receta) error
	DeleteReceta(string) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddReceta(r Receta) error {
	insertReceta := `INSERT INTO receta (nombre, duracion, dificultad) VALUES (?,?,?)`
	s.db.MustExec(insertReceta, r.Nombre, r.Duracion, r.Dificultad)

	return nil
}

func (s service) UpdateReceta(r Receta) error {
	insertReceta := `UPDATE receta SET nombre = ?, duracion = ?, dificultad = ? WHERE id = ?`
	s.db.MustExec(insertReceta, r.Nombre, r.Duracion, r.Dificultad, r.ID)

	return nil
}

func (s service) FindByID(ID string) *Receta {
	return &Receta{}
}

func (s service) FindAll() []*Receta {
	var list []*Receta
	if err := s.db.Select(&list, "SELECT * FROM receta"); err != nil {
		panic(err)
	}
	return list
}

func (s service) DeleteReceta(id string) error {
	insertReceta := `DELETE FROM receta WHERE id = ?`
	s.db.MustExec(insertReceta, id)

	return nil
}
