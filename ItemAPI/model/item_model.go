package model

import (
	"time"
)

type Item struct {
	ID            string     `json:"id"`
	Customer_name string     `json:"name"`
	Order_date    *time.Time `json:"order_date"`
	Product       string     `json:"product"`
	Quantity      int64	     `json:"quantity"`
	Price         float64    `json:"price"`
	Details       *string    `json:"details"`
	TotalPrice	  float64	 `json:"total_price"`
}

type ItemDTO struct {
	Customer_name string  `json:"name"`
	Order_date    *string `json:"order_date"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Details       *string `json:"details"`
}

