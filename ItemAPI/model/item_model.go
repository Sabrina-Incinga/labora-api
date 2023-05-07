package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/labora-api/ItemAPI/config"
)

type Item struct {
	ID            string     `json:"id"`
	Customer_name string     `json:"name"`
	Order_date    *time.Time `json:"order_date"`
	Product       string     `json:"product"`
	Quantity      string     `json:"quantity"`
	Price         string     `json:"price"`
	Details       *string    `json:"details"`
}

type ItemDTO struct {
	Customer_name string  `json:"name"`
	Order_date    *string `json:"order_date"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Details       *string `json:"details"`
}

var connection, error = config.GetConnection()

func GetAll() []Item {
	rows, err := connection.Query("SELECT id, customer_name, order_date, product, quantity, price, details  FROM items")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var items []Item = make([]Item, 0)

	for rows.Next() {
		var item Item

		err = rows.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)

	}

	return items
}

func GetItemById(id int) *Item {
	row := connection.QueryRow("SELECT id, customer_name, order_date, product, quantity, price, details  FROM items WHERE id = $1", id)

	var item Item

	err := row.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil
		} else{
			log.Fatal(err)
		}
	}

	return &item
}

func Create(item ItemDTO) int64 {
	var id int64
	orderDate, err := time.Parse("2006-01-02", *item.Order_date)
	if err != nil {
		log.Fatal(err)
	}

	row := connection.QueryRow(`INSERT INTO public.items(
								 customer_name, order_date, product, quantity, price, details)
								VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`, item.Customer_name, orderDate, item.Product, item.Quantity, item.Price, item.Details)

	row.Scan(&id)

	return id
}

func Update(dto ItemDTO, id int) int64 {
	orderDate, err := time.Parse("2006-01-02", *dto.Order_date)
	if err != nil {
		log.Fatal(err)
	}

	result, err := connection.Exec(`UPDATE public.items
				SET customer_name=$1, order_date=$2, product=$3, quantity=$4, price=$5, details=$6
				WHERE id = $7;`, dto.Customer_name, orderDate, dto.Product, dto.Quantity, dto.Price, dto.Details, id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}

func Delete(id int) int64 {
	result, err := connection.Exec(`DELETE FROM public.items
									WHERE id = $1;`, id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}
