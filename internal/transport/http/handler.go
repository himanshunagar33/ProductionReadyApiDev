package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//handler -stores the pointer to our comments service
type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

// setupRoutes -sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up the routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive !! ")
	})
}
