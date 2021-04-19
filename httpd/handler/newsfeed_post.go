package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

type itemRequest struct {
	*newsfeed.Item
}

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request := map[string]string{}
		// json.NewDecoder(r.Body).Decode(&request)
		data := &itemRequest{}

		if err := render.Bind(r, data); err != nil {
			http.Error(w, http.StatusText("Invalid request."), 400)
			return
		}

		feed.Add(data.Item)
		w.Write([]byte("OK"))
	}
}
