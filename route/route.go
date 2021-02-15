package route

import (
	"gobook/handler"
	"log"
	"net/http"
)

type Route struct {
	Handler *handler.Handler
}

func NewRoute(h *handler.Handler) *Route {
	return &Route{h}
}

func (r *Route) InitRoute() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.Handler.IndexHandler)
	// mux.HandleFunc("/create", r.Handler.CreateHandler)

	log.Println("Starting web on localhost:8000")

	err := http.ListenAndServe("localhost:8000", mux)
	log.Fatal(err)
}
