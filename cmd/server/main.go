package main

import (
	"fmt"
	"net/http"

	transporthttp "github.com/himanshunagar33/ProductionReadyApiDev"
)

// App - struct which contains things like pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")
	handler := transporthttp.NewHandler()
	handler.setupRoutes()
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
