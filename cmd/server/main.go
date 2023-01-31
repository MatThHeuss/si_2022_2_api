package main

import (
	"database/sql"
	"fmt"
	"github.com/MatThHeuss/si_2020_2_api/configs"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("successfully connected to the database.")
	defer db.Close()

	userDb := database.NewUserDb(db)
	announcementDb := database.NewAnnouncementDb(db)
	announcementImageDb := database.NewAnnouncementImagesDb(db)
	chatDb := database.NewChatDb(db)

	announcementHandler := handlers.NewAnnouncementHandler(announcementDb, announcementImageDb)
	userHandler := handlers.NewUserHandler(userDb)
	chatHandler := handlers.NewChatHandler(chatDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.FindByEmail)
	r.Post("/announcements", announcementHandler.CreateAnnouncement)
	r.Get("/announcements", announcementHandler.GetAllAnnouncements)
	r.Get("/announcements/{id}", announcementHandler.GetAnnouncementById)
	r.Get("/announcements/category/{category}", announcementHandler.GetAnnouncementByCategory)
	r.Post("/chat", chatHandler.Create)
	r.Get("/chat", chatHandler.GetAllMessage)

	fmt.Printf("Starting Server at port: %s\n", configs.WebServerPort)
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
		Timeout:              time.Second * 10,
		Loc:                  time.Local,
		ParseTime:            true,
	}

	return cfg.FormatDSN()

}
