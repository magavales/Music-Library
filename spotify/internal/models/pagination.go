package models

import (
	"errors"
	"net/url"
	"strconv"
)

type Pagination struct {
	params map[string]int
}

func newPagination() *Pagination {
	return &Pagination{
		params: make(map[string]int),
	}
}

func (pag *Pagination) parse(url *url.URL) {
	if url.Query().Has("limit") {
		pag.params["limit"], _ = strconv.Atoi(url.Query().Get("limit"))
		if url.Query().Has("offset") {
			pag.params["offset"], _ = strconv.Atoi(url.Query().Get("offset"))
		} else {
			pag.params["offset"] = 0
		}
	}
}

func (pag *Pagination) get() (map[string]int, error) {
	var err error
	if len(pag.params) == 0 {
		err = errors.New("query doesn't contain limit and offset")
		return nil, err
	}

	return pag.params, err
}
