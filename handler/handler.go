package handler

import (
	"fmt"
	"gobook/controller"
	"gobook/helper"
	"gobook/models/book"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
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

func (h Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" && strings.TrimSuffix(r.URL.Path, "/") != "/book" {
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

func (h Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "create.html"), path.Join("views/layout", "layout.html"))
	if err != nil {
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}

	wd := Meta{Title: "Create Book"}
	respon := Response{
		WebData: wd,
	}

	err = tmpl.Execute(w, respon)
	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}
}

func (h Handler) StoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Http Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}

	price, _ := strconv.Atoi(r.Form.Get("price"))

	newBook := book.Book{
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		Price:       price,
	}

	h.Controller.StoreBook(&newBook)

	http.Redirect(w, r, "/book", http.StatusMovedPermanently)
}

func (h Handler) EditHandler(w http.ResponseWriter, r *http.Request) {
	idQuery := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 1 {
		http.Error(w, "ID not Valid", http.StatusUnprocessableEntity)
		return
	}

	book, err := h.Controller.ShowBook(id)

	if err != nil {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "edit.html"), path.Join("views/layout", "layout.html"))
	if err != nil {
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	title := fmt.Sprintf("Edit: #%d %s book's", id, book.Title)
	wd := Meta{Title: title}
	respon := Response{
		WebData: wd,
		Data:    book,
	}

	err = tmpl.Execute(w, respon)
	if err != nil {
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}
}

func (h Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}

	price, _ := strconv.Atoi(r.Form.Get("price"))
	id, _ := strconv.Atoi(r.Form.Get("id"))

	updateBook := book.Book{
		ID:          id,
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		Price:       price,
	}

	h.Controller.UpdateBook(&updateBook)

	http.Redirect(w, r, "/book", http.StatusMovedPermanently)
}
