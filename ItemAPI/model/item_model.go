package model

import (
	"log"

	"github.com/labora-api/ItemAPI/config"
)

type Item struct {
	ID            string `json:"id"`
	Customer_name string `json:"name"`
	Product       string `json:"product"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	Details       string `json:"details"`
}

var connection, error = config.GetConnection()

func GetAll() []Item{
	rows, err := connection.Query("SELECT id, customer_name, product, quantity, price  FROM items")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var items []Item = make([]Item, 0)

	for rows.Next(){
		var item Item
		err := rows.Scan(&item.ID, &item.Customer_name, &item.Product, &item.Quantity, &item.Price)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

	}

	return items
}
