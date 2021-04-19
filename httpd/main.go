package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pyankovzhe/chi-router/httpd/handler"
	"github.com/pyankovzhe/chi-router/platform/newsfeed"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *chi.Mux
	store  *newsfeed.Repo
	logger *logrus.Logger
}

func (s *server) configureRouter() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: s.logger, NoColor: false}))
	s.router.Use(middleware.Recoverer)

	s.router.Get("/newsfeed", handler.NewsfeedGet(s.store))
	s.router.Post("/newsfeed", handler.NewsfeedPost(s.store))
	s.router.Get("/newsfeed/items/{itemTitle}", handler.NewsfeedShow(s.store))
	s.router.Get("/panic-test", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})
}

func NewServer(logger *logrus.Logger, store *newsfeed.Repo) *server {
	s := &server{
		router: chi.NewRouter(),
		store:  store,
		logger: logger,
	}

	s.configureRouter()

	s.logger.Info("Server initialized")
	return s
}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	store := newsfeed.New()
	srv := NewServer(logger, store)
	logger.Info("Server started")
	http.ListenAndServe(":3000", srv.router)
}
