package receta

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/config"
)

// Receta ...
type Receta struct {
	ID         int64
	Nombre     string
	Duracion   int
	Dificultad string
}

// Service ...
type Service interface {
	AddReceta(Receta) (int64, error)
	FindByID(string) *Receta
	FindAll() []*Receta
	UpdateReceta(Receta, string)
	DeleteReceta(string) *Receta
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddReceta(r Receta) (int64, error) {
	query := `INSERT INTO receta (nombre, duracion, dificultad) VALUES (?,?,?)`
	return s.db.MustExec(query, r.Nombre, r.Duracion, r.Dificultad).LastInsertId()

}

func (s service) UpdateReceta(r Receta, id string) {
	query := `UPDATE receta SET nombre = ?, duracion = ?, dificultad = ? WHERE id = ?`
	s.db.MustExec(query, r.Nombre, r.Duracion, r.Dificultad, id)
}

func (s service) FindByID(ID string) *Receta {
	receta := &Receta{}
	query := `SELECT * FROM receta WHERE id = ?`
	err := s.db.Get(receta, query, ID)
	if err != nil {
		return nil
	}
	return receta
}

func (s service) FindAll() []*Receta {
	var list []*Receta
	if err := s.db.Select(&list, "SELECT * FROM receta"); err != nil {
		panic(err)
	}
	return list
}

func (s service) DeleteReceta(id string) *Receta {
	receta := s.FindByID(id)
	query := `DELETE FROM receta WHERE id = ?`
	s.db.MustExec(query, id).RowsAffected()
	return receta
}
