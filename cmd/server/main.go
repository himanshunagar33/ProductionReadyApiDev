package main

import (
	"fmt"
	"net/http"

	"github.com/himanshunagar33/ProductionReadyApiDev/internal/comments"
	"github.com/himanshunagar33/ProductionReadyApiDev/internal/database"
	transporthttp "github.com/himanshunagar33/ProductionReadyApiDev/internal/transport/http"
)

// App - struct which contains things like pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	commentService := comments.NewService(db)

	handler := transporthttp.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to set up  server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API dev")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
