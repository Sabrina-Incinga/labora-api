package model

import (
	"time"
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

