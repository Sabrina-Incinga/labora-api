package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Item struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
}

var items []Item = make([] Item, 10)

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	itemsJson, _ := json.Marshal(items)
	w.Write(itemsJson)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var itemBuscado *Item
	var itemBuscadoNoPuntero Item

	for i := 0; i < len(items); i++ {
		if items[i].ID == vars["id"] {
			itemBuscado = &items[i]
			break
		}
	}

	if itemBuscado == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item no encontrado"))
		return
	}

	itemBuscadoNoPuntero = *itemBuscado
	itemJson, err := json.Marshal(itemBuscado)
	itemNoPuntero, err := json.Marshal(itemBuscadoNoPuntero)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No se pudo convertir el elemento encontrado"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(itemNoPuntero)
	w.Write(itemJson)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
    err := json.NewDecoder(r.Body).Decode(&item)
    defer r.Body.Close()

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	items = append(items, item)

	w.Write([]byte("Item creado correctamente"))
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemUpdate Item
    err := json.NewDecoder(r.Body).Decode(&itemUpdate)
    defer r.Body.Close()
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	for i, item := range items {
		if item.ID == vars["id"] {
			items[i] = itemUpdate
			w.Write([]byte("Item actualizado correctamente"))
			return
		}
	}

	w.Write([]byte("No se pudo actualizar el item"))
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    // TODO Función para eliminar un elemento
	vars := mux.Vars(r)

	for i, item := range items {
		if item.ID == vars["id"] {
			nuevoSlice := make([]Item, len(items)-1)

    		nuevoSlice = append(items[:i], items[i+1:]...)

			items = nuevoSlice
			w.Write([]byte("Item eliminado correctamente"))
			return
		}
	}

	w.Write([]byte("Item no pudo ser eliminado"))
}

func getItemByName(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	variables := r.URL.Query()
	name := variables.Get("name")

	for _, item := range items {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Item{})
}

func main(){
	router := mux.NewRouter()

	names := []string {"Auriculares", "Auriculares Bluetooth", "Teclado", "Monitor", "Mouse pad", "Mouse", "Parlante Bluetooth", "Helena", "Francisca", "Prudencia"}

	for i := 0; i < len(names); i++ {
		items[i] = Item{ID: strconv.Itoa(i), Name: names[i]}
	}

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/itemsbyname", getItemByName).Methods("GET")

    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}