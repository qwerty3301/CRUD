package main

import (
	"MyCRUD/internal/handlers"
	"MyCRUD/internal/repository"
	"MyCRUD/internal/server"
	"MyCRUD/internal/service"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	config, err := repository.LoadConfig("config/config.toml")
	if err != nil {
		log.Fatalf("Error loading configuration file: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
	config.Database.Password = os.Getenv("DB_PASSWORD")

	db := &repository.DatabaseConfig{}
	db.NewPostgreDB(&config.Database)
	conn := db.Connect()

	repos := repository.UserRepo{Connection: conn}
	serv := service.UserService{Repo: &repos}
	handl := handlers.UserHandler{Service: &serv}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(conn, context.Background())

	s := &server.Server{Port: ":8080", Handler: handl}
	router := s.Create()

	s.Start(router)

}
