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

func (r Route) InitRoute() {
	mux := http.NewServeMux()

	r.RouteList(mux)

	log.Println("Starting web on localhost:8000")

	err := http.ListenAndServe("localhost:8000", mux)
	log.Fatal(err)
}

func (r Route) RouteList(mux *http.ServeMux) {
	// Index
	mux.HandleFunc("/", r.Handler.IndexHandler)
	mux.HandleFunc("/book", r.Handler.IndexHandler)

	// Create
	mux.HandleFunc("/book/create", r.Handler.CreateHandler)
	mux.HandleFunc("/book/store", r.Handler.StoreHandler)

	// edit
	mux.HandleFunc("/book/edit", r.Handler.EditHandler)
	mux.HandleFunc("/book/update", r.Handler.UpdateHandler)

	// delete
	mux.HandleFunc("/book/delete", r.Handler.DeleteHandler)
}
