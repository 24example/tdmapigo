package api

import (
	"log/slog"
	"os"
	"sync"

	"github.com/24example/tdmapigo/client"
	"github.com/24example/tdmapigo/models"
)

type (
	Api struct {
		token string
		mu    sync.Mutex
		log   *slog.Logger
	}

	Method interface {
		// Start Category Method
		Categories() ([]*models.Category, error)
		CategoriesByID(id int) (*models.Category, error)
		CategoriesTree() ([]*models.CategoryTree, error)
		CategoriesTreeByID(id int) ([]*models.CategoryTreeID, error)
		CategoriesTreeParentsByID(id int) ([]*models.CategoryTreeParentsID, error)
		// End Category Method

		// Start Product Method
		ProductList() ([]*models.ProductData, error)
		ProductListParams() ([]*models.ProductSKUParams, error)
		ProductListByID(id int) (*models.ProductID, error)
		ProductListByIDStocks(id int) (*models.ProductIDStocks, error)
		ProductListByIDPassports(id int) (*models.ProductIDPassport, error)
		ProductListByIDPrices(id int) (*models.ProductIDPrices, error)
		ProductListPricesfull() ([]*models.ProductIDPrices, error)
		ProductListByIDCerts(id int) (*models.ProductIDCerts, error)
		ProductListByIDSizes(id int) (*models.ProductIDSizes, error)
		ProductListSKUParams(sku string) (*models.ProductSKUParams, error)
		ProductListByIDWarranty(id int) (*models.ProductIDWarranty, error)
		// End Product Method

		// Start Stocks Method
		Stocks() ([]*models.Stock, error)
		// End Stocks Method

		// Start Passports Method
		Passports() (*models.Passport, error)
		// End Passports Method

		// Start Images Method
		Images() ([]*models.ImageData, error)
		ImagesByID(id int) (*models.ImageID, error)
		// End Images Method
	}
)

func NewApi(token string, logger *slog.Logger) Method {
	if logger == nil {
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return &Api{
		token: token,
		log:   logger,
		mu:    sync.Mutex{},
	}
}

func (a *Api) Client() *client.Client {
	return client.NewClient(a.token)
}
