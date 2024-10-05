package models

import (
	"encoding/json"
	"io"
	"strings"
)

type Song struct {
	ID          int64  `json:"id"`
	GroupName   string `json:"group_name"`
	SongName    string `json:"song_name"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (s *Song) ParseRows(values []interface{}) {
	s.ID = values[0].(int64)
	s.GroupName = values[1].(string)
	s.SongName = values[2].(string)
	s.ReleaseDate = values[3].(string)
	s.Text = values[4].(string)
	s.Link = values[5].(string)
}

func (s *Song) DecodeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&s)

	return err
}

func (s *Song) EncodeJSON() ([]byte, error) {
	body, err := json.Marshal(s)

	return body, err
}

func (s *Song) Merge(songOld Song) {
	if s.GroupName == "" {
		s.GroupName = songOld.GroupName
	}
	if s.SongName == "" {
		s.SongName = songOld.SongName
	}
	if s.ReleaseDate == "" {
		s.ReleaseDate = songOld.ReleaseDate
	}
	if s.Text == "" {
		s.Text = songOld.Text
	}
	if s.Link == "" {
		s.Link = songOld.Link
	}
}

func (s *Song) PaginationForText(page map[string]int) []string {
	var (
		indexBegin int
		indexEnd   int
		arrayText  []string
	)
	verse := "["
	text := s.Text
	indexBegin = strings.Index(text, verse)
	for i := page["offset"] + 1; i <= page["limit"]; i++ {
		indexBegin = strings.Index(text, verse)
		if indexBegin == -1 {
			break
		}
		indexEnd = strings.Index(text[indexBegin:], "\n\n")
		text = text[indexBegin:]
		arrayText = append(arrayText, strings.TrimSpace(text[:indexEnd]))
		text = text[indexEnd+2:]
	}

	return arrayText
}
