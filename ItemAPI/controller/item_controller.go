package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labora-api/ItemAPI/service"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items := service.GetAllItems()

	w.WriteHeader(http.StatusOK)
	itemsJson, _ := json.Marshal(items)
	w.Write(itemsJson)
}