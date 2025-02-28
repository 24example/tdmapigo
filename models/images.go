package models

import "time"

//
// General
type ImagePrice struct {
	ID         int       `json:"id"`
	ProductID  int       `json:"product_id"`
	CurrencyID int       `json:"currency_id"`
	TdmID      string    `json:"tdm_id"`
	Price      string    `json:"price"`
	Type       string    `json:"type"`
	IsSumTax   bool      `json:"is_sum_tax"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ImageImage struct {
	ID             int       `json:"id"`
	IsMain         bool      `json:"is_main"`
	URL            string    `json:"url"`
	ImagetableID   int       `json:"imagetable_id"`
	ImagetableType string    `json:"imagetable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Order          int       `json:"order"`
	Enabled        bool      `json:"enabled"`
	Transcode      bool      `json:"transcode"`
}

// General
//

//
// Image
type ImageData struct {
	ID     int          `json:"id"`
	SKU    string       `json:"sku"`
	Price  string       `json:"price"`
	Images []ImageImage `json:"images"`
	Prices []ImagePrice `json:"prices"`
}

type ImageLink struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type Image struct {
	CurrentPage  int         `json:"current_page"`
	Data         []ImageData `json:"data"`
	FirstPageURL string      `json:"first_page_url"`
	From         int         `json:"from"`
	LastPage     int         `json:"last_page"`
	LastPageURL  string      `json:"last_page_url"`
	Links        []ImageLink `json:"links"`
	NextPageURL  *string     `json:"next_page_url"`
	Path         string      `json:"path"`
	PerPage      int         `json:"per_page"`
	PrevPageURL  *string     `json:"prev_page_url"`
	To           int         `json:"to"`
	Total        int         `json:"total"`
}

// Image
//

//
// ImageID
type ImageID struct {
	ID     int          `json:"id"`
	SKU    string       `json:"sku"`
	Price  string       `json:"price"`
	Images []ImageImage `json:"images"`
	Prices []ImagePrice `json:"prices"`
}

// ImageID
//
