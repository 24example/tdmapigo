package api

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/24example/tdmapigo/models"
)

func (a *Api) Stocks() ([]*models.Stock, error) {

	const op = "API.Stocks"

	log := a.log.With(
		slog.String("event", op),
	)

	var stocks []*models.Stock

	limit := 100
	page := 1

	for {
		response, err := a.Client().
			Request.
			SetResult(&models.Response[[]*models.Stock]{}).
			SetQueryParam("page", strconv.Itoa(page)).
			SetQueryParam("per_page", strconv.Itoa(limit)).
			Get("/stocks")
		if err != nil {
			log.Error("failed getting stocks to api. Error: " + err.Error())
			return nil, err
		}

		if response.StatusCode() != 200 {
			log.Error("failed getting stocks to api. Status: " + response.Status())
			return nil, errors.New(response.Status())
		}

		result := response.Result().(*models.Response[[]*models.Stock])

		stocks = append(stocks, result.Data...)

		if page == result.LastPage {
			break
		}

		page++

	}

	return stocks, nil
}
