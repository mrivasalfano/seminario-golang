package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/config"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/database"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/service/receta"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := receta.New(db, cfg)
	HTTPService := receta.NewHTTPTransport(service)

	r := gin.Default()
	HTTPService.Register(r)
	r.Run()
	// for _, m := range service.FindAll() {
	// 	fmt.Println(m)
	// }
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS receta (
		id integer primary key autoincrement,
		nombre varchar(50),
		duracion int,
		dificultad varchar (20)
	);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
