package repository

import (
	"database/sql"
	"log"
	"math"
	"sync"
	"time"
	"github.com/labora-api/ItemAPI/model"
)

type PostgresDBHandler struct {
	Db *sql.DB
}

var mu sync.Mutex

func (p *PostgresDBHandler) GetAll(page int, itemsPerPage int) model.ItemsResponse {
	rows, err := p.Db.Query("SELECT id, customer_name, order_date, product, quantity, price, details  FROM items limit $1 offset $2 ", itemsPerPage, itemsPerPage*page)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var ItemsResponse model.ItemsResponse
	ItemsResponse.Items = make([]model.Item, 0)
	ItemsResponse.ItemCount = p.getItemsCount()

	for rows.Next() {
		var item model.Item

		err = rows.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)

		if err != nil {
			log.Fatal(err)
		}

		//item.TotalPrice = calculateTotalPrice(item.Price, item.Quantity)
		item.GetTotalPrice()

		ItemsResponse.Items = append(ItemsResponse.Items, item)

	}

	return ItemsResponse
}

func (p *PostgresDBHandler) GetItemById(id int) *model.ItemResponse {
	p.incrementViewCount(id, &mu)

	row := p.Db.QueryRow(`SELECT id, customer_name, order_date, product, quantity, price, details, "viewCount"  FROM items WHERE id = $1`, id)

	var item model.ItemResponse

	err := row.Scan(&item.Item.ID, &item.Item.Customer_name, &item.Item.Order_date, &item.Item.Product, &item.Item.Quantity, &item.Item.Price, &item.Item.Details, &item.ViewCount)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			log.Fatal(err)
		}
	}

	item.Item.TotalPrice = p.calculateTotalPrice(item.Item.Price, item.Item.Quantity)

	return &item
}

func (p *PostgresDBHandler) Create(item model.ItemDTO) int64 {
	var id int64
	orderDate, err := time.Parse("2006-01-02", *item.Order_date)
	if err != nil {
		log.Fatal(err)
	}

	row := p.Db.QueryRow(`INSERT INTO public.items(
								 customer_name, order_date, product, quantity, price, details)
								VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`, item.Customer_name, orderDate, item.Product, item.Quantity, item.Price, item.Details)

	row.Scan(&id)

	return id
}

func (p *PostgresDBHandler) Update(dto model.ItemDTO, id int) int64 {
	orderDate, err := time.Parse("2006-01-02", *dto.Order_date)
	if err != nil {
		log.Fatal(err)
	}

	result, err := p.Db.Exec(`UPDATE public.items
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

func (p *PostgresDBHandler) Delete(id int) int64 {
	result, err := p.Db.Exec(`DELETE FROM public.items
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

func (p *PostgresDBHandler) getItemsCount() int {
	row := p.Db.QueryRow("SELECT count(id) FROM items")

	var itemCount int

	err := row.Scan(&itemCount)

	if err != nil {
		log.Fatal(err)
	}

	return itemCount
}

func (p *PostgresDBHandler) calculateTotalPrice(price float64, quantity int64) float64 {
	totalPrice := price * float64(quantity)

	return math.Round(totalPrice*100) / 100
}

func (p *PostgresDBHandler) incrementViewCount(id int, mu *sync.Mutex) int64 {
	mu.Lock()
	result, err := p.Db.Exec(`UPDATE public.items
							SET "viewCount"= "viewCount"+1
							WHERE id=$1;`, id)
	mu.Unlock()

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}

func (p *PostgresDBHandler) getViewCount(id int) int {
	row := p.Db.QueryRow(`SELECT  "viewCount"  FROM items WHERE id = $1`, id)

	var viewCount int

	err := row.Scan(&viewCount)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0
		} else {
			log.Fatal(err)
		}
	}

	return viewCount
}
