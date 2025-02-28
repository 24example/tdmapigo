package api

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/24example/tdmapigo/models"
)

func (a *Api) Images() ([]*models.ImageData, error) {

	const op = "API.Images"

	log := a.log.With(
		slog.String("event", op),
	)

	var products []*models.ImageData

	limit := 100
	page := 1

	for {
		response, err := a.Client().
			Request.
			SetResult(&models.Response[[]*models.ImageData]{}).
			SetQueryParam("page", strconv.Itoa(page)).
			SetQueryParam("per_page", strconv.Itoa(limit)).
			Get("/images")
		if err != nil {
			log.Error("failed getting images to api. Error: " + err.Error())
			return nil, err
		}

		if response.StatusCode() != 200 {
			log.Error("failed getting images to api. Status: " + response.Status())
			return nil, errors.New(response.Status())
		}

		result := response.Result().(*models.Response[[]*models.ImageData])

		products = append(products, result.Data...)

		if page == result.LastPage {
			break
		}

		page++

	}

	return products, nil

}

func (a *Api) ImagesByID(id int) (*models.ImageID, error) {

	const op = "API.ImageByID"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.ImageID{}).
		Get("/images/" + strconv.Itoa(id))
	if err != nil {
		log.Error("failed getting images by id to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting images by id to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.ImageID), nil

}
