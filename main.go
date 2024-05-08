package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/MaheshMoholkar/lenslocked/controllers"
	"github.com/MaheshMoholkar/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.html")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "about.html")))
	r.Get("/about", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Print("Server started:3000\n")
	http.ListenAndServe(":3000", r)
}
