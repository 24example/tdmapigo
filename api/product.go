package api

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/24example/tdmapigo/models"
)

func (a *Api) ProductList() ([]*models.ProductData, error) {

	const op = "API.ProductList"

	log := a.log.With(
		slog.String("event", op),
	)

	var products []*models.ProductData

	limit := 100
	page := 1

	for {
		response, err := a.Client().
			Request.
			SetResult(&models.Response[[]*models.ProductData]{}).
			SetQueryParam("page", strconv.Itoa(page)).
			SetQueryParam("per_page", strconv.Itoa(limit)).
			Get("/product_list")
		if err != nil {
			log.Error("failed getting products list to api. Error: " + err.Error())
			return nil, err
		}

		if response.StatusCode() != 200 {
			log.Error("failed getting products list to api. Status: " + response.Status())
			return nil, errors.New(response.Status())
		}

		result := response.Result().(*models.Response[[]*models.ProductData])

		products = append(products, result.Data...)

		if page == result.LastPage {
			break
		}

		page++

	}

	return products, nil

}

func (a *Api) ProductListParams() ([]*models.ProductSKUParams, error) {

	const op = "API.ProductListParams"

	log := a.log.With(
		slog.String("event", op),
	)

	var products []*models.ProductSKUParams

	limit := 100
	page := 1

	for {
		response, err := a.Client().
			Request.
			SetResult(&models.Response[[]*models.ProductSKUParams]{}).
			SetQueryParam("page", strconv.Itoa(page)).
			SetQueryParam("per_page", strconv.Itoa(limit)).
			Get("/product_list/params")
		if err != nil {
			log.Error("failed getting products list params to api. Error: " + err.Error())
			return nil, err
		}

		if response.StatusCode() != 200 {
			log.Error("failed getting products list params to api. Status: " + response.Status())
			return nil, errors.New(response.Status())
		}

		result := response.Result().(*models.Response[[]*models.ProductSKUParams])

		products = append(products, result.Data...)

		if page == result.LastPage {
			break
		}

		page++

	}

	return products, nil

}

func (a *Api) ProductListByID(id int) (*models.ProductID, error) {

	const op = "API.ProductListByID"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductID{}).
		Get("/product_list/" + strconv.Itoa(id))
	if err != nil {
		log.Error("failed getting product list by id to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductID), nil

}

func (a *Api) ProductListByIDStocks(id int) (*models.ProductIDStocks, error) {

	const op = "API.ProductListByIDStocks"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDStocks{}).
		Get("/product_list/" + strconv.Itoa(id) + "/stocks")
	if err != nil {
		log.Error("failed getting product list by id stocks to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id stocks to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDStocks), nil

}

func (a *Api) ProductListByIDPassports(id int) (*models.ProductIDPassport, error) {

	const op = "API.ProductListByIDPassports"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDPassport{}).
		Get("/product_list/" + strconv.Itoa(id) + "/passport")
	if err != nil {
		log.Error("failed getting product list by id passports to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id passports to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDPassport), nil

}

func (a *Api) ProductListByIDPrices(id int) (*models.ProductIDPrices, error) {

	const op = "API.ProductListByIDPrices"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDPrices{}).
		Get("/product_list/" + strconv.Itoa(id) + "/prices")
	if err != nil {
		log.Error("failed getting product list by id prices to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id prices to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDPrices), nil

}

func (a *Api) ProductListPricesfull() ([]*models.ProductIDPrices, error) {

	const op = "API.ProductListPricesfull"

	log := a.log.With(
		slog.String("event", op),
	)

	var products []*models.ProductIDPrices

	limit := 100
	page := 1

	for {
		response, err := a.Client().
			Request.
			SetResult(&models.Response[[]*models.ProductIDPrices]{}).
			SetQueryParam("page", strconv.Itoa(page)).
			SetQueryParam("per_page", strconv.Itoa(limit)).
			Get("/product_list/pricesfull")
		if err != nil {
			log.Error("failed getting products list params to api. Error: " + err.Error())
			return nil, err
		}

		if response.StatusCode() != 200 {
			log.Error("failed getting products list params to api. Status: " + response.Status())
			return nil, errors.New(response.Status())
		}

		result := response.Result().(*models.Response[[]*models.ProductIDPrices])

		products = append(products, result.Data...)

		if page == result.LastPage {
			break
		}

		page++

	}

	return products, nil

}

func (a *Api) ProductListByIDCerts(id int) (*models.ProductIDCerts, error) {

	const op = "API.ProductListByIDCerts"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDCerts{}).
		Get("/product_list/" + strconv.Itoa(id) + "/certs")
	if err != nil {
		log.Error("failed getting product list by id certs to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id certs to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDCerts), nil

}

func (a *Api) ProductListByIDSizes(id int) (*models.ProductIDSizes, error) {

	const op = "API.ProductListByIDSizes"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDSizes{}).
		Get("/product_list/" + strconv.Itoa(id) + "/sizes")
	if err != nil {
		log.Error("failed getting product list by id sizes to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id sizes to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDSizes), nil

}

func (a *Api) ProductListSKUParams(sku string) (*models.ProductSKUParams, error) {

	const op = "API.ProductListSKUParams"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductSKUParams{}).
		Get("/product_list/" + sku + "/params")
	if err != nil {
		log.Error("failed getting product list by sku params to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by sku params to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductSKUParams), nil

}

func (a *Api) ProductListByIDWarranty(id int) (*models.ProductIDWarranty, error) {

	const op = "API.ProductListByIDWarranty"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ProductIDWarranty{}).
		Get("/product_list/" + strconv.Itoa(id) + "/warranty")
	if err != nil {
		log.Error("failed getting product list by id warranty to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting product list by id warranty to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ProductIDWarranty), nil

}
