package models

import "time"

type StockLink struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type StockShortStockStock struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type StockShortStock struct {
	StockID   int                  `json:"stock_id"`
	Count     int                  `json:"count"`
	UpdatedAt time.Time            `json:"updated_at"`
	ProductID int                  `json:"product_id"`
	Stock     StockShortStockStock `json:"stock"`
}

type Stock struct {
	ID          int               `json:"id"`
	SKU         string            `json:"sku"`
	ShortStocks []StockShortStock `json:"short_stocks"`
}
