package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &itemRequest{}

		if err := render.Bind(r, data); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		feed.Add(data.Item)
		w.Write([]byte("OK"))
	}
}

type itemRequest struct {
	*newsfeed.Item
}

func (i *itemRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if i.Item == nil {
		return errors.New("missing required Item fields.")
	}

	return nil
}
