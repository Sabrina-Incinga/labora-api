package main

import (
	"encoding/json"
	"math"
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
	variables := r.URL.Query()
	pageStr := variables.Get("page")
	itemsPerPageStr := variables.Get("itemsPerPage")
	var page, itemsPerPage int
	var err error
	if pageStr == ""{
		page = 0
	} else{
		page, err = strconv.Atoi(pageStr)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error al convertir página a int"))
			return
		}
	}

	if pageStr == ""{
		itemsPerPage = 5
	} else{
		itemsPerPage, err = strconv.Atoi(itemsPerPageStr)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error al convertir cantidad de ítems por página a int"))
			return
		}
	}

	var sliceToShow []Item 
	init := page*itemsPerPage
	limit := itemsPerPage*(page+1)
	nroPage := float64(len(items)) / float64(itemsPerPage)
    nroPage = math.Ceil(nroPage)
	if page <= int(nroPage) {
		if limit <= len(items) {
			sliceToShow = items[init:limit]
		} else{
			sliceToShow = items[init:]
		}
	}

	w.WriteHeader(http.StatusOK)
	itemsJson, _ := json.Marshal(sliceToShow)
	w.Write(itemsJson)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var searchedItem *Item

	for i := 0; i < len(items); i++ {
		if items[i].ID == vars["id"] {
			searchedItem = &items[i]
			break
		}
	}

	if searchedItem == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item no encontrado"))
		return
	}

	itemJson, err := json.Marshal(searchedItem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No se pudo convertir el elemento encontrado"))
	}
	w.WriteHeader(http.StatusOK)
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
	vars := mux.Vars(r)

	for i, item := range items {
		if item.ID == vars["id"] {
			newSlice := make([]Item, len(items)-1)

    		newSlice = append(items[:i], items[i+1:]...)

			items = newSlice
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