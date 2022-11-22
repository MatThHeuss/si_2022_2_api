package main

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/matheus/Desktop/si-projeto-backend/si-backend.json")
	db, err := sql.Open("mysql", MysqlConnectString())
	if err != nil {
		log.Fatalf("Error initializing db: %s", err)
		panic(err)
	}

	userDb := database.NewUserDb(db)
	userHandler := handlers.NewUserHandler(userDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8000", r)
}

func MysqlConnectString() string {
	cfg := mysql.Config{
		User:                 "root",               // Username
		Passwd:               "root",               // Password (requires User)
		Net:                  "tcp",                // Network type
		Addr:                 "localhost:3306",     // Network address (requires Net)
		DBName:               "si-backend",         // Database name
		Collation:            "utf8mb4_general_ci", // Connection collation
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		Timeout:              time.Second * 5,
		Loc:                  time.Local,
		ParseTime:            true,
	}

	return cfg.FormatDSN()

}
