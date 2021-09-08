package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/models"
)

func (h *HTTPHandler) AddNewsHandler(w http.ResponseWriter, r *http.Request) {
	var news models.News
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	err = json.Unmarshal(req, &news)
	if err != nil {

	}
	h.database.AddData(fmt.Sprintf("INSERT INTO news (Title,Topic,Status) VALUES ('%s','%s','%s');", news.Title, news.Topic, news.Status))
	for _, tag := range news.Tags {
		h.database.AddData(fmt.Sprintf("INSERT INTO tags (Tag,Title) VALUES ('%s','%s'); ", tag, news.Title))
	}
}
