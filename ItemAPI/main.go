package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labora-api/ItemAPI/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetItems).Methods("GET")

	http.ListenAndServe(":8000", router)

}