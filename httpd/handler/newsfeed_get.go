package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

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
		var item *newsfeed.Item
		var err error

		itemIdStr := chi.URLParam(r, "id")
		if itemIdStr == "" {
			render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: "id is required"})
			return
		}

		itemId, err := parseInt32FromParam(itemIdStr)
		if err != nil {
			render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: err.Error()})
			return
		}

		item, err = feed.FindItem(itemId)
		if err != nil {
			render.Render(w, r, &ErrResponse{Code: http.StatusNotFound, Message: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}

func parseInt32FromParam(s string) (int32, error) {
	var id64, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(id64), nil
}
