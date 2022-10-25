package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // add this

	"github.com/undercode99/restql/db"
	"github.com/undercode99/restql/http_rest"
)

func main() {

	// load .env file
	godotenv.Load()

	dbConnect, err := db.NewDatabaseConnect(&db.Database{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Driver:   "postgres",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer dbConnect.Close()

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8863"
	}

	apiApp := http_rest.NewApi(&http_rest.ApiConfig{
		Port: appPort,
		Mode: os.Getenv("APP_MODE"),
	}, dbConnect)

	apiApp.Run()
}
