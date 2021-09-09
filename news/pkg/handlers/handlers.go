package handlers

import (
	"news/pkg/database"
	"news/pkg/router"
)

type HTTPHandler struct {
	database *database.DBInstance
	router   *router.RouterInstance
}

func NewHttpHandlers(DatabaseInstance *database.DBInstance, RouterInstance *router.RouterInstance) *HTTPHandler {
	return &HTTPHandler{
		database: DatabaseInstance,
		router:   RouterInstance,
	}
}

func (h *HTTPHandler) RegisterAllHandlers() {
	h.router.RegisterHandler("/news", h.AddNewsHandler, "POST")
	h.router.RegisterHandler("/news", h.GetAllNewsHandler, "GET")
	h.router.RegisterHandler("/news/{title}", h.GetNewsByTitleHandler, "GET")
	h.router.RegisterHandler("/news/{title}", h.DeleteNewsHandler, "DELETE")
	h.router.RegisterHandler("/news/{title}", h.ModifyNewsHandler, "PATCH")
	h.router.RegisterHandler("/news/topic/{topic}", h.GetNewsByTopicHandler, "GET")
	h.router.RegisterHandler("/news/status/{status}", h.GetNewsByStatusHandler, "GET")
}
