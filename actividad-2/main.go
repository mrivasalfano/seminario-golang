package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sqlx.DB

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	createSchema(db)

	// r := gin.Default()
	// r.GET("/users", getUsersHandler)
	// r.GET("/users/:name", getUserNameHandler)
	// r.POST("/users", postUsersHandler)
	// r.Run()
}

//Receta ...
type Receta struct {
	ID         int    `db:"id"`
	Nombre     string `db:"nombre"`
	Duracion   int    `db:"duracion"`
	Dificultad string `db:"dificultad"`
}

func createSchema(db *sqlx.DB) {
	schema := `CREATE TABLE receta (
		id integer PRIMARY KEY AUTOINCREMENT,
		nombre varchar(30),
		duracion integer,
		dificultad varchar(10));`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err.Error())
	}

	db.MustExec("INSERT INTO receta(nombre, duracion, dificultad) VALUES (?,?,?)",
		"Tortilla", 20, "Media")
	rows, err := db.Queryx("SELECT * FROM receta")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var r Receta
		rows.StructScan(&r)
		fmt.Println(r)
	}
}

func getUsersHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func getUserNameHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   c.Param("name"),
		"edad":   c.Query("edad"),
	})
}

type user struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func postUsersHandler(c *gin.Context) {
	requestBody := user{}
	c.Bind(&requestBody)
	user := user{
		Nombre: requestBody.Nombre,
		Edad:   requestBody.Edad,
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, user)
}
