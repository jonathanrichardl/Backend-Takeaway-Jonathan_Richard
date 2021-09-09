package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"news/pkg/models"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) ModifyNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsUpdate models.NewsUpdate
	request := mux.Vars(r)
	newsTitle := request["title"]
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading modify news request: ", err.Error())

	}
	err = json.Unmarshal(req, &newsUpdate)
	if err != nil {
		log.Println("Error unmarshaling modify news request: ", err.Error())
	}
	if newsUpdate.Title != nil {
		h.modifyTitle(newsTitle, newsUpdate.Title)
		newsTitle = *newsUpdate.Title
	}
	if newsUpdate.Topic != nil {
		h.modifyTopic(newsTitle, newsUpdate.Topic)
	}
	if newsUpdate.Tags != nil {
		h.modifyTags(newsTitle, newsUpdate.Tags)
	}
	if newsUpdate.Status != nil {
		h.modifyStatus(newsTitle, newsUpdate.Status)
	}

}
