package models

import (
	"encoding/json"
	"fmt"
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
	verse := "[Куплет]"
	indexBegin = strings.Index(s.Text, verse)
	if indexBegin != -1 {
		indexEnd = strings.Index(s.Text[indexBegin:], "\n\n")
		tempText := s.Text[indexBegin:]
		arrayText = append(arrayText, strings.TrimSpace(tempText[:indexEnd]))
		return arrayText
	} else {
		for i := page["offset"] + 1; i <= page["limit"]; i++ {
			verse = fmt.Sprintf("[Куплет %d", i)
			indexBegin = strings.Index(s.Text, verse)
			if indexBegin == -1 {
				break
			}
			indexEnd = strings.Index(s.Text[indexBegin:], "\n\n")
			tempText := s.Text[indexBegin:]
			arrayText = append(arrayText, strings.TrimSpace(tempText[:indexEnd]))
		}
	}

	return arrayText
}
