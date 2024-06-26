package main

import (
	"fmt"
	"net/http"

	"github.com/MaheshMoholkar/lenslocked/controllers"
	"github.com/MaheshMoholkar/lenslocked/models"
	"github.com/MaheshMoholkar/lenslocked/templates"
	"github.com/MaheshMoholkar/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "layout-page.html", "home.html"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.html", "about.html"))
	r.Get("/about", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.html", "faq.html"))
	r.Get("/faq", controllers.FAQ(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	us := models.UserService{
		DB: db,
	}

	uc := controllers.Users{
		UserService: &us,
	}

	uc.Templates.New = views.Must(views.ParseFS(templates.FS, "layout-page.html", "signup.html"))
	uc.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "layout-page.html", "signin.html"))

	r.Get("/signup", uc.New)
	r.Get("/signin", uc.SignIn)
	r.Post("/users", uc.Create)
	r.Post("/signin", uc.ProcessSignIn)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	csrfKey := "gFvi35R4fy5xNBlnEeZtQbfAVCYeIAUC"
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false))
	fmt.Println("server started:3000")
	http.ListenAndServe(":3000", csrfMw(r))
}
