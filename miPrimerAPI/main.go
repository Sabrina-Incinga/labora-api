package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
	ViewCount int  `json:"viewCount"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

var detailedItems []ItemDetails

var mutex sync.Mutex

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

func getItemDetails(id string) (ItemDetails, error) {
	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
	time.Sleep(2 * time.Millisecond)
	var foundItem *Item = nil
	for _, item := range items {
		if item.ID == id {
			foundItem = &item
			break
		}
	}

	if foundItem == nil{
		return ItemDetails{}, errors.New("Detalles del item no encontrado")
	}
	//Obviamente, aquí iria un SELECT si es SQL o un llamado a un servicio externo
	//pero esta busqueda del item junto con Details, la hacemos a mano.
	return ItemDetails{
		Item:    *foundItem, 
		Details: fmt.Sprintf("Detalles para el item %s", id),
	}, nil
}

func getDetailedItems(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	wg := &sync.WaitGroup{}
	detailsChannel := make(chan ItemDetails, len(items))
	itemIds := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} //slice con ids de items para forzar error con ids inexistentes
	for _, item := range itemIds{
		wg.Add(1) // Creamos el escucha, sin aun crearse la gorutina
		go func(id string) {
			defer wg.Done() //Completamos el trabajo del escucha, al final de esta ejecución
			detailedItem, err := getItemDetails(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			detailsChannel <- detailedItem
		}(item)
	}

    go func() {
		wg.Wait()
		close(detailsChannel)
	}()

	for details := range detailsChannel {
		detailedItems = append(detailedItems, details)
	}
	json.NewEncoder(w).Encode(detailedItems)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)


	searchedItem := getItemById(vars["id"])
	shouldReturn := searchedItem == nil
	if shouldReturn {
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

func getItemById(id string) *Item {
	var searchedItem *Item
	for i := 0; i < len(items); i++ {
		if items[i].ID == id {
			searchedItem = &items[i]
			break
		}
	}

	if searchedItem == nil {
		return nil
	}

	
	mutex.Lock()
	searchedItem.ViewCount++
	mutex.Unlock()

	return searchedItem
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

/* func main2(){
	router := mux.NewRouter()

	names := []string {"Auriculares", "Auriculares Bluetooth", "Teclado", "Monitor", "Mouse pad", "Mouse", "Parlante Bluetooth", "Micrófono", "Xbox", "Play Station 5"}

	for i := 0; i < len(names); i++ {
		items[i] = Item{ID: strconv.Itoa(i), Name: names[i], ViewCount: 0}
	}

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/details", getDetailedItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/itemsbyname", getItemByName).Methods("GET")

    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
} */

func main(){
	var wg sync.WaitGroup
	//var mu sync.Mutex

	item := items[3]
	initialViewCount := item.ViewCount

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			getItemById(item.ID)
			wg.Done()
		}()
	}

	wg.Wait()

	finalViewCount := item.ViewCount

	fmt.Printf("El valor esperado %d difiere del obtenido %d", initialViewCount+100, finalViewCount)
}