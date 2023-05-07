package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labora-api/ItemAPI/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", controller.GetById).Methods("GET")
	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")


	http.ListenAndServe(":8000", router)

}