package repository

import (
	"database/sql"
	"log"
	"math"
	"time"

	"github.com/labora-api/ItemAPI/config"
	"github.com/labora-api/ItemAPI/model"
)

var connection, error = config.GetConnection()

func GetAll() []model.Item {
	rows, err := connection.Query("SELECT id, customer_name, order_date, product, quantity, price, details  FROM items")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var items []model.Item = make([]model.Item, 0)

	for rows.Next() {
		var item model.Item

		err = rows.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)

		if err != nil {
			log.Fatal(err)
		}

		//item.TotalPrice = calculateTotalPrice(item.Price, item.Quantity)
		item.GetTotalPrice()
		
		items = append(items, item)

	}

	return items
}

func GetItemById(id int) *model.Item {
	row := connection.QueryRow("SELECT id, customer_name, order_date, product, quantity, price, details  FROM items WHERE id = $1", id)

	var item model.Item

	err := row.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil
		} else{
			log.Fatal(err)
		}
	}

	item.TotalPrice = calculateTotalPrice(item.Price, item.Quantity)

	return &item
}

func Create(item model.ItemDTO) int64 {
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

func Update(dto model.ItemDTO, id int) int64 {
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

func calculateTotalPrice(price float64, quantity int64) float64{
	totalPrice := price * float64(quantity)

	return math.Round(totalPrice*100)/100
}