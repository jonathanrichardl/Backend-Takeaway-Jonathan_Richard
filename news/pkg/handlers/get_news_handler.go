package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetAllNewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cache, err := h.redis.GetData("news")
	if err != nil {
		retrievedData, err := h.database.RetrieveData("SELECT title,topic,status FROM news")
		if err != nil {
			h.logger.ErrorLogger.Println("Can't retrieve news from SQL: ", err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response := h.retrieveNews(retrievedData)
		resp, _ := json.Marshal(response)
		err = h.redis.SetData("news", string(resp))
		if err != nil {
			h.logger.ErrorLogger.Println("Error saving cache to redis,", err.Error())
		}
		fmt.Fprintf(w, string(resp))
		return
	}
	fmt.Fprintf(w, cache)

}

func (h *HTTPHandler) GetNewsByTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := mux.Vars(r)
	topicName := strings.ToLower(request["topic"])
	cache, err := h.redis.GetData(topicName)
	if err != nil {
		retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,status,topic FROM news WHERE topic = '%s';", topicName))
		if err != nil {
			h.logger.ErrorLogger.Println("Can't retrieve news: ", err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response := h.retrieveNews(retrievedData)
		resp, _ := json.Marshal(response)
		err = h.redis.SetData(topicName, string(resp))
		if err != nil {
			h.logger.ErrorLogger.Println("Error saving cache to redis,", err.Error())
		}
		fmt.Fprintf(w, string(resp))
		return
	}
	fmt.Fprintf(w, cache)
}

func (h *HTTPHandler) GetNewsByTitleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := mux.Vars(r)
	newsName := strings.ToLower(request["title"])
	cache, err := h.redis.GetData(newsName)
	if err != nil {
		retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,topic,status FROM news WHERE title = '%s';", newsName))
		if err != nil {
			h.logger.ErrorLogger.Println("Can't retrieve news: ", err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response := h.retrieveNews(retrievedData)
		resp, _ := json.Marshal(response)
		err = h.redis.SetData(newsName, string(resp))
		if err != nil {
			h.logger.ErrorLogger.Println("Error saving cache to redis,", err.Error())
			return
		}
		fmt.Fprintf(w, string(resp))
		return
	}
	fmt.Fprintf(w, cache)

}

func (h *HTTPHandler) GetNewsByStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp []byte
	request := mux.Vars(r)
	status := strings.ToLower(request["status"])
	cache, err := h.redis.GetData(status)
	if err != nil {
		switch status {
		case "deleted":
			retrievedData, err := h.database.RetrieveData("SELECT * FROM deleted;")
			if err != nil {
				h.logger.ErrorLogger.Println("Can't retrieve deleted news: ", err.Error())
				w.WriteHeader(http.StatusNotFound)
				return
			}
			response := h.retrieveDeletedNews(retrievedData)
			resp, _ = json.Marshal(response)

		default:
			retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,topic,status FROM news WHERE status = '%s';", status))
			if err != nil {
				h.logger.ErrorLogger.Println("Can't retrieve data: ", err.Error())
				w.WriteHeader(http.StatusNotFound)
				return
			}
			response := h.retrieveNews(retrievedData)
			resp, _ = json.Marshal(response)
		}
		err = h.redis.SetData(status, string(resp))
		if err != nil {
			h.logger.ErrorLogger.Println("Error saving cache to redis,", err.Error())
		}
		fmt.Fprintf(w, string(resp))
		return
	}
	fmt.Fprintf(w, cache)

}
