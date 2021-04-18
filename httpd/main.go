package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pyankovzhe/chi-router/httpd/handler"
	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	http.ListenAndServe(":3000", r)
}
