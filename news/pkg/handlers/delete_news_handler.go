package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	newsTitle := request["title"]
	err := h.database.AddData(fmt.Sprintf("INSERT INTO deleted SELECT id,title FROM news WHERE title = '%s';", newsTitle))
	if err != nil {
		h.logger.ErrorLogger.Println("Error moving news into deleted table: ", err.Error())
	}
	err = h.database.DeleteData(fmt.Sprintf("DELETE FROM tags WHERE news_id = (SELECT id FROM news WHERE title = '%s');", newsTitle))
	if err != nil {
		h.logger.ErrorLogger.Println("Failed to delete tags: ", err.Error())
	}
	err = h.database.DeleteData(fmt.Sprintf("DELETE FROM news WHERE title = '%s';", newsTitle))
	if err != nil {
		h.logger.ErrorLogger.Println("Failed to delete news: ", err.Error())
	}
	w.WriteHeader(204)
	w.Write([]byte(fmt.Sprintf("204 - News '%s' has been deleted", newsTitle)))
	h.logger.InfoLogger.Printf("News %s has been deleted from database", newsTitle)
}
