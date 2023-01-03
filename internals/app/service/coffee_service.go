package service

import (
	data "coffee/internals/app/data"
	model "coffee/internals/app/models"
	"errors"
)

type CoffeeService struct {
	Storage *data.CoffeeStorage
}

func NewCoffeeService(storage *data.CoffeeStorage) *CoffeeService {
	service := new(CoffeeService)
	service.Storage = storage
	return service
}

func (service *CoffeeService) FindCoffeeById(id int64) (model.Coffee, error) {
	var err error
	if id == 0 {
		err = errors.New("id is 0")
	}

	return service.Storage.FindCoffeeById(id), err
}
