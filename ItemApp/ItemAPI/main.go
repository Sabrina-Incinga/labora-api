package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labora-api/ItemAPI/config"
	"github.com/labora-api/ItemAPI/controller"
	"github.com/labora-api/ItemAPI/repository"
	"github.com/labora-api/ItemAPI/service"
	"github.com/rs/cors"
)

func main() {
	var db, error = config.GetConnection()
	if error != nil {
		log.Fatal(error)
	}
	dbHandler := &repository.PostgresDBHandler{Db: db}
	itemService := &service.ItemService{DbHandler: dbHandler}
	controller := &controller.ItemController{ItemServiceImpl: *itemService}

	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetItems).Methods("GET")
	router.HandleFunc("/items/getById/{id}", controller.GetById).Methods("GET")
	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items/update/{id}", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/delete/{id}", controller.DeleteItem).Methods("DELETE")

	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// Agregar el middleware CORS a todas las rutas
	handler := corsOptions.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
