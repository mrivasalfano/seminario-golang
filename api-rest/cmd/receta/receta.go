package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrivasalfano/seminario-golang/api-rest/internal/config"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/database"
	"github.com/mrivasalfano/seminario-golang/api-rest/internal/service/receta"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := receta.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
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
