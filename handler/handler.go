package handler

import (
	"gobook/controller"
	"gobook/helper"
	"html/template"
	"net/http"
	"path"
)

type Handler struct {
	Controller *controller.Controller
}

type Meta struct {
	Title string
}

type Response struct {
	WebData Meta
	Data    interface{}
}

func NewHandler(c *controller.Controller) *Handler {
	return &Handler{c}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// get books list from database
	books := h.Controller.GetBooks()

	// add additional function to gohtml
	funcs := template.FuncMap{"add": helper.Add}

	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles(path.Join("views", "index.html"), path.Join("views/layout", "layout.html"))
	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	wd := Meta{Title: "Book List"}
	respon := Response{
		WebData: wd,
		Data:    books,
	}

	// execute the template
	err = tmpl.Execute(w, respon)
	// err = tmpl.Execute(w, books)

	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {

}
