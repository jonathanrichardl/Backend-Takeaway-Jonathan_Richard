package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"news/pkg/database"
	"news/pkg/handlers"
	"news/pkg/logger"
	"news/pkg/models"
	"news/pkg/redis"
	"news/pkg/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Init() *handlers.HTTPHandler {
	logger := logger.NewLogger("log.txt")
	logger.InfoLogger.Println("Initializing Program")
	Database, err := database.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")
	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
	}
	router := router.NewRouterInstance()
	redis := redis.NewRedisClient("localhost", 6379, "", 0, 60)
	handlers := handlers.NewHttpHandlers(Database, router, logger, redis)
	return handlers

}
func TestAddNewsHandler(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news", handlers.AddNewsHandler, "POST")
	newNews := models.News{
		Title:  "WhatisGo?",
		Tags:   []string{"Programming", "Go"},
		Topic:  "Technology",
		Status: "Draft",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPost, "/news", payload)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 201, resp.Code, "Created response")

}

func TestRetreiveNews(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/{title}", handlers.GetNewsByTitleHandler, "GET")
	req, _ := http.NewRequest(http.MethodGet, "/news/WhatIsGo", nil)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")
}

func TestDeleteNews(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/{title}", handlers.DeleteNewsHandler, "DELETE")
	req, _ := http.NewRequest(http.MethodDelete, "/news/WhatIsGo", nil)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 204, resp.Code, "Deleted response")

}

func TestRetreiveNewsByTopic(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/topic/{topic}", handlers.GetNewsByTopicHandler, "GET")
	req, _ := http.NewRequest(http.MethodGet, "/news/topic/investation", nil)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestRetreiveNewsByState(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/status/{status}", handlers.GetNewsByStatusHandler, "GET")
	req, _ := http.NewRequest(http.MethodGet, "/news/status/draft", nil)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestModifyExistingNews(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/{title}", handlers.ModifyNewsHandler, "PATCH")
	Title := "NewTitle"
	Tags := []string{"Golang", "Program"}
	Topic := "Programming"
	newNews := models.NewsUpdate{
		Title: &Title,
		Tags:  &Tags,
		Topic: &Topic,
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPatch, "/news/HealthyInvestation", payload)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestAddNewTagsIntoNews(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/{title}/tags", handlers.AddNewTags, "POST")
	newNews := models.TagsUpdate{
		Tags: []string{"Golang", "Program"},
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPost, "/news/WhatIsGo/tags", payload)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestDeleteTagsofNews(t *testing.T) {
	handlers := Init()
	handlers.Router.RegisterHandler("/news/{title}/tags/{tags}", handlers.RemoveTags, "DELETE")
	newNews := models.TagsUpdate{
		Tags: []string{"Golang", "Program"},
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodDelete, "/news/WhatIsGo/tags/Golang", payload)
	resp := httptest.NewRecorder()
	handlers.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 204, resp.Code, "Deleted response")

}
