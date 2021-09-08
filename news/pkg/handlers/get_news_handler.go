package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news/pkg/models"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetAllNewsHandler(w http.ResponseWriter, r *http.Request) {
	var response []models.News
	retrievedData, err := h.database.RetrieveData("SELECT * FROM news")
	if err != nil {

	}
	for retrievedData.Data.Next() {
		var each models.News
		err := retrievedData.Data.Scan(&each.Title, &each.Tags, &each.Topic)
		if err != nil {

		}
		response = append(response, each)
	}
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetNewsByTopicHandler(w http.ResponseWriter, r *http.Request) {
	var response []models.News
	request := mux.Vars(r)
	topicName := request["name"]
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM NEWS WHERE TOPIC = %s;", topicName))
	if err != nil {

	}
	for retrievedData.Data.Next() {
		var each models.News
		err := retrievedData.Data.Scan(&each.Title, &each.Tags, &each.Topic)
		if err != nil {

		}
		response = append(response, each)
	}
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetNewsByNameHandler(w http.ResponseWriter, r *http.Request) {
	var response []models.News
	request := mux.Vars(r)
	newsName := request["name"]
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM news WHERE name = %s;", newsName))
	if err != nil {

	}
	for retrievedData.Data.Next() {
		var each models.News
		err := retrievedData.Data.Scan(&each.Title, &each.Tags, &each.Topic)
		if err != nil {

		}
		response = append(response, each)
	}
	json.NewEncoder(w).Encode(response)
}
