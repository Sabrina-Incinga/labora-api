package model

import (
	"math"
	"time"
)

type Item struct {
	ID            int    `json:"id"`
	Customer_name string     `json:"name"`
	Order_date    *time.Time `json:"order_date"`
	Product       string     `json:"product"`
	Quantity      int64      `json:"quantity"`
	Price         float64    `json:"price"`
	Details       *string    `json:"details"`
	TotalPrice    float64    `json:"total_price"`
}

type ItemDTO struct {
	Customer_name string  `json:"name"`
	Order_date    *string `json:"order_date"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Details       *string `json:"details"`
}

func (i *Item) GetTotalPrice() {
	totalPrice := i.Price * float64(i.Quantity)

	i.TotalPrice = math.Round(totalPrice*100) / 100
}

type ItemsResponse struct {
	Items      []Item
	ItemCount int
}

type ItemResponse struct {
	Item      Item
	ViewCount int
}
