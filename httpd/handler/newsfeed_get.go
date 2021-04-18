package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
