package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func init(){
	names := []string {"Auriculares", "Auriculares Bluetooth", "Teclado", "Monitor", "Mouse pad", "Mouse", "Parlante Bluetooth", "Micrófono", "Xbox", "Play Station 5"}

	for i := 0; i < len(names); i++ {
		items[i] = Item{ID: strconv.Itoa(i), Name: names[i], ViewCount: 0}
	}

	fmt.Printf("%+v", items)
}

func TestGetItemById(t *testing.T){
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

	if finalViewCount != initialViewCount+100 {
		t.Errorf("El valor esperado %d difiere del obtenido %d", initialViewCount+100, finalViewCount)
	}

}