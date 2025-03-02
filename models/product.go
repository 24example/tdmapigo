package models

//
// General

type Unit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCategoryPivot struct {
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}

type ProductCategory struct {
	ID    int                  `json:"id"`
	Name  string               `json:"name"`
	Slug  string               `json:"slug"`
	Pivot ProductCategoryPivot `json:"pivot"`
}

type CertificatePivot struct {
	FileableType  string `json:"fileable_type"`
	FileableID    int    `json:"fileable_id"`
	FileStorageID int    `json:"file_storage_id"`
}

type Certificate struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Preview   *string          `json:"preview"`
	File      string           `json:"file"`
	Type      string           `json:"type"`
	Title     string           `json:"title"`
	Subtitle  *string          `json:"subtitle"`
	Size      int              `json:"size"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	Enabled   bool             `json:"enabled"`
	OldIdent  *int             `json:"old_ident"`
	Pivot     CertificatePivot `json:"pivot"`
}

type Size struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Type      string `json:"type"`
	Count     int    `json:"count"`
	Height    string `json:"height"`
	Width     string `json:"width"`
	Depth     string `json:"depth"`
	Weight    string `json:"weight"`
	Volume    string `json:"volume"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Barcode   string `json:"barcode"`
}

type ParamPivot struct {
	ProductID int    `json:"product_id"`
	ParamID   int    `json:"param_id"`
	Value     string `json:"value"`
	EtimID    string `json:"etim_id"`
}

type Param struct {
	ID            int         `json:"id"`
	Slug          string      `json:"slug"`
	Name          string      `json:"name"`
	Order         int         `json:"order"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Tab           string      `json:"tab"`
	Type          *string     `json:"type"`
	FilterEnabled bool        `json:"filter_enabled"`
	Pivot         *ParamPivot `json:"pivot"`
}

type Link struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

// General
//

//
// Product

type ProductDataStockStock struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	TdmID       string `json:"tdm_id"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type ProductDataStock struct {
	ID             int                   `json:"id"`
	StockID        int                   `json:"stock_id"`
	ProductID      int                   `json:"product_id"`
	Count          int                   `json:"count"`
	CreatedAt      string                `json:"created_at"`
	UpdatedAt      string                `json:"updated_at"`
	CountRemaining int                   `json:"count_remaining"`
	Stock          ProductDataStockStock `json:"stock"`
	Reserves       []interface{}         `json:"reserves"`
}
type ProductDataMainImage struct {
	ID             int    `json:"id"`
	IsMain         bool   `json:"is_main"`
	URL            string `json:"url"`
	ImagetableID   int    `json:"imagetable_id"`
	ImagetableType string `json:"imagetable_type"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Order          int    `json:"order"`
	Enabled        bool   `json:"enabled"`
	Transcode      bool   `json:"transcode"`
}

type ProductData struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name"`
	Slug           string               `json:"slug"`
	SKU            string               `json:"sku"`
	IsActive       bool                 `json:"is_active"`
	IsStock        bool                 `json:"is_stock"`
	Multiple       int                  `json:"multiple"`
	Barcode        string               `json:"barcode"`
	Type           string               `json:"type"`
	DateAssortment string               `json:"date_assortment"`
	WarrantyPeriod int                  `json:"warranty_period"`
	WarrantyEd     string               `json:"warranty_ed"`
	Etim           string               `json:"etim"`
	CountryName    string               `json:"country_name"`
	Series         string               `json:"series"`
	Unit           Unit                 `json:"unit"`
	Categories     []ProductCategory    `json:"categories"`
	Certificates   []Certificate        `json:"certificates"`
	Sizes          []Size               `json:"sizes"`
	Tags           []interface{}        `json:"tags"`
	Stocks         []ProductDataStock   `json:"stocks"`
	MainImage      ProductDataMainImage `json:"main_image"`
	Params         []Param              `json:"params"`
}

// Product
//

//
// ProductID

type ProductID struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Slug           string            `json:"slug"`
	SKU            string            `json:"sku"`
	IsActive       bool              `json:"is_active"`
	IsStock        bool              `json:"is_stock"`
	Multiple       int               `json:"multiple"`
	Barcode        string            `json:"barcode"`
	Type           string            `json:"type"`
	DateAssortment string            `json:"date_assortment"`
	WarrantyPeriod int               `json:"warranty_period"`
	WarrantyEd     string            `json:"warranty_ed"`
	Etim           string            `json:"etim"`
	CountryName    string            `json:"country_name"`
	Series         string            `json:"series"`
	Unit           Unit              `json:"unit"`
	Categories     []ProductCategory `json:"categories"`
}

// ProductID
//

//
// ProductIDStocks

type ProductIDStocksShortStockStock struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ProductIDStocksShortStock struct {
	StockID   int                            `json:"stock_id"`
	Count     int                            `json:"count"`
	UpdatedAt string                         `json:"updated_at"`
	ProductID int                            `json:"product_id"`
	Stock     ProductIDStocksShortStockStock `json:"stock"`
}

type ProductIDStocks struct {
	ID          int                         `json:"id"`
	SKU         string                      `json:"sku"`
	ShortStocks []ProductIDStocksShortStock `json:"short_stocks"`
}

// ProductIDStocks
//

//
// ProductIDPassport

type ProductIDPassport struct {
	ID        int           `json:"id"`
	SKU       string        `json:"sku"`
	Passports []Certificate `json:"passports"`
}

// ProductIDPassport
//

//
// ProductIDPrices

type ProductIDPricesPrice struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Price     string  `json:"price"`
	Type      string  `json:"type"`
	IsSumTax  bool    `json:"is_sum_tax"`
	UpdatedAt *string `json:"updated_at"`
}

type ProductIDPrices struct {
	ID     int                    `json:"id"`
	SKU    string                 `json:"sku"`
	Price  string                 `json:"price"`
	Prices []ProductIDPricesPrice `json:"prices"`
}

// ProductIDPrices
//

//
// ProductIDCerts

type ProductIDCerts struct {
	ID           int           `json:"id"`
	SKU          string        `json:"sku"`
	Certificates []Certificate `json:"certificates"`
}

// ProductIDCerts
//

//
// ProductIDSizes

type ProductIDSizes struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Sizes []Size `json:"sizes"`
}

// ProductIDSizes
//

//
// ProductSKUParams

type ProductSKUParams struct {
	ID     int     `json:"id"`
	SKU    string  `json:"sku"`
	Params []Param `json:"params"`
}

// ProductSKUParams
//

//
// ProductIDWarranty

type ProductIDWarranty struct {
	ID             int    `json:"id"`
	SKU            string `json:"sku"`
	WarrantyPeriod int    `json:"warranty_period"`
	WarrantyEd     string `json:"warranty_ed"`
}

// ProductIDWarranty
//
