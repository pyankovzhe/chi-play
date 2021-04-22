package newsfeed

import (
	"errors"
)

type Getter interface {
	GetAll() []*Item
	FindItem(int32) (*Item, error)
}

type Adder interface {
	Add(item *Item) *Item
}

type Item struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items map[int32]*Item
}

func New() *Repo {
	return &Repo{
		Items: map[int32]*Item{},
	}
}

func (r *Repo) Add(item *Item) *Item {
	item.ID = int32(len(r.Items) + 1)
	r.Items[item.ID] = item

	return item
}

func (r *Repo) GetAll() []*Item {
	items := make([]*Item, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, item)
	}

	return items
}

func (r *Repo) FindItem(id int32) (*Item, error) {
	foundItem, ok := r.Items[id]

	if !ok {
		return nil, errors.New("Record not found")
	}

	return foundItem, nil
}
