package main

import (
	"net/http"

	"github.com/kavit-munshi/bookings/pkg/config"
	"github.com/kavit-munshi/bookings/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	/* 	mux := pat.New()

	   	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	   	mux.Get("/about/", http.HandlerFunc(handlers.Repo.About))
	   	mux.Get("/greet/:name", http.HandlerFunc(handlers.Repo.Greet))
	   	mux.Get("/greet/", http.HandlerFunc(handlers.Repo.Greet)) */

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about/", http.HandlerFunc(handlers.Repo.About))
	//mux.Get("/greet/:name", http.HandlerFunc(handlers.Repo.Greet))
	mux.Get("/greet/", http.HandlerFunc(handlers.Repo.Greet))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
