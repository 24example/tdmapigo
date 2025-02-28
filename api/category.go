package api

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/24example/tdmapigo/models"
)

func (a *Api) Categories() ([]*models.Category, error) {

	const op = "API.Categories"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&[]*models.Category{}).
		Get("/categories")
	if err != nil {
		log.Error("failed getting categories to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting categories to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	result := response.Result().(*[]*models.Category)

	return *result, nil
}

func (a *Api) CategoriesByID(id int) (*models.Category, error) {

	const op = "API.CategoriesByID"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&models.Category{}).
		Get("/categories/" + strconv.Itoa(id))
	if err != nil {
		log.Error("failed getting categories by id to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting categories by id to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	return response.Result().(*models.Category), nil

}

func (a *Api) CategoriesTree() ([]*models.CategoryTree, error) {

	const op = "API.CategoriesTree"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&[]*models.CategoryTree{}).
		Get("/categories/tree")
	if err != nil {
		log.Error("failed getting categories tree to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting categories tree to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	result := response.Result().(*[]*models.CategoryTree)

	return *result, nil
}

func (a *Api) CategoriesTreeByID(id int) ([]*models.CategoryTreeID, error) {

	const op = "API.CategoriesTreeByID"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&[]*models.CategoryTreeID{}).
		Get("/categories/tree/" + strconv.Itoa(id))
	if err != nil {
		log.Error("failed getting categories tree to api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting categories tree to api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	result := response.Result().(*[]*models.CategoryTreeID)

	return *result, nil
}

func (a *Api) CategoriesTreeParentsByID(id int) ([]*models.CategoryTreeParentsID, error) {

	const op = "API.CategoriesTreeParentsByID"

	log := a.log.With(
		slog.String("event", op),
	)

	response, err := a.Client().
		Request.
		SetResult(&[]*models.CategoryTreeParentsID{}).
		Get("/categories/tree/parents/" + strconv.Itoa(id))
	if err != nil {
		log.Error("failed getting categories tree to parents api. Error: " + err.Error())
		return nil, err
	}

	if response.StatusCode() != 200 {
		log.Error("failed getting categories tree to parents api. Status: " + response.Status())
		return nil, errors.New(response.Status())
	}

	result := response.Result().(*[]*models.CategoryTreeParentsID)

	return *result, nil
}
