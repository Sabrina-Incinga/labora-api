package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
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


	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Agregar el middleware CORS a todas las rutas
	handler := corsOptions.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}