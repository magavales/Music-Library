package models

import (
	"errors"
	"net/url"
)

type Filtration struct {
	params map[string]interface{}
}

func newFiltration() *Filtration {
	return &Filtration{
		params: make(map[string]interface{}),
	}
}

func (f *Filtration) parse(url *url.URL) {
	if url.Query().Has("id") {
		f.params["id"] = url.Query().Get("id")
	}
	if url.Query().Has("group_name") {
		f.params["group_name"] = url.Query().Get("group_name")
	}
	if url.Query().Has("song_name") {
		f.params["song_name"] = url.Query().Get("song_name")
	}
	if url.Query().Has("release_date") {
		f.params["release_date"] = url.Query().Get("release_date")
	}
	if url.Query().Has("text") {
		f.params["text"] = url.Query().Get("text")
	}
	if url.Query().Has("link") {
		f.params["link"] = url.Query().Get("link")
	}
}

func (f *Filtration) get() (map[string]interface{}, error) {
	var err error
	if len(f.params) == 0 {
		err = errors.New("no params for filtering")
		return nil, err
	}
	return f.params, err
}
