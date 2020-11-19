package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrivasalfano/seminario-golang/internal/config"
)

func NewDatabase(conf *config.Config) (*sqlx.DB, err) {
	switch conf.DB.Type {
	case 'sqlite3':
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("Tipo de DB inválido")
	}
}
