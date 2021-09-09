package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"news/pkg/models"
)

func (h *HTTPHandler) AddNewsHandler(w http.ResponseWriter, r *http.Request) {
	var news models.News
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading add news request: ", err.Error())

	}
	err = json.Unmarshal(req, &news)
	if err != nil {
		log.Println("Error unmarshalling request data: ", err.Error())

	}
	if h.database.CheckIfExists(fmt.Sprintf("SELECT EXISTS(SELECT * FROM news WHERE title = '%s')", news.Title)) {
		log.Printf("News with title %s already exists!\n", news.Title)
		return
	}
	h.database.AddData(fmt.Sprintf("INSERT INTO news (Title,Topic,Status) VALUES ('%s','%s','%s');", news.Title, news.Topic, news.Status))
	h.addTags(news.Title, news.Tags)
}
