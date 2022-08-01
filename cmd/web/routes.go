package main

import (
	"net/http"
	"pkg/pkg/config"
	"pkg/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// Se establece un middleware propio y basico
	//mux.Use(WriteToConsole)

	//se establece el middleware de Chi
	mux.Use(middleware.Recoverer)

	mux.Use(NoSurf)
	mux.Use((SessionLoad))
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	// se establece el directorio ./static como servidor de archivos
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
