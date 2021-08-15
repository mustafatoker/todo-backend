package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	cors "github.com/rs/cors/wrapper/gin"
	"golang-todo-list/controllers"
	"golang-todo-list/routes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDatabase(
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
	)

	migrateDatabase(db)

	if err != nil {
		log.Fatal("err", err)
	}

	todoController := controllers.NewTodoController(db)

	r := gin.Default()

	r.Use(cors.AllowAll())

	routes.SetupRoutes(r, todoController)

	servErr := r.Run()
	if servErr != nil {
		log.Fatal("Server cannot run..")
	}
}

func connectDatabase(user, password, dbname, host, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	return sql.Open("postgres", connStr)
}

func migrateDatabase(db *sql.DB) {
	query, err := ioutil.ReadFile("migrations/create_todos_table.sql")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(string(query)); err != nil {
		panic(err)
	}
}
