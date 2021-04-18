package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pyankovzhe/chi-router/httpd/handler"
	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

type server struct {
	router *chi.Mux
	store  *newsfeed.Repo
}

func (s *server) configureRouter() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Get("/newsfeed", handler.NewsfeedGet(s.store))
	s.router.Post("/newsfeed", handler.NewsfeedPost(s.store))
	s.router.Get("/panic-test", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})
}

func NewServer() *server {
	s := &server{
		router: chi.NewRouter(),
		store:  newsfeed.New(),
	}

	s.configureRouter()

	return s
}

func main() {
	srv := NewServer()
	http.ListenAndServe(":3000", srv.router)
}
