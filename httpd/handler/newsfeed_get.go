package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}

func NewsfeedShow(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item newsfeed.Item

		if itemTitle := chi.URLParam(r, "itemTitle"); itemTitle != "" {
			item = feed.FindItem(itemTitle)
		} else {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}
