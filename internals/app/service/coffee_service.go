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

func (service *CoffeeService) FindById(id int64) (model.Coffee, error) {
	return service.Storage.FindById(id)
}

func (service *CoffeeService) AddNew(coffee model.Coffee) error {
	if coffee.Name == "" {
		return errors.New("Coffee name is empty")
	}

	if coffee.Price == 0.0 {
		return errors.New("Coffee price is empty")
	}

	return service.AddNew(coffee)
}

func (service *CoffeeService) GetAll() []model.Coffee {
	return service.Storage.GetAll()
}
