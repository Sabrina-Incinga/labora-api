package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labora-api/ItemAPI/model"
	"github.com/labora-api/ItemAPI/service"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
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

	items := service.GetAllItems(page, itemsPerPage)

	w.WriteHeader(http.StatusOK)
	itemsJson, _ := json.Marshal(items)
	w.Write(itemsJson)
}

func GetById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	item := service.GetItemById(id)

	if item == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Objeto con id %d no encontrado", id)))	
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var dto model.ItemDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	id := service.CreateItem(dto)

	if id != 0 {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("Objeto creado correctamente con id: %d", id)))	
	} else{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("No fue posible crear el item solicitado"))
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var dto model.ItemDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	vars := mux.Vars(r)

	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := service.UpdateItem(dto, id)

	if result {
		w.Header().Set("X-Message", fmt.Sprintf("Objeto con id: %d actualizado correctamente", id))
		w.WriteHeader(http.StatusNoContent)
	} else{
		w.Header().Set("X-Message",  "No fue posible actualizar el item solicitado")
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := service.DeleteItem(id)

	if result {
		w.Header().Set("X-Message", fmt.Sprintf("Objeto con id: %d eliminado correctamente", id))
		w.WriteHeader(http.StatusNoContent)
	} else{
		w.Header().Set("X-Message",  "No fue posible eliminar el item solicitado")
		w.WriteHeader(http.StatusNoContent)
	}
}