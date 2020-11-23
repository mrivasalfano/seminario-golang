package receta

import (
	"database/sql"

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
	AddReceta(Receta) sql.Result
	FindByID(string) *Receta
	FindAll() []*Receta
	UpdateReceta(Receta) sql.Result
	DeleteReceta(Receta) sql.Result
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddReceta(r Receta) sql.Result {
	query := `INSERT INTO receta (nombre, duracion, dificultad) VALUES (?,?,?)`
	return s.db.MustExec(query, r.Nombre, r.Duracion, r.Dificultad)
}

func (s service) UpdateReceta(r Receta) sql.Result {
	query := `UPDATE receta SET nombre = ?, duracion = ?, dificultad = ? WHERE id = ?`
	return s.db.MustExec(query, r.Nombre, r.Duracion, r.Dificultad, r.ID)
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

func (s service) DeleteReceta(r Receta) sql.Result {
	query := `DELETE FROM receta WHERE id = ?`
	return s.db.MustExec(query, r.ID)
}
