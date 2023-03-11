package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"estimator/cmd/api/config"
	apictx "estimator/cmd/api/context"
	v1 "estimator/cmd/api/v1"
	"estimator/services"
	"estimator/storage/mysql"

	"github.com/beeker1121/httprouter"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Parse the API configuration file.
	cfg, err := config.ParseConfigFile("config.json")
	if err != nil {
		panic(err)
	}

	// Get the configuration environment variables.
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPass = os.Getenv("DB_PASS")
	cfg.APIHost = os.Getenv("API_HOST")
	cfg.APIPort = os.Getenv("API_PORT")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	// TODO: Add logger.

	// Connect to the MySQL database.
	db, err := sql.Open("mysql", cfg.DBUser+":"+cfg.DBPass+"@tcp("+cfg.DBHost+":"+cfg.DBPort+")/"+cfg.DBName+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test database connection.
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Create a new MySQL storage implementation.
	store := mysql.New(db)

	// Create new services.
	serv := services.New(store)

	// Create a new router.
	router := httprouter.New()

	// Create a new API context.
	ac := apictx.New(serv)

	// Create a new v1 API.
	v1.New(ac, router)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Running server...")

	// Start the HTTP server.
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
