package main

import (
	"os"

	"github.com/undercode99/restql/db"
	"github.com/undercode99/restql/http_rest"
)

func main() {

	// load .env file

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8863"
	}

	apiApp := http_rest.NewApi(&http_rest.ApiConfig{
		Port: appPort,
		Mode: os.Getenv("APP_MODE"),
	}, db.NewListDatabaseConnect())

	apiApp.Run()
}
