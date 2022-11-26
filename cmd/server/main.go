package main

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/configs"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, MysqlConnectString())
	if err != nil {
		log.Fatalf("Error initializing db: %s", err)
		panic(err)
	}

	userDb := database.NewUserDb(db)
	announcementDb := database.NewAnnouncementDb(db)
	announcementImageDb := database.NewAnnouncementImagesDb(db)
	announcementHandler := handlers.NewAnnouncementHandler(announcementDb, announcementImageDb)
	userHandler := handlers.NewUserHandler(userDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.FindByEmail)
	r.Post("/announcements", announcementHandler.CreateAnnouncement)
	r.Get("/announcements", announcementHandler.GetAllAnnouncements)

	http.ListenAndServe(configs.WebServerPort, r)
}

func MysqlConnectString() string {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	cfg := mysql.Config{
		User:                 configs.MysqlUser,     // Username
		Passwd:               configs.MysqlPassword, // Password (requires User)
		Net:                  "tcp",                 // Network type
		Addr:                 configs.DBHost,        // Network address (requires Net)
		DBName:               configs.MysqlDatabase, // Database name
		Collation:            "utf8mb4_general_ci",  // Connection collation
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		Timeout:              time.Second * 5,
		Loc:                  time.Local,
		ParseTime:            true,
	}

	return cfg.FormatDSN()

}
