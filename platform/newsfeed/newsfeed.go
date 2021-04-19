package newsfeed

import "errors"

type Getter interface {
	GetAll() []*Item
	FindItem(string) (*Item, error)
}

type Adder interface {
	Add(item *Item)
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items map[string]*Item
}

func New() *Repo {
	return &Repo{
		Items: map[string]*Item{},
	}
}

func (r *Repo) Add(item *Item) {
	r.Items[item.Title] = item
}

func (r *Repo) GetAll() []*Item {
	items := make([]*Item, 0, len(r.Items))
	for _, item := range r.Items {
		items = append(items, item)
	}

	return items
}

func (r *Repo) FindItem(title string) (*Item, error) {
	foundItem, ok := r.Items[title]

	if !ok {
		return nil, errors.New("Record not found")
	}

	return foundItem, nil
}
