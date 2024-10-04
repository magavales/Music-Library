package models

import (
	"github.com/doug-martin/goqu/v9"
	"net/url"
)

type Parameter struct {
	page   *Pagination
	filter *Filtration
}

func NewParameter() *Parameter {
	return &Parameter{
		page:   newPagination(),
		filter: newFiltration(),
	}
}

func (p *Parameter) Parse(url *url.URL) {
	p.page.parse(url)
	p.filter.parse(url)
}

func (p *Parameter) GetSQL() (string, error) {
	var (
		filter goqu.Ex
		page   map[string]int
		sql    string
		err    error
	)
	filter, err = p.filter.get()
	if err != nil {
		page, err = p.page.get()
		if err != nil {
			return "", err
		} else {
			sql, _, err = goqu.From("music_library").Limit(uint(page["limit"])).Offset(uint(page["offset"])).ToSQL()
			return sql, nil
		}
	} else {
		page, err = p.page.get()
		if err != nil {
			sql, _, err = goqu.From("music_library").Where(filter).ToSQL()
			return sql, nil
		} else {
			sql, _, err = goqu.From("music_library").Where(filter).Limit(uint(page["limit"])).Offset(uint(page["offset"])).ToSQL()
			return sql, nil
		}
	}
}

func (p *Parameter) GetPage() (map[string]int, error) {
	return p.page.get()
}
