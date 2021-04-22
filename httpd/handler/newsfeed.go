package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/pyankovzhe/chi-router/platform/newsfeed"
)

func NewsfeedList(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()

		render.RenderList(w, r, NewsfeedResponse(items))
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

		render.Render(w, r, &itemRequest{Item: item})
	}
}

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &itemRequest{}

		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: err.Error()})
			return
		}

		item := feed.Add(data.Item)
		render.Status(r, http.StatusCreated)
		render.Render(w, r, &itemRequest{Item: item})
	}
}

type itemRequest struct {
	*newsfeed.Item
}

func (i *itemRequest) Bind(r *http.Request) error {
	// a.Item is nil if no Item fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if i.Item == nil {
		return errors.New("missing required Item fields.")
	}

	return nil
}

func (i *itemRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewsfeedResponse(items []*newsfeed.Item) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		list = append(list, &itemRequest{Item: item})
	}
	return list
}

func parseInt32FromParam(s string) (int32, error) {
	var id64, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(id64), nil
}
