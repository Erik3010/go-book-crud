package handler

import (
	"gobook/controller"
	"html/template"
	"net/http"
	"path"
)

type Handler struct {
	Controller *controller.Controller
}

func NewHandler(c *controller.Controller) *Handler {
	return &Handler{c}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	books := h.Controller.GetBooks()

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, books)

	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {

}
