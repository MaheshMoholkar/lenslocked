package main

import (
	"fmt"
	"net/http"

	"github.com/MaheshMoholkar/lenslocked/controllers"
	"github.com/MaheshMoholkar/lenslocked/templates"
	"github.com/MaheshMoholkar/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.html"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "about.html"))
	r.Get("/about", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.html"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Print("Server started:3000\n")
	http.ListenAndServe(":3000", r)
}
