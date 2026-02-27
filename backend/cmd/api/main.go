package main

import (
	"database/sql"
	"log"

	"myapp/internal/http"
	"myapp/internal/http/handlers"
	"myapp/internal/repository/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "host=127.0.0.1 port=5432 user=postgres password=pequena@@ dbname=monitence_db sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	userRepo := postgres.NewUserRepository(db)
	handler := handlers.NewHandler(userRepo)

	// se seu Handler NÃO tiver NewHandler, deixa como você já usa:
	// handler := &handlers.Handler{}
	// mas o ideal é injetar:
	http.RegisterRoutes(r, handler)

	log.Println("API rodando em http://localhost:8081")
	r.Run(":8081")
}