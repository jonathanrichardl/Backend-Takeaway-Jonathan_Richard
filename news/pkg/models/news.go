package models

import "encoding/json"

type News struct {
	Title  string   `json:"news"`
	Topic  string   `json:"topic"`
	Tags   []string `json:"tags"`
	Status string   `json:"status"`
}

func (a *News) Unmarshal(data []byte) (*News, error) {
	var NewNews News
	err := json.Unmarshal(data, &NewNews)
	return &NewNews, err
}
